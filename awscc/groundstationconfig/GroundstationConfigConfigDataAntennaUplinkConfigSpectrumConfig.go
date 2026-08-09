// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationconfig


type GroundstationConfigConfigDataAntennaUplinkConfigSpectrumConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/groundstation_config#center_frequency GroundstationConfig#center_frequency}.
	CenterFrequency *GroundstationConfigConfigDataAntennaUplinkConfigSpectrumConfigCenterFrequency `field:"optional" json:"centerFrequency" yaml:"centerFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/groundstation_config#polarization GroundstationConfig#polarization}.
	Polarization *string `field:"optional" json:"polarization" yaml:"polarization"`
}

