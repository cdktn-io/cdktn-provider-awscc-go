// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessendpoint


type Ec2VerifiedAccessEndpointLoadBalancerOptions struct {
	// The ARN of the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_endpoint#load_balancer_arn Ec2VerifiedAccessEndpoint#load_balancer_arn}
	LoadBalancerArn *string `field:"optional" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// The IP port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_endpoint#port Ec2VerifiedAccessEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The list of port range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_endpoint#port_ranges Ec2VerifiedAccessEndpoint#port_ranges}
	PortRanges interface{} `field:"optional" json:"portRanges" yaml:"portRanges"`
	// The IP protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_endpoint#protocol Ec2VerifiedAccessEndpoint#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The IDs of the subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_verified_access_endpoint#subnet_ids Ec2VerifiedAccessEndpoint#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

