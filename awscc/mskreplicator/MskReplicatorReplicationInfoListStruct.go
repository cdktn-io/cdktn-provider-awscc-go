// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorReplicationInfoListStruct struct {
	// Configuration relating to consumer group replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_replicator#consumer_group_replication MskReplicator#consumer_group_replication}
	ConsumerGroupReplication *MskReplicatorReplicationInfoListConsumerGroupReplication `field:"required" json:"consumerGroupReplication" yaml:"consumerGroupReplication"`
	// The type of compression to use writing records to target Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_replicator#target_compression_type MskReplicator#target_compression_type}
	TargetCompressionType *string `field:"required" json:"targetCompressionType" yaml:"targetCompressionType"`
	// Configuration relating to topic replication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_replicator#topic_replication MskReplicator#topic_replication}
	TopicReplication *MskReplicatorReplicationInfoListTopicReplication `field:"required" json:"topicReplication" yaml:"topicReplication"`
	// Amazon Resource Name of the source Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_replicator#source_kafka_cluster_arn MskReplicator#source_kafka_cluster_arn}
	SourceKafkaClusterArn *string `field:"optional" json:"sourceKafkaClusterArn" yaml:"sourceKafkaClusterArn"`
	// The ID of the source Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_replicator#source_kafka_cluster_id MskReplicator#source_kafka_cluster_id}
	SourceKafkaClusterId *string `field:"optional" json:"sourceKafkaClusterId" yaml:"sourceKafkaClusterId"`
	// Amazon Resource Name of the target Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_replicator#target_kafka_cluster_arn MskReplicator#target_kafka_cluster_arn}
	TargetKafkaClusterArn *string `field:"optional" json:"targetKafkaClusterArn" yaml:"targetKafkaClusterArn"`
	// The ID of the target Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_replicator#target_kafka_cluster_id MskReplicator#target_kafka_cluster_id}
	TargetKafkaClusterId *string `field:"optional" json:"targetKafkaClusterId" yaml:"targetKafkaClusterId"`
}

