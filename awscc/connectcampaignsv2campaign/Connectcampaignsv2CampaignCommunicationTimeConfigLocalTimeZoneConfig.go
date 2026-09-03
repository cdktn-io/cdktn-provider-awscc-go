// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignCommunicationTimeConfigLocalTimeZoneConfig struct {
	// Time Zone Id in the IANA format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connectcampaignsv2_campaign#default_time_zone Connectcampaignsv2Campaign#default_time_zone}
	DefaultTimeZone *string `field:"optional" json:"defaultTimeZone" yaml:"defaultTimeZone"`
	// Local TimeZone Detection method list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connectcampaignsv2_campaign#local_time_zone_detection Connectcampaignsv2Campaign#local_time_zone_detection}
	LocalTimeZoneDetection *[]*string `field:"optional" json:"localTimeZoneDetection" yaml:"localTimeZoneDetection"`
	// Local TimeZone Detection scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connectcampaignsv2_campaign#local_time_zone_detection_scope Connectcampaignsv2Campaign#local_time_zone_detection_scope}
	LocalTimeZoneDetectionScope *string `field:"optional" json:"localTimeZoneDetectionScope" yaml:"localTimeZoneDetectionScope"`
}

