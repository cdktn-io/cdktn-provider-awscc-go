// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamDirectPutSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/kinesisfirehose_delivery_stream#throughput_hint_in_m_bs KinesisfirehoseDeliveryStream#throughput_hint_in_m_bs}.
	ThroughputHintInMBs *float64 `field:"optional" json:"throughputHintInMBs" yaml:"throughputHintInMBs"`
}

