// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsplan


type SsmcontactsPlanStagesTargets struct {
	// Information about the contact channel that SSM Incident Manager uses to engage the contact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_plan#channel_target_info SsmcontactsPlan#channel_target_info}
	ChannelTargetInfo *SsmcontactsPlanStagesTargetsChannelTargetInfo `field:"optional" json:"channelTargetInfo" yaml:"channelTargetInfo"`
	// The contact that SSM Incident Manager is engaging during an incident.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/ssmcontacts_plan#contact_target_info SsmcontactsPlan#contact_target_info}
	ContactTargetInfo *SsmcontactsPlanStagesTargetsContactTargetInfo `field:"optional" json:"contactTargetInfo" yaml:"contactTargetInfo"`
}

