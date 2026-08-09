// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelTopicConfigurationListRecordConverter struct {
	// Value converter for topic data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#value_converter MskChannel#value_converter}
	ValueConverter *string `field:"required" json:"valueConverter" yaml:"valueConverter"`
}

