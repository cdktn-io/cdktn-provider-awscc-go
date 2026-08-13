// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator


type MskReplicatorKafkaClusters struct {
	// Details of an Amazon MSK cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_replicator#amazon_msk_cluster MskReplicator#amazon_msk_cluster}
	AmazonMskCluster *MskReplicatorKafkaClustersAmazonMskCluster `field:"optional" json:"amazonMskCluster" yaml:"amazonMskCluster"`
	// Details of an Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_replicator#apache_kafka_cluster MskReplicator#apache_kafka_cluster}
	ApacheKafkaCluster *MskReplicatorKafkaClustersApacheKafkaCluster `field:"optional" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
	// Details of the client authentication used by the Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_replicator#client_authentication MskReplicator#client_authentication}
	ClientAuthentication *MskReplicatorKafkaClustersClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// Details of encryption in transit to the Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_replicator#encryption_in_transit MskReplicator#encryption_in_transit}
	EncryptionInTransit *MskReplicatorKafkaClustersEncryptionInTransit `field:"optional" json:"encryptionInTransit" yaml:"encryptionInTransit"`
	// Details of an Amazon VPC which has network connectivity to the Apache Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/msk_replicator#vpc_config MskReplicator#vpc_config}
	VpcConfig *MskReplicatorKafkaClustersVpcConfig `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

