// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationconfig


type GroundstationConfigConfigDataAntennaUplinkConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_config#spectrum_config GroundstationConfig#spectrum_config}.
	SpectrumConfig *GroundstationConfigConfigDataAntennaUplinkConfigSpectrumConfig `field:"optional" json:"spectrumConfig" yaml:"spectrumConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_config#target_eirp GroundstationConfig#target_eirp}.
	TargetEirp *GroundstationConfigConfigDataAntennaUplinkConfigTargetEirp `field:"optional" json:"targetEirp" yaml:"targetEirp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/groundstation_config#transmit_disabled GroundstationConfig#transmit_disabled}.
	TransmitDisabled interface{} `field:"optional" json:"transmitDisabled" yaml:"transmitDisabled"`
}

