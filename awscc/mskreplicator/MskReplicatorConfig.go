// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskreplicator

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MskReplicatorConfig struct {
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
	// Specifies a list of Kafka clusters which are targets of the replicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#kafka_clusters MskReplicator#kafka_clusters}
	KafkaClusters interface{} `field:"required" json:"kafkaClusters" yaml:"kafkaClusters"`
	// A list of replication configurations, where each configuration targets a given source cluster to target cluster replication flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#replication_info_list MskReplicator#replication_info_list}
	ReplicationInfoList interface{} `field:"required" json:"replicationInfoList" yaml:"replicationInfoList"`
	// The name of the replicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#replicator_name MskReplicator#replicator_name}
	ReplicatorName *string `field:"required" json:"replicatorName" yaml:"replicatorName"`
	// The Amazon Resource Name (ARN) of the IAM role used by the replicator to access external resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#service_execution_role_arn MskReplicator#service_execution_role_arn}
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// A summary description of the replicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#description MskReplicator#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Configuration for log delivery for the replicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#log_delivery MskReplicator#log_delivery}
	LogDelivery *MskReplicatorLogDelivery `field:"optional" json:"logDelivery" yaml:"logDelivery"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_replicator#tags MskReplicator#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

