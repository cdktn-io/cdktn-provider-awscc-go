// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsgurunotificationchannel


type DevopsguruNotificationChannelConfigA struct {
	// Information about filters of a notification channel configured in DevOpsGuru to filter for insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsguru_notification_channel#filters DevopsguruNotificationChannel#filters}
	Filters *DevopsguruNotificationChannelConfigFilters `field:"optional" json:"filters" yaml:"filters"`
	// Information about a notification channel configured in DevOps Guru to send notifications when insights are created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/devopsguru_notification_channel#sns DevopsguruNotificationChannel#sns}
	Sns *DevopsguruNotificationChannelConfigSns `field:"optional" json:"sns" yaml:"sns"`
}

