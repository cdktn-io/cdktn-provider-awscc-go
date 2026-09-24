// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationconfig


type GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/groundstation_config#bandwidth GroundstationConfig#bandwidth}.
	Bandwidth *GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigBandwidth `field:"optional" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/groundstation_config#center_frequency GroundstationConfig#center_frequency}.
	CenterFrequency *GroundstationConfigConfigDataAntennaDownlinkConfigSpectrumConfigCenterFrequency `field:"optional" json:"centerFrequency" yaml:"centerFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/groundstation_config#polarization GroundstationConfig#polarization}.
	Polarization *string `field:"optional" json:"polarization" yaml:"polarization"`
}

