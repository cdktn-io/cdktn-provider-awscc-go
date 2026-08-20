// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachecachecluster


type ElasticacheCacheClusterLogDeliveryConfigurations struct {
	// Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_cache_cluster#destination_details ElasticacheCacheCluster#destination_details}
	DestinationDetails *ElasticacheCacheClusterLogDeliveryConfigurationsDestinationDetails `field:"optional" json:"destinationDetails" yaml:"destinationDetails"`
	// Specify either CloudWatch Logs or Kinesis Data Firehose as the destination type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_cache_cluster#destination_type ElasticacheCacheCluster#destination_type}
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Valid values are either json or text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_cache_cluster#log_format ElasticacheCacheCluster#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// Valid value is either slow-log, which refers to slow-log or engine-log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_cache_cluster#log_type ElasticacheCacheCluster#log_type}
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
}

