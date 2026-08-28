// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecttasktemplate


type ConnectTaskTemplateConstraints struct {
	// The list of the task template's invisible fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_task_template#invisible_fields ConnectTaskTemplate#invisible_fields}
	InvisibleFields interface{} `field:"optional" json:"invisibleFields" yaml:"invisibleFields"`
	// The list of the task template's read only fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_task_template#read_only_fields ConnectTaskTemplate#read_only_fields}
	ReadOnlyFields interface{} `field:"optional" json:"readOnlyFields" yaml:"readOnlyFields"`
	// The list of the task template's required fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_task_template#required_fields ConnectTaskTemplate#required_fields}
	RequiredFields interface{} `field:"optional" json:"requiredFields" yaml:"requiredFields"`
}

