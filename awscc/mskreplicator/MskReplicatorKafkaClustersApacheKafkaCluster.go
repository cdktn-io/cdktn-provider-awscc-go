// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClustersApacheKafkaCluster struct {
	// The ID of the Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#apache_kafka_cluster_id MskReplicator#apache_kafka_cluster_id}
	ApacheKafkaClusterId *string `field:"optional" json:"apacheKafkaClusterId" yaml:"apacheKafkaClusterId"`
	// The bootstrap broker string of the Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/msk_replicator#bootstrap_broker_string MskReplicator#bootstrap_broker_string}
	BootstrapBrokerString *string `field:"optional" json:"bootstrapBrokerString" yaml:"bootstrapBrokerString"`
}

