// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactscontact


type SsmcontactsContactPlanTargets struct {
	// Information about the contact channel that SSM Incident Manager uses to engage the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmcontacts_contact#channel_target_info SsmcontactsContact#channel_target_info}
	ChannelTargetInfo *SsmcontactsContactPlanTargetsChannelTargetInfo `field:"optional" json:"channelTargetInfo" yaml:"channelTargetInfo"`
	// The contact that SSM Incident Manager is engaging during an incident.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmcontacts_contact#contact_target_info SsmcontactsContact#contact_target_info}
	ContactTargetInfo *SsmcontactsContactPlanTargetsContactTargetInfo `field:"optional" json:"contactTargetInfo" yaml:"contactTargetInfo"`
}

