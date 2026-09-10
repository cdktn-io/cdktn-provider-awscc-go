// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evsenvironment


type EvsEnvironmentConnectivityInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evs_environment#private_route_server_peerings EvsEnvironment#private_route_server_peerings}.
	PrivateRouteServerPeerings *[]*string `field:"optional" json:"privateRouteServerPeerings" yaml:"privateRouteServerPeerings"`
}

