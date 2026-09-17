// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskcluster

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MskClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#broker_node_group_info MskCluster#broker_node_group_info}.
	BrokerNodeGroupInfo *MskClusterBrokerNodeGroupInfo `field:"required" json:"brokerNodeGroupInfo" yaml:"brokerNodeGroupInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#cluster_name MskCluster#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#kafka_version MskCluster#kafka_version}.
	KafkaVersion *string `field:"required" json:"kafkaVersion" yaml:"kafkaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#number_of_broker_nodes MskCluster#number_of_broker_nodes}.
	NumberOfBrokerNodes *float64 `field:"required" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#client_authentication MskCluster#client_authentication}.
	ClientAuthentication *MskClusterClientAuthentication `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#configuration_info MskCluster#configuration_info}.
	ConfigurationInfo *MskClusterConfigurationInfo `field:"optional" json:"configurationInfo" yaml:"configurationInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#encryption_info MskCluster#encryption_info}.
	EncryptionInfo *MskClusterEncryptionInfo `field:"optional" json:"encryptionInfo" yaml:"encryptionInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#enhanced_monitoring MskCluster#enhanced_monitoring}.
	EnhancedMonitoring *string `field:"optional" json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#logging_info MskCluster#logging_info}.
	LoggingInfo *MskClusterLoggingInfo `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#open_monitoring MskCluster#open_monitoring}.
	OpenMonitoring *MskClusterOpenMonitoring `field:"optional" json:"openMonitoring" yaml:"openMonitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#rebalancing MskCluster#rebalancing}.
	Rebalancing *MskClusterRebalancing `field:"optional" json:"rebalancing" yaml:"rebalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#storage_mode MskCluster#storage_mode}.
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#tags MskCluster#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/msk_cluster#zookeeper_access MskCluster#zookeeper_access}.
	ZookeeperAccess *MskClusterZookeeperAccess `field:"optional" json:"zookeeperAccess" yaml:"zookeeperAccess"`
}

