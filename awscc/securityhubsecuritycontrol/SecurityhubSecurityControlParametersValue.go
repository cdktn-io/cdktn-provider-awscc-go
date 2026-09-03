// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubsecuritycontrol


type SecurityhubSecurityControlParametersValue struct {
	// A control parameter that is a boolean.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#boolean SecurityhubSecurityControl#boolean}
	Boolean interface{} `field:"optional" json:"boolean" yaml:"boolean"`
	// A control parameter that is a double.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#double SecurityhubSecurityControl#double}
	Double *float64 `field:"optional" json:"double" yaml:"double"`
	// A control parameter that is a enum.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#enum SecurityhubSecurityControl#enum}
	Enum *string `field:"optional" json:"enum" yaml:"enum"`
	// A control parameter that is a list of enums.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#enum_list SecurityhubSecurityControl#enum_list}
	EnumList *[]*string `field:"optional" json:"enumList" yaml:"enumList"`
	// A control parameter that is a integer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#integer SecurityhubSecurityControl#integer}
	Integer *float64 `field:"optional" json:"integer" yaml:"integer"`
	// A control parameter that is a list of integers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#integer_list SecurityhubSecurityControl#integer_list}
	IntegerList *[]*float64 `field:"optional" json:"integerList" yaml:"integerList"`
	// A control parameter that is a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#string SecurityhubSecurityControl#string}
	String *string `field:"optional" json:"string" yaml:"string"`
	// A control parameter that is a list of strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#string_list SecurityhubSecurityControl#string_list}
	StringList *[]*string `field:"optional" json:"stringList" yaml:"stringList"`
}

