// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicySecurityHubSecurityControlsConfigurationSecurityControlCustomParametersParametersValue struct {
	// A control parameter that is a boolean.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#boolean SecurityhubConfigurationPolicy#boolean}
	Boolean interface{} `field:"optional" json:"boolean" yaml:"boolean"`
	// A control parameter that is a double.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#double SecurityhubConfigurationPolicy#double}
	Double *float64 `field:"optional" json:"double" yaml:"double"`
	// A control parameter that is an enum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#enum SecurityhubConfigurationPolicy#enum}
	Enum *string `field:"optional" json:"enum" yaml:"enum"`
	// A control parameter that is a list of enums.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#enum_list SecurityhubConfigurationPolicy#enum_list}
	EnumList *[]*string `field:"optional" json:"enumList" yaml:"enumList"`
	// A control parameter that is an integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#integer SecurityhubConfigurationPolicy#integer}
	Integer *float64 `field:"optional" json:"integer" yaml:"integer"`
	// A control parameter that is a list of integers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#integer_list SecurityhubConfigurationPolicy#integer_list}
	IntegerList *[]*float64 `field:"optional" json:"integerList" yaml:"integerList"`
	// A control parameter that is a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#string SecurityhubConfigurationPolicy#string}
	String *string `field:"optional" json:"string" yaml:"string"`
	// A control parameter that is a list of strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/securityhub_configuration_policy#string_list SecurityhubConfigurationPolicy#string_list}
	StringList *[]*string `field:"optional" json:"stringList" yaml:"stringList"`
}

