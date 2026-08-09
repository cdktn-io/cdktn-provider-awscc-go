// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectuser


type ConnectUserUserProficiencies struct {
	// The name of user's proficiency. You must use name of predefined attribute present in the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#attribute_name ConnectUser#attribute_name}
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// The value of user's proficiency. You must use value of predefined attribute present in the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#attribute_value ConnectUser#attribute_value}
	AttributeValue *string `field:"optional" json:"attributeValue" yaml:"attributeValue"`
	// The level of the proficiency. The valid values are 0 to 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_user#level ConnectUser#level}
	Level *float64 `field:"optional" json:"level" yaml:"level"`
}

