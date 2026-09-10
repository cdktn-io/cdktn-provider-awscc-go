// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2flowlog


type Ec2FlowLogTagFieldSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#resource_type Ec2FlowLog#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_flow_log#tag_keys Ec2FlowLog#tag_keys}.
	TagKeys *[]*string `field:"optional" json:"tagKeys" yaml:"tagKeys"`
}

