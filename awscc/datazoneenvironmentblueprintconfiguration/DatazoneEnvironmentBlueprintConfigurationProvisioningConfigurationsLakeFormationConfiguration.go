// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneenvironmentblueprintconfiguration


type DatazoneEnvironmentBlueprintConfigurationProvisioningConfigurationsLakeFormationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_environment_blueprint_configuration#location_registration_exclude_s3_locations DatazoneEnvironmentBlueprintConfiguration#location_registration_exclude_s3_locations}.
	LocationRegistrationExcludeS3Locations *[]*string `field:"optional" json:"locationRegistrationExcludeS3Locations" yaml:"locationRegistrationExcludeS3Locations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datazone_environment_blueprint_configuration#location_registration_role DatazoneEnvironmentBlueprintConfiguration#location_registration_role}.
	LocationRegistrationRole *string `field:"optional" json:"locationRegistrationRole" yaml:"locationRegistrationRole"`
}

