// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream


type KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceAuthenticationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesisfirehose_delivery_stream#secrets_manager_configuration KinesisfirehoseDeliveryStream#secrets_manager_configuration}.
	SecretsManagerConfiguration *KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationDatabaseSourceAuthenticationConfigurationSecretsManagerConfiguration `field:"optional" json:"secretsManagerConfiguration" yaml:"secretsManagerConfiguration"`
}

