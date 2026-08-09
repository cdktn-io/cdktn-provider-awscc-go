// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsecurityprofile


type IotSecurityProfileBehaviorsMetricDimension struct {
	// A unique identifier for the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_security_profile#dimension_name IotSecurityProfile#dimension_name}
	DimensionName *string `field:"optional" json:"dimensionName" yaml:"dimensionName"`
	// Defines how the dimensionValues of a dimension are interpreted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_security_profile#operator IotSecurityProfile#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

