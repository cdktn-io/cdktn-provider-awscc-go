// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2applicationstatuscheck


type Ec2ApplicationStatusCheckTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#key Ec2ApplicationStatusCheck#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_application_status_check#value Ec2ApplicationStatusCheck#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

