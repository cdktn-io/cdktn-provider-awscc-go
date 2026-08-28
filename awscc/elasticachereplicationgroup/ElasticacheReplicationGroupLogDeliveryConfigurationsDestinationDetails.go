// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachereplicationgroup


type ElasticacheReplicationGroupLogDeliveryConfigurationsDestinationDetails struct {
	// The configuration details of the CloudWatch Logs destination.
	//
	// Note that this field is marked as required but only if CloudWatch Logs was chosen as the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_replication_group#cloudwatch_logs_details ElasticacheReplicationGroup#cloudwatch_logs_details}
	CloudwatchLogsDetails *ElasticacheReplicationGroupLogDeliveryConfigurationsDestinationDetailsCloudwatchLogsDetails `field:"optional" json:"cloudwatchLogsDetails" yaml:"cloudwatchLogsDetails"`
	// The configuration details of the Kinesis Data Firehose destination.
	//
	// Note that this field is marked as required but only if Kinesis Data Firehose was chosen as the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_replication_group#kinesis_firehose_details ElasticacheReplicationGroup#kinesis_firehose_details}
	KinesisFirehoseDetails *ElasticacheReplicationGroupLogDeliveryConfigurationsDestinationDetailsKinesisFirehoseDetails `field:"optional" json:"kinesisFirehoseDetails" yaml:"kinesisFirehoseDetails"`
}

