// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisstream

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisStreamConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The final list of shard-level metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesis_stream#desired_shard_level_metrics KinesisStream#desired_shard_level_metrics}
	DesiredShardLevelMetrics *[]*string `field:"optional" json:"desiredShardLevelMetrics" yaml:"desiredShardLevelMetrics"`
	// Maximum size of a data record in KiB allowed to be put into Kinesis stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesis_stream#max_record_size_in_ki_b KinesisStream#max_record_size_in_ki_b}
	MaxRecordSizeInKiB *float64 `field:"optional" json:"maxRecordSizeInKiB" yaml:"maxRecordSizeInKiB"`
	// The name of the Kinesis stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesis_stream#name KinesisStream#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The number of hours for the data records that are stored in shards to remain accessible.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesis_stream#retention_period_hours KinesisStream#retention_period_hours}
	RetentionPeriodHours *float64 `field:"optional" json:"retentionPeriodHours" yaml:"retentionPeriodHours"`
	// The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesis_stream#shard_count KinesisStream#shard_count}
	ShardCount *float64 `field:"optional" json:"shardCount" yaml:"shardCount"`
	// When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesis_stream#stream_encryption KinesisStream#stream_encryption}
	StreamEncryption *KinesisStreamStreamEncryption `field:"optional" json:"streamEncryption" yaml:"streamEncryption"`
	// The mode in which the stream is running.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesis_stream#stream_mode_details KinesisStream#stream_mode_details}
	StreamModeDetails *KinesisStreamStreamModeDetails `field:"optional" json:"streamModeDetails" yaml:"streamModeDetails"`
	// An arbitrary set of tags (key-value pairs) to associate with the Kinesis stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesis_stream#tags KinesisStream#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Target warm throughput in MiB/s for the stream. This property can ONLY be set when StreamMode is ON_DEMAND.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kinesis_stream#warm_throughput_mi_bps KinesisStream#warm_throughput_mi_bps}
	WarmThroughputMiBps *float64 `field:"optional" json:"warmThroughputMiBps" yaml:"warmThroughputMiBps"`
}

