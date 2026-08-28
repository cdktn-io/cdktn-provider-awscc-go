// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package msktopic

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MskTopicConfig struct {
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
	// The Amazon Resource Name (ARN) of the MSK cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_topic#cluster_arn MskTopic#cluster_arn}
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// The number of partitions for the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_topic#partition_count MskTopic#partition_count}
	PartitionCount *float64 `field:"required" json:"partitionCount" yaml:"partitionCount"`
	// The replication factor for the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_topic#replication_factor MskTopic#replication_factor}
	ReplicationFactor *float64 `field:"required" json:"replicationFactor" yaml:"replicationFactor"`
	// The name of the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_topic#topic_name MskTopic#topic_name}
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
	// Base64 encoded configuration properties of the topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/msk_topic#configs MskTopic#configs}
	Configs *string `field:"optional" json:"configs" yaml:"configs"`
}

