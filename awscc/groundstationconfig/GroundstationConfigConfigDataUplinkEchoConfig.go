// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationconfig


type GroundstationConfigConfigDataUplinkEchoConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/groundstation_config#antenna_uplink_config_arn GroundstationConfig#antenna_uplink_config_arn}.
	AntennaUplinkConfigArn *string `field:"optional" json:"antennaUplinkConfigArn" yaml:"antennaUplinkConfigArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/groundstation_config#enabled GroundstationConfig#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

