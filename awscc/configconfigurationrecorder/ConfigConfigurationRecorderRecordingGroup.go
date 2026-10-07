// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigurationrecorder


type ConfigConfigurationRecorderRecordingGroup struct {
	// Specifies whether AWS Config records configuration changes for all supported resource types, excluding the global IAM resource types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/config_configuration_recorder#all_supported ConfigConfigurationRecorder#all_supported}
	AllSupported interface{} `field:"optional" json:"allSupported" yaml:"allSupported"`
	// An object that specifies how AWS Config excludes resource types from being recorded by the configuration recorder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/config_configuration_recorder#exclusion_by_resource_types ConfigConfigurationRecorder#exclusion_by_resource_types}
	ExclusionByResourceTypes *ConfigConfigurationRecorderRecordingGroupExclusionByResourceTypes `field:"optional" json:"exclusionByResourceTypes" yaml:"exclusionByResourceTypes"`
	// This option is a bundle which only applies to the global IAM resource types: IAM users, groups, roles, and customer managed policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/config_configuration_recorder#include_global_resource_types ConfigConfigurationRecorder#include_global_resource_types}
	IncludeGlobalResourceTypes interface{} `field:"optional" json:"includeGlobalResourceTypes" yaml:"includeGlobalResourceTypes"`
	// An object that specifies the recording strategy for the configuration recorder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/config_configuration_recorder#recording_strategy ConfigConfigurationRecorder#recording_strategy}
	RecordingStrategy *ConfigConfigurationRecorderRecordingGroupRecordingStrategy `field:"optional" json:"recordingStrategy" yaml:"recordingStrategy"`
	// A comma-separated list that specifies which resource types AWS Config records.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/config_configuration_recorder#resource_types ConfigConfigurationRecorder#resource_types}
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

