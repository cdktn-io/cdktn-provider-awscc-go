// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationconfig


type GroundstationConfigConfigData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#antenna_downlink_config GroundstationConfig#antenna_downlink_config}.
	AntennaDownlinkConfig *GroundstationConfigConfigDataAntennaDownlinkConfig `field:"optional" json:"antennaDownlinkConfig" yaml:"antennaDownlinkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#antenna_downlink_demod_decode_config GroundstationConfig#antenna_downlink_demod_decode_config}.
	AntennaDownlinkDemodDecodeConfig *GroundstationConfigConfigDataAntennaDownlinkDemodDecodeConfig `field:"optional" json:"antennaDownlinkDemodDecodeConfig" yaml:"antennaDownlinkDemodDecodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#antenna_uplink_config GroundstationConfig#antenna_uplink_config}.
	AntennaUplinkConfig *GroundstationConfigConfigDataAntennaUplinkConfig `field:"optional" json:"antennaUplinkConfig" yaml:"antennaUplinkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#dataflow_endpoint_config GroundstationConfig#dataflow_endpoint_config}.
	DataflowEndpointConfig *GroundstationConfigConfigDataDataflowEndpointConfig `field:"optional" json:"dataflowEndpointConfig" yaml:"dataflowEndpointConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#s3_recording_config GroundstationConfig#s3_recording_config}.
	S3RecordingConfig *GroundstationConfigConfigDataS3RecordingConfig `field:"optional" json:"s3RecordingConfig" yaml:"s3RecordingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#telemetry_sink_config GroundstationConfig#telemetry_sink_config}.
	TelemetrySinkConfig *GroundstationConfigConfigDataTelemetrySinkConfig `field:"optional" json:"telemetrySinkConfig" yaml:"telemetrySinkConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#tracking_config GroundstationConfig#tracking_config}.
	TrackingConfig *GroundstationConfigConfigDataTrackingConfig `field:"optional" json:"trackingConfig" yaml:"trackingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/groundstation_config#uplink_echo_config GroundstationConfig#uplink_echo_config}.
	UplinkEchoConfig *GroundstationConfigConfigDataUplinkEchoConfig `field:"optional" json:"uplinkEchoConfig" yaml:"uplinkEchoConfig"`
}

