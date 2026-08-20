// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachereplicationgroup


type ElasticacheReplicationGroupLogDeliveryConfigurationsDestinationDetailsKinesisFirehoseDetails struct {
	// The name of the Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_replication_group#delivery_stream ElasticacheReplicationGroup#delivery_stream}
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

