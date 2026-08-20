// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelLoggingInfoFirehose struct {
	// The Firehose delivery stream for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_channel#delivery_stream MskChannel#delivery_stream}
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
	// Whether Firehose logging is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_channel#enabled MskChannel#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

