/*
 Copyright 2021 The Hybridnet Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package validating

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"reflect"

	multiclusterv1 "github.com/alibaba/hybridnet/pkg/apis/multicluster/v1"
	networkingv1 "github.com/alibaba/hybridnet/pkg/apis/networking/v1"
	"github.com/alibaba/hybridnet/pkg/feature"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var networkGVK = gvkConverter(networkingv1.GroupVersion.WithKind("Network"))

func init() {
	createHandlers[networkGVK] = NetworkCreateValidation
	updateHandlers[networkGVK] = NetworkUpdateValidation
	deleteHandlers[networkGVK] = NetworkDeleteValidation
}

func NetworkCreateValidation(ctx context.Context, req *admission.Request, handler *Handler) admission.Response {
	network := &networkingv1.Network{}
	err := handler.Decoder.Decode(*req, network)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	switch networkingv1.GetNetworkType(network) {
	case networkingv1.NetworkTypeUnderlay:
		if network.Spec.NodeSelector == nil || len(network.Spec.NodeSelector) == 0 {
			return admission.Denied("must have node selector for underlay network")
		}
	case networkingv1.NetworkTypeOverlay:
		// check uniqueness
		networks := &networkingv1.NetworkList{}
		if err = handler.Client.List(ctx, networks); err != nil {
			return admission.Errored(http.StatusInternalServerError, err)
		}
		for i := range networks.Items {
			if networkingv1.GetNetworkType(&networks.Items[i]) == networkingv1.NetworkTypeOverlay {
				return admission.Denied("must have one overlay network at most")
			}
		}

		// check node selector
		if network.Spec.NodeSelector != nil && len(network.Spec.NodeSelector) > 0 {
			return admission.Denied("must not assign node selector for overlay network")
		}

		// check net id
		if network.Spec.NetID == nil {
			return admission.Denied("must assign net ID for overlay network")
		}
	default:
		return admission.Denied(fmt.Sprintf("unknown network type %s", networkingv1.GetNetworkType(network)))
	}

	switch networkingv1.GetNetworkMode(network) {
	case networkingv1.NetworkModeBGP:
		// check net id
		if network.Spec.NetID == nil {
			return admission.Denied("must assign net ID for bgp network")
		}

		if len(network.Spec.Config.BGPPeers) != 1 {
			return admission.Denied("one and only one bgp router need to be set")
		}

		for _, peer := range network.Spec.Config.BGPPeers {
			if net.ParseIP(peer.Address) == nil {
				return admission.Denied(fmt.Sprintf("invalid bgp peer ip address %v", peer.Address))
			}
		}
	case networkingv1.NetworkModeVlan:
	case networkingv1.NetworkModeVxlan:
	default:
		return admission.Denied(fmt.Sprintf("unknown network mode %s", networkingv1.GetNetworkMode(network)))
	}

	return admission.Allowed("validation pass")
}

func NetworkUpdateValidation(ctx context.Context, req *admission.Request, handler *Handler) admission.Response {
	var err error
	oldN, newN := &networkingv1.Network{}, &networkingv1.Network{}
	if err = handler.Decoder.DecodeRaw(req.Object, newN); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}
	if err = handler.Decoder.DecodeRaw(req.OldObject, oldN); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	if networkingv1.GetNetworkType(oldN) != networkingv1.GetNetworkType(newN) {
		return admission.Denied("network type must not be changed")
	}

	switch networkingv1.GetNetworkType(newN) {
	case networkingv1.NetworkTypeUnderlay:
		if newN.Spec.NodeSelector == nil || len(newN.Spec.NodeSelector) == 0 {
			return admission.Denied("must have node selector for underlay network")
		}
	case networkingv1.NetworkTypeOverlay:
		if newN.Spec.NodeSelector != nil && len(newN.Spec.NodeSelector) > 0 {
			return admission.Denied("node selector must not be assigned for overlay network")
		}
	default:
		return admission.Denied(fmt.Sprintf("unknown network type %s", networkingv1.GetNetworkType(newN)))
	}

	if oldN.Spec.Mode != newN.Spec.Mode {
		return admission.Denied("network mode must not be changed")
	}

	switch networkingv1.GetNetworkMode(newN) {
	case networkingv1.NetworkModeBGP:
		if len(newN.Spec.Config.BGPPeers) != 1 {
			return admission.Denied("one and only one bgp router need to be set")
		}

		for _, peer := range newN.Spec.Config.BGPPeers {
			if net.ParseIP(peer.Address) == nil {
				return admission.Denied(fmt.Sprintf("invalid bgp peer ip address %v", peer.Address))
			}
		}
	case networkingv1.NetworkModeVlan:
	case networkingv1.NetworkModeVxlan:
	default:
		return admission.Denied(fmt.Sprintf("unknown network mode %s", networkingv1.GetNetworkMode(newN)))
	}

	if !reflect.DeepEqual(oldN.Spec.NetID, newN.Spec.NetID) {
		return admission.Denied("net ID must not be changed")
	}

	return admission.Allowed("validation pass")
}

func NetworkDeleteValidation(ctx context.Context, req *admission.Request, handler *Handler) admission.Response {
	var err error
	network := &networkingv1.Network{}
	if err = handler.Client.Get(ctx, types.NamespacedName{Name: req.Name}, network); err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	subnetList := &networkingv1.SubnetList{}
	if err = handler.Client.List(ctx, subnetList); err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	for _, subnet := range subnetList.Items {
		if subnet.Spec.Network == network.Name {
			return admission.Denied(fmt.Sprintf("still have child subnet %s", subnet.Name))
		}
	}

	if network.Spec.Type == networkingv1.NetworkTypeOverlay && feature.MultiClusterEnabled() {
		remoteClusterList := &multiclusterv1.RemoteClusterList{}
		if err = handler.Client.List(ctx, remoteClusterList); err != nil {
			return admission.Errored(http.StatusInternalServerError, err)
		}
		if len(remoteClusterList.Items) != 0 {
			return admission.Denied(fmt.Sprintf("still have remote cluster. number=%v", len(remoteClusterList.Items)))
		}
	}

	return admission.Allowed("validation pass")
}
