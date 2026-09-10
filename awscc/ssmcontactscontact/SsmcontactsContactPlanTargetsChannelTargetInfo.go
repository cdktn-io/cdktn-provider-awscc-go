// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactscontact


type SsmcontactsContactPlanTargetsChannelTargetInfo struct {
	// The Amazon Resource Name (ARN) of the contact channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmcontacts_contact#channel_id SsmcontactsContact#channel_id}
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// The number of minutes to wait to retry sending engagement in the case the engagement initially fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ssmcontacts_contact#retry_interval_in_minutes SsmcontactsContact#retry_interval_in_minutes}
	RetryIntervalInMinutes *float64 `field:"optional" json:"retryIntervalInMinutes" yaml:"retryIntervalInMinutes"`
}

