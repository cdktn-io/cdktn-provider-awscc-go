// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2componentversion


type Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_component_version#devices Greengrassv2ComponentVersion#devices}.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_component_version#memory_size_in_kb Greengrassv2ComponentVersion#memory_size_in_kb}.
	MemorySizeInKb *float64 `field:"optional" json:"memorySizeInKb" yaml:"memorySizeInKb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_component_version#mount_ro_sysfs Greengrassv2ComponentVersion#mount_ro_sysfs}.
	MountRoSysfs interface{} `field:"optional" json:"mountRoSysfs" yaml:"mountRoSysfs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/greengrassv2_component_version#volumes Greengrassv2ComponentVersion#volumes}.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

