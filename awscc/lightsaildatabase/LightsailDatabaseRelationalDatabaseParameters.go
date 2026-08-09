// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildatabase


type LightsailDatabaseRelationalDatabaseParameters struct {
	// Specifies the valid range of values for the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_database#allowed_values LightsailDatabase#allowed_values}
	AllowedValues *string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// Indicates when parameter updates are applied. Can be immediate or pending-reboot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_database#apply_method LightsailDatabase#apply_method}
	ApplyMethod *string `field:"optional" json:"applyMethod" yaml:"applyMethod"`
	// Specifies the engine-specific parameter type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_database#apply_type LightsailDatabase#apply_type}
	ApplyType *string `field:"optional" json:"applyType" yaml:"applyType"`
	// Specifies the valid data type for the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_database#data_type LightsailDatabase#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Provides a description of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_database#description LightsailDatabase#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A Boolean value indicating whether the parameter can be modified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_database#is_modifiable LightsailDatabase#is_modifiable}
	IsModifiable interface{} `field:"optional" json:"isModifiable" yaml:"isModifiable"`
	// Specifies the name of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_database#parameter_name LightsailDatabase#parameter_name}
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Specifies the value of the parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lightsail_database#parameter_value LightsailDatabase#parameter_value}
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

