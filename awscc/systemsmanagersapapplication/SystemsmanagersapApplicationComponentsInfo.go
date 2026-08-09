// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package systemsmanagersapapplication


type SystemsmanagersapApplicationComponentsInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/systemsmanagersap_application#component_type SystemsmanagersapApplication#component_type}.
	ComponentType *string `field:"optional" json:"componentType" yaml:"componentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/systemsmanagersap_application#ec_2_instance_id SystemsmanagersapApplication#ec_2_instance_id}.
	Ec2InstanceId *string `field:"optional" json:"ec2InstanceId" yaml:"ec2InstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/systemsmanagersap_application#sid SystemsmanagersapApplication#sid}.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}

