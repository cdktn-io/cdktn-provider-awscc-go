// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationTables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kinesisfirehose_delivery_stream#exclude KinesisfirehoseDeliveryStream#exclude}.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kinesisfirehose_delivery_stream#include KinesisfirehoseDeliveryStream#include}.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}

