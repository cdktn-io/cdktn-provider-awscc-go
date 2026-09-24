// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KafkaconnectConnectorConfig struct {
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
	// Information about the capacity allocated to the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#capacity KafkaconnectConnector#capacity}
	Capacity *KafkaconnectConnectorCapacity `field:"required" json:"capacity" yaml:"capacity"`
	// The configuration for the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#connector_configuration KafkaconnectConnector#connector_configuration}
	ConnectorConfiguration *map[string]*string `field:"required" json:"connectorConfiguration" yaml:"connectorConfiguration"`
	// The name of the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#connector_name KafkaconnectConnector#connector_name}
	ConnectorName *string `field:"required" json:"connectorName" yaml:"connectorName"`
	// Details of how to connect to the Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#kafka_cluster KafkaconnectConnector#kafka_cluster}
	KafkaCluster *KafkaconnectConnectorKafkaCluster `field:"required" json:"kafkaCluster" yaml:"kafkaCluster"`
	// Details of the client authentication used by the Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#kafka_cluster_client_authentication KafkaconnectConnector#kafka_cluster_client_authentication}
	KafkaClusterClientAuthentication *KafkaconnectConnectorKafkaClusterClientAuthentication `field:"required" json:"kafkaClusterClientAuthentication" yaml:"kafkaClusterClientAuthentication"`
	// Details of encryption in transit to the Kafka cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#kafka_cluster_encryption_in_transit KafkaconnectConnector#kafka_cluster_encryption_in_transit}
	KafkaClusterEncryptionInTransit *KafkaconnectConnectorKafkaClusterEncryptionInTransit `field:"required" json:"kafkaClusterEncryptionInTransit" yaml:"kafkaClusterEncryptionInTransit"`
	// The version of Kafka Connect. It has to be compatible with both the Kafka cluster's version and the plugins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#kafka_connect_version KafkaconnectConnector#kafka_connect_version}
	KafkaConnectVersion *string `field:"required" json:"kafkaConnectVersion" yaml:"kafkaConnectVersion"`
	// List of plugins to use with the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#plugins KafkaconnectConnector#plugins}
	Plugins interface{} `field:"required" json:"plugins" yaml:"plugins"`
	// The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon S3 objects and other external resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#service_execution_role_arn KafkaconnectConnector#service_execution_role_arn}
	ServiceExecutionRoleArn *string `field:"required" json:"serviceExecutionRoleArn" yaml:"serviceExecutionRoleArn"`
	// A summary description of the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#connector_description KafkaconnectConnector#connector_description}
	ConnectorDescription *string `field:"optional" json:"connectorDescription" yaml:"connectorDescription"`
	// Details of what logs are delivered and where they are delivered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#log_delivery KafkaconnectConnector#log_delivery}
	LogDelivery *KafkaconnectConnectorLogDelivery `field:"optional" json:"logDelivery" yaml:"logDelivery"`
	// The network type of the Connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#network_type KafkaconnectConnector#network_type}
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#tags KafkaconnectConnector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies the worker configuration to use with the connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#worker_configuration KafkaconnectConnector#worker_configuration}
	WorkerConfiguration *KafkaconnectConnectorWorkerConfiguration `field:"optional" json:"workerConfiguration" yaml:"workerConfiguration"`
}

