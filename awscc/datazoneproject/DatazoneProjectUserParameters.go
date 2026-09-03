// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneproject


type DatazoneProjectUserParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_project#environment_configuration_name DatazoneProject#environment_configuration_name}.
	EnvironmentConfigurationName *string `field:"optional" json:"environmentConfigurationName" yaml:"environmentConfigurationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_project#environment_id DatazoneProject#environment_id}.
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_project#environment_parameters DatazoneProject#environment_parameters}.
	EnvironmentParameters interface{} `field:"optional" json:"environmentParameters" yaml:"environmentParameters"`
}

