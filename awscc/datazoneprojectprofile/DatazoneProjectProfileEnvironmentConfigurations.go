// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneprojectprofile


type DatazoneProjectProfileEnvironmentConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_project_profile#aws_account DatazoneProjectProfile#aws_account}.
	AwsAccount *DatazoneProjectProfileEnvironmentConfigurationsAwsAccount `field:"optional" json:"awsAccount" yaml:"awsAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_project_profile#aws_region DatazoneProjectProfile#aws_region}.
	AwsRegion *DatazoneProjectProfileEnvironmentConfigurationsAwsRegion `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_project_profile#configuration_parameters DatazoneProjectProfile#configuration_parameters}.
	ConfigurationParameters *DatazoneProjectProfileEnvironmentConfigurationsConfigurationParameters `field:"optional" json:"configurationParameters" yaml:"configurationParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_project_profile#deployment_mode DatazoneProjectProfile#deployment_mode}.
	DeploymentMode *string `field:"optional" json:"deploymentMode" yaml:"deploymentMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_project_profile#deployment_order DatazoneProjectProfile#deployment_order}.
	DeploymentOrder *float64 `field:"optional" json:"deploymentOrder" yaml:"deploymentOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_project_profile#description DatazoneProjectProfile#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_project_profile#environment_blueprint_id DatazoneProjectProfile#environment_blueprint_id}.
	EnvironmentBlueprintId *string `field:"optional" json:"environmentBlueprintId" yaml:"environmentBlueprintId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_project_profile#environment_configuration_id DatazoneProjectProfile#environment_configuration_id}.
	EnvironmentConfigurationId *string `field:"optional" json:"environmentConfigurationId" yaml:"environmentConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datazone_project_profile#name DatazoneProjectProfile#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

