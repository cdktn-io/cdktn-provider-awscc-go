// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachereplicationgroup


type ElasticacheReplicationGroupLogDeliveryConfigurations struct {
	// Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticache_replication_group#destination_details ElasticacheReplicationGroup#destination_details}
	DestinationDetails *ElasticacheReplicationGroupLogDeliveryConfigurationsDestinationDetails `field:"optional" json:"destinationDetails" yaml:"destinationDetails"`
	// Specify either CloudWatch Logs or Kinesis Data Firehose as the destination type. Valid values are either cloudwatch-logs or kinesis-firehose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticache_replication_group#destination_type ElasticacheReplicationGroup#destination_type}
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// Valid values are either json or text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticache_replication_group#log_format ElasticacheReplicationGroup#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// Valid value is either slow-log, which refers to slow-log or engine-log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/elasticache_replication_group#log_type ElasticacheReplicationGroup#log_type}
	LogType *string `field:"optional" json:"logType" yaml:"logType"`
}

