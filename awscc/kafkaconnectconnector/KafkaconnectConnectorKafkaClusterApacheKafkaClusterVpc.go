// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kafkaconnectconnector


type KafkaconnectConnectorKafkaClusterApacheKafkaClusterVpc struct {
	// The AWS security groups to associate with the elastic network interfaces in order to specify what the connector has access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#security_groups KafkaconnectConnector#security_groups}
	SecurityGroups *[]*string `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The list of subnets to connect to in the virtual private cloud (VPC).
	//
	// AWS creates elastic network interfaces inside these subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kafkaconnect_connector#subnets KafkaconnectConnector#subnets}
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
}

