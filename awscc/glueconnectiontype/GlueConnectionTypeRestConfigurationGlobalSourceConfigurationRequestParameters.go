// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnectiontype


type GlueConnectionTypeRestConfigurationGlobalSourceConfigurationRequestParameters struct {
	// A list of allowed values for the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#allowed_values GlueConnectionType#allowed_values}
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// The default value for the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#default_value GlueConnectionType#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// A key name to use when sending this property in API requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#key_override GlueConnectionType#key_override}
	KeyOverride *string `field:"optional" json:"keyOverride" yaml:"keyOverride"`
	// The name of the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#name GlueConnectionType#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies where this property should be included in REST requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#property_location GlueConnectionType#property_location}
	PropertyLocation *string `field:"optional" json:"propertyLocation" yaml:"propertyLocation"`
	// The data type of this property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#property_type GlueConnectionType#property_type}
	PropertyType *string `field:"optional" json:"propertyType" yaml:"propertyType"`
	// Indicates whether the property is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_connection_type#required GlueConnectionType#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

