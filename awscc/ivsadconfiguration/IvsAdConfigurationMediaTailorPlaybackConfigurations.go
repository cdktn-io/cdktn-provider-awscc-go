// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivsadconfiguration


type IvsAdConfigurationMediaTailorPlaybackConfigurations struct {
	// ARN of the customer-created EMT PlaybackConfiguration resource in the same region and account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/ivs_ad_configuration#playback_configuration_arn IvsAdConfiguration#playback_configuration_arn}
	PlaybackConfigurationArn *string `field:"optional" json:"playbackConfigurationArn" yaml:"playbackConfigurationArn"`
}

