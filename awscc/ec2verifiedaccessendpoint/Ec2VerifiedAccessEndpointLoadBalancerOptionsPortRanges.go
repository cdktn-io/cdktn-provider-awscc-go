// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessendpoint


type Ec2VerifiedAccessEndpointLoadBalancerOptionsPortRanges struct {
	// The first port in the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_endpoint#from_port Ec2VerifiedAccessEndpoint#from_port}
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The last port in the range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_endpoint#to_port Ec2VerifiedAccessEndpoint#to_port}
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

