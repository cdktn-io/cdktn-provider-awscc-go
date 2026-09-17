// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachecachecluster


type ElasticacheCacheClusterLogDeliveryConfigurationsDestinationDetailsKinesisFirehoseDetails struct {
	// The name of the Kinesis Data Firehose delivery stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticache_cache_cluster#delivery_stream ElasticacheCacheCluster#delivery_stream}
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}

