// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationSlackTransmissionTarget struct {
	// Destination for IncidentResponse agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#incident_response_target DevopsagentAssociation#incident_response_target}
	IncidentResponseTarget *DevopsagentAssociationConfigurationSlackTransmissionTargetIncidentResponseTarget `field:"optional" json:"incidentResponseTarget" yaml:"incidentResponseTarget"`
}

