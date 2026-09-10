// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationconfig


type GroundstationConfigConfigDataTelemetrySinkConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/groundstation_config#telemetry_sink_data GroundstationConfig#telemetry_sink_data}.
	TelemetrySinkData *GroundstationConfigConfigDataTelemetrySinkConfigTelemetrySinkData `field:"optional" json:"telemetrySinkData" yaml:"telemetrySinkData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/groundstation_config#telemetry_sink_type GroundstationConfig#telemetry_sink_type}.
	TelemetrySinkType *string `field:"optional" json:"telemetrySinkType" yaml:"telemetrySinkType"`
}

