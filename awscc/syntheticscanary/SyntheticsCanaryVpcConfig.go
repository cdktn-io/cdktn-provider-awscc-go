// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package syntheticscanary


type SyntheticsCanaryVpcConfig struct {
	// Allow outbound IPv6 traffic on VPC canaries that are connected to dual-stack subnets if set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#ipv_6_allowed_for_dual_stack SyntheticsCanary#ipv_6_allowed_for_dual_stack}
	Ipv6AllowedForDualStack interface{} `field:"optional" json:"ipv6AllowedForDualStack" yaml:"ipv6AllowedForDualStack"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#security_group_ids SyntheticsCanary#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#subnet_ids SyntheticsCanary#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/synthetics_canary#vpc_id SyntheticsCanary#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

