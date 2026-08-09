// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsgurunotificationchannel


type DevopsguruNotificationChannelConfigFilters struct {
	// DevOps Guru message types to filter for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsguru_notification_channel#message_types DevopsguruNotificationChannel#message_types}
	MessageTypes *[]*string `field:"optional" json:"messageTypes" yaml:"messageTypes"`
	// DevOps Guru insight severities to filter for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsguru_notification_channel#severities DevopsguruNotificationChannel#severities}
	Severities *[]*string `field:"optional" json:"severities" yaml:"severities"`
}

