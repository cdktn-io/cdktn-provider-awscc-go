// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectdatatableattribute


type ConnectDataTableAttributeValidationEnum struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_data_table_attribute#strict ConnectDataTableAttribute#strict}.
	Strict interface{} `field:"optional" json:"strict" yaml:"strict"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_data_table_attribute#values ConnectDataTableAttribute#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

