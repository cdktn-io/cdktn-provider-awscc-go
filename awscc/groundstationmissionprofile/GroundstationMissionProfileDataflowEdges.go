// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package groundstationmissionprofile


type GroundstationMissionProfileDataflowEdges struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/groundstation_mission_profile#destination GroundstationMissionProfile#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/groundstation_mission_profile#source GroundstationMissionProfile#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

