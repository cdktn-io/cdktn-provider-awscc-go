// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nimblestudiostudiocomponent


type NimblestudioStudioComponentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/nimblestudio_studio_component#active_directory_configuration NimblestudioStudioComponent#active_directory_configuration}.
	ActiveDirectoryConfiguration *NimblestudioStudioComponentConfigurationActiveDirectoryConfiguration `field:"optional" json:"activeDirectoryConfiguration" yaml:"activeDirectoryConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/nimblestudio_studio_component#compute_farm_configuration NimblestudioStudioComponent#compute_farm_configuration}.
	ComputeFarmConfiguration *NimblestudioStudioComponentConfigurationComputeFarmConfiguration `field:"optional" json:"computeFarmConfiguration" yaml:"computeFarmConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/nimblestudio_studio_component#license_service_configuration NimblestudioStudioComponent#license_service_configuration}.
	LicenseServiceConfiguration *NimblestudioStudioComponentConfigurationLicenseServiceConfiguration `field:"optional" json:"licenseServiceConfiguration" yaml:"licenseServiceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/nimblestudio_studio_component#shared_file_system_configuration NimblestudioStudioComponent#shared_file_system_configuration}.
	SharedFileSystemConfiguration *NimblestudioStudioComponentConfigurationSharedFileSystemConfiguration `field:"optional" json:"sharedFileSystemConfiguration" yaml:"sharedFileSystemConfiguration"`
}

