// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectdatatableattribute


type ConnectDataTableAttributeValidation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#enum ConnectDataTableAttribute#enum}.
	Enum *ConnectDataTableAttributeValidationEnum `field:"optional" json:"enum" yaml:"enum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#exclusive_maximum ConnectDataTableAttribute#exclusive_maximum}.
	ExclusiveMaximum *float64 `field:"optional" json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#exclusive_minimum ConnectDataTableAttribute#exclusive_minimum}.
	ExclusiveMinimum *float64 `field:"optional" json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#maximum ConnectDataTableAttribute#maximum}.
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#max_length ConnectDataTableAttribute#max_length}.
	MaxLength *float64 `field:"optional" json:"maxLength" yaml:"maxLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#max_values ConnectDataTableAttribute#max_values}.
	MaxValues *float64 `field:"optional" json:"maxValues" yaml:"maxValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#minimum ConnectDataTableAttribute#minimum}.
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#min_length ConnectDataTableAttribute#min_length}.
	MinLength *float64 `field:"optional" json:"minLength" yaml:"minLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#min_values ConnectDataTableAttribute#min_values}.
	MinValues *float64 `field:"optional" json:"minValues" yaml:"minValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_data_table_attribute#multiple_of ConnectDataTableAttribute#multiple_of}.
	MultipleOf *float64 `field:"optional" json:"multipleOf" yaml:"multipleOf"`
}

