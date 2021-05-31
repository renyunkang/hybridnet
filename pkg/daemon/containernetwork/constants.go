/*
Copyright 2021 The Rama Authors.

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

package containernetwork

const (
	DockerNetnsDir     = "/var/run/docker/netns"
	ContainerdNetnsDir = "/var/run/netns/"

	ContainerHostLinkSuffix = "_h"
	ContainerHostLinkMac    = "ee:ee:ee:ee:ee:ee"
	ContainerInitLinkSuffix = "_c"
	VxlanLinkInfix          = ".vxlan"
	ContainerNicName        = "eth0"

	ProxyArpSysctl       = "/proc/sys/net/ipv4/conf/%s/proxy_arp"
	ProxyDelaySysctl     = "/proc/sys/net/ipv4/neigh/%s/proxy_delay"
	RouteLocalNetSysctl  = "/proc/sys/net/ipv4/conf/%s/route_localnet"
	IPv4ForwardingSysctl = "/proc/sys/net/ipv4/conf/%s/forwarding"

	RpFilterSysctl = "/proc/sys/net/ipv4/conf/%s/rp_filter"

	ProxyNdpSysctl       = "/proc/sys/net/ipv6/conf/%s/proxy_ndp"
	IPv6ForwardingSysctl = "/proc/sys/net/ipv6/conf/%s/forwarding"

	IPv4AppSolicitSysctl = "/proc/sys/net/ipv4/neigh/%s/app_solicit"
	IPv6AppSolicitSysctl = "/proc/sys/net/ipv6/neigh/%s/app_solicit"

	AcceptDADSysctl = "/proc/sys/net/ipv6/conf/%s/accept_dad"
	AcceptRASysctl  = "/proc/sys/net/ipv6/conf/%s/accept_ra"

	IPv4BaseReachableTimeMSSysctl = "/proc/sys/net/ipv4/neigh/%s/base_reachable_time_ms"
	IPv6BaseReachableTimeMSSysctl = "/proc/sys/net/ipv6/neigh/%s/base_reachable_time_ms"

	// IP Masks that have no effect on IP Address
	DefaultIP4Mask = "255.255.255.255"
	DefaultIP6Mask = "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"
)