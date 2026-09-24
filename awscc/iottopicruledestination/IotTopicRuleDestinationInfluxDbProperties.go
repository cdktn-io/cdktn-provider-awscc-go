// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicruledestination


type IotTopicRuleDestinationInfluxDbProperties struct {
	// The endpoint URL of the InfluxDB database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule_destination#endpoint IotTopicRuleDestination#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The version of the InfluxDB database (for example, V2 or V3).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule_destination#influx_db_version IotTopicRuleDestination#influx_db_version}
	InfluxDbVersion *string `field:"optional" json:"influxDbVersion" yaml:"influxDbVersion"`
	// The ARN or name of the Secrets Manager secret containing the InfluxDB API token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule_destination#secret_id IotTopicRuleDestination#secret_id}
	SecretId *string `field:"optional" json:"secretId" yaml:"secretId"`
	// The key name within the secret that contains the InfluxDB token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule_destination#secret_key IotTopicRuleDestination#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// The type of the secret value (SecretString or SecretBinary).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_topic_rule_destination#secret_type IotTopicRuleDestination#secret_type}
	SecretType *string `field:"optional" json:"secretType" yaml:"secretType"`
}

