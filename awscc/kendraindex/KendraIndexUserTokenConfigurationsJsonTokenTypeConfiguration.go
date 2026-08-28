// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraindex


type KendraIndexUserTokenConfigurationsJsonTokenTypeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_index#group_attribute_field KendraIndex#group_attribute_field}.
	GroupAttributeField *string `field:"optional" json:"groupAttributeField" yaml:"groupAttributeField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_index#user_name_attribute_field KendraIndex#user_name_attribute_field}.
	UserNameAttributeField *string `field:"optional" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

