// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigurationrecorder

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigConfigurationRecorderConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Amazon Resource Name (ARN) of the role that allows AWS Config to read S3 objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/config_configuration_recorder#role_arn ConfigConfigurationRecorder#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the configuration recorder.
	//
	// By default, AWS Config assigns the name "default" when creating the configuration recorder. To change the configuration recorder name, you must use the DeleteConfigurationRecorder action to delete your current configuration recorder, and then you must use the PutConfigurationRecorder command to create a configuration recorder that has the desired name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/config_configuration_recorder#name ConfigConfigurationRecorder#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies which resource types are in scope for the configuration recorder to record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/config_configuration_recorder#recording_group ConfigConfigurationRecorder#recording_group}
	RecordingGroup *ConfigConfigurationRecorderRecordingGroup `field:"optional" json:"recordingGroup" yaml:"recordingGroup"`
	// Specifies the default recording frequency for the configuration recorder.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/config_configuration_recorder#recording_mode ConfigConfigurationRecorder#recording_mode}
	RecordingMode *ConfigConfigurationRecorderRecordingMode `field:"optional" json:"recordingMode" yaml:"recordingMode"`
	// Defaults to 'true'.
	//
	// Controls whether the recorder starts recording after the create operation. Set this to 'false' for development and testing purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/config_configuration_recorder#started_on_create ConfigConfigurationRecorder#started_on_create}
	StartedOnCreate interface{} `field:"optional" json:"startedOnCreate" yaml:"startedOnCreate"`
}

