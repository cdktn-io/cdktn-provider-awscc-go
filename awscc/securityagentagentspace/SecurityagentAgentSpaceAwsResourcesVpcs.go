// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceAwsResourcesVpcs struct {
	// List of security group ARNs in the customer VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityagent_agent_space#security_group_arns SecurityagentAgentSpace#security_group_arns}
	SecurityGroupArns *[]*string `field:"optional" json:"securityGroupArns" yaml:"securityGroupArns"`
	// List of subnet ARNs in the customer VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityagent_agent_space#subnet_arns SecurityagentAgentSpace#subnet_arns}
	SubnetArns *[]*string `field:"optional" json:"subnetArns" yaml:"subnetArns"`
	// ARN of the customer VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityagent_agent_space#vpc_arn SecurityagentAgentSpace#vpc_arn}
	VpcArn *string `field:"optional" json:"vpcArn" yaml:"vpcArn"`
}

