// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneprojectprofile


type DatazoneProjectProfileEnvironmentConfigurationsConfigurationParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_project_profile#parameter_overrides DatazoneProjectProfile#parameter_overrides}.
	ParameterOverrides interface{} `field:"optional" json:"parameterOverrides" yaml:"parameterOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_project_profile#resolved_parameters DatazoneProjectProfile#resolved_parameters}.
	ResolvedParameters interface{} `field:"optional" json:"resolvedParameters" yaml:"resolvedParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_project_profile#ssm_path DatazoneProjectProfile#ssm_path}.
	SsmPath *string `field:"optional" json:"ssmPath" yaml:"ssmPath"`
}

