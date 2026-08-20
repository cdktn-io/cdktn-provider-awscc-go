// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivschannel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IvsChannelConfig struct {
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
	// Whether the channel is authorized.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#authorized IvsChannel#authorized}
	Authorized interface{} `field:"optional" json:"authorized" yaml:"authorized"`
	// Indicates which content-packaging format is used (MPEG-TS or fMP4).
	//
	// If multitrackInputConfiguration is specified and enabled is true, then containerFormat is required and must be set to FRAGMENTED_MP4. Otherwise, containerFormat may be set to TS or FRAGMENTED_MP4. Default: TS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#container_format IvsChannel#container_format}
	ContainerFormat *string `field:"optional" json:"containerFormat" yaml:"containerFormat"`
	// Whether the channel allows insecure ingest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#insecure_ingest IvsChannel#insecure_ingest}
	InsecureIngest interface{} `field:"optional" json:"insecureIngest" yaml:"insecureIngest"`
	// Channel latency mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#latency_mode IvsChannel#latency_mode}
	LatencyMode *string `field:"optional" json:"latencyMode" yaml:"latencyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#multitrack_input_configuration IvsChannel#multitrack_input_configuration}.
	MultitrackInputConfiguration *IvsChannelMultitrackInputConfiguration `field:"optional" json:"multitrackInputConfiguration" yaml:"multitrackInputConfiguration"`
	// Channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#name IvsChannel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Optional transcode preset for the channel.
	//
	// This is selectable only for ADVANCED_HD and ADVANCED_SD channel types. For those channel types, the default preset is HIGHER_BANDWIDTH_DELIVERY. For other channel types (BASIC and STANDARD), preset is the empty string ("").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#preset IvsChannel#preset}
	Preset *string `field:"optional" json:"preset" yaml:"preset"`
	// Recording Configuration ARN.
	//
	// A value other than an empty string indicates that recording is enabled. Default: "" (recording is disabled).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#recording_configuration_arn IvsChannel#recording_configuration_arn}
	RecordingConfigurationArn *string `field:"optional" json:"recordingConfigurationArn" yaml:"recordingConfigurationArn"`
	// A list of key-value pairs that contain metadata for the asset model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#tags IvsChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Channel type, which determines the allowable resolution and bitrate.
	//
	// If you exceed the allowable resolution or bitrate, the stream probably will disconnect immediately.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ivs_channel#type IvsChannel#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

