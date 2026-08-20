// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorLogDeliveryReplicatorLogDeliveryFirehose struct {
	// The Firehose delivery stream that is the destination for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#delivery_stream MskReplicator#delivery_stream}
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
	// Whether log delivery to Firehose is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#enabled MskReplicator#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

