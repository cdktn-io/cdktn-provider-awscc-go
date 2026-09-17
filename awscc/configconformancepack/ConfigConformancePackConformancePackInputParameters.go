// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconformancepack


type ConfigConformancePackConformancePackInputParameters struct {
	// Key part of key-value pair with value being parameter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/config_conformance_pack#parameter_name ConfigConformancePack#parameter_name}
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// Value part of key-value pair with key being parameter Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/config_conformance_pack#parameter_value ConfigConformancePack#parameter_value}
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

