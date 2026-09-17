// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationmissionprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GroundstationMissionProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#dataflow_edges GroundstationMissionProfile#dataflow_edges}.
	DataflowEdges interface{} `field:"required" json:"dataflowEdges" yaml:"dataflowEdges"`
	// Visibilities with shorter duration than the specified minimum viable contact duration will be ignored when searching for available contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#minimum_viable_contact_duration_seconds GroundstationMissionProfile#minimum_viable_contact_duration_seconds}
	MinimumViableContactDurationSeconds *float64 `field:"required" json:"minimumViableContactDurationSeconds" yaml:"minimumViableContactDurationSeconds"`
	// A name used to identify a mission profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#name GroundstationMissionProfile#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#tracking_config_arn GroundstationMissionProfile#tracking_config_arn}.
	TrackingConfigArn *string `field:"required" json:"trackingConfigArn" yaml:"trackingConfigArn"`
	// Post-pass time needed after the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#contact_post_pass_duration_seconds GroundstationMissionProfile#contact_post_pass_duration_seconds}
	ContactPostPassDurationSeconds *float64 `field:"optional" json:"contactPostPassDurationSeconds" yaml:"contactPostPassDurationSeconds"`
	// Pre-pass time needed before the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#contact_pre_pass_duration_seconds GroundstationMissionProfile#contact_pre_pass_duration_seconds}
	ContactPrePassDurationSeconds *float64 `field:"optional" json:"contactPrePassDurationSeconds" yaml:"contactPrePassDurationSeconds"`
	// The ARN of a KMS Key used for encrypting data during transmission from the source to destination locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#streams_kms_key GroundstationMissionProfile#streams_kms_key}
	StreamsKmsKey *GroundstationMissionProfileStreamsKmsKey `field:"optional" json:"streamsKmsKey" yaml:"streamsKmsKey"`
	// The ARN of the KMS Key or Alias Key role used to define permissions on KMS Key usage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#streams_kms_role GroundstationMissionProfile#streams_kms_role}
	StreamsKmsRole *string `field:"optional" json:"streamsKmsRole" yaml:"streamsKmsRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#tags GroundstationMissionProfile#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// ARN of a Config resource of type TelemetrySinkConfig used for telemetry data sink configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_mission_profile#telemetry_sink_config_arn GroundstationMissionProfile#telemetry_sink_config_arn}
	TelemetrySinkConfigArn *string `field:"optional" json:"telemetrySinkConfigArn" yaml:"telemetrySinkConfigArn"`
}

