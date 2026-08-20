// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectpredefinedattribute


type ConnectPredefinedAttributeAttributeConfiguration struct {
	// Enables customers to enforce strict validation on the specific values that this predefined attribute can hold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_predefined_attribute#enable_value_validation_on_association ConnectPredefinedAttribute#enable_value_validation_on_association}
	EnableValueValidationOnAssociation interface{} `field:"optional" json:"enableValueValidationOnAssociation" yaml:"enableValueValidationOnAssociation"`
	// Allows the predefined attribute to show up and be managed in the Amazon Connect UI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/connect_predefined_attribute#is_read_only ConnectPredefinedAttribute#is_read_only}
	IsReadOnly interface{} `field:"optional" json:"isReadOnly" yaml:"isReadOnly"`
}

