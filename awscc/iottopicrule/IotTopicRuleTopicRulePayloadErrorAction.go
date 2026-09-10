// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadErrorAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#cloudwatch_alarm IotTopicRule#cloudwatch_alarm}.
	CloudwatchAlarm *IotTopicRuleTopicRulePayloadErrorActionCloudwatchAlarm `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#cloudwatch_logs IotTopicRule#cloudwatch_logs}.
	CloudwatchLogs *IotTopicRuleTopicRulePayloadErrorActionCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#cloudwatch_metric IotTopicRule#cloudwatch_metric}.
	CloudwatchMetric *IotTopicRuleTopicRulePayloadErrorActionCloudwatchMetric `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#dynamo_db IotTopicRule#dynamo_db}.
	DynamoDb *IotTopicRuleTopicRulePayloadErrorActionDynamoDb `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#dynamo_d_bv_2 IotTopicRule#dynamo_d_bv_2}.
	DynamoDBv2 *IotTopicRuleTopicRulePayloadErrorActionDynamoDBv2 `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#elasticsearch IotTopicRule#elasticsearch}.
	Elasticsearch *IotTopicRuleTopicRulePayloadErrorActionElasticsearch `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#firehose IotTopicRule#firehose}.
	Firehose *IotTopicRuleTopicRulePayloadErrorActionFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#http IotTopicRule#http}.
	Http *IotTopicRuleTopicRulePayloadErrorActionHttp `field:"optional" json:"http" yaml:"http"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#influx_db IotTopicRule#influx_db}.
	InfluxDb *IotTopicRuleTopicRulePayloadErrorActionInfluxDb `field:"optional" json:"influxDb" yaml:"influxDb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#iot_analytics IotTopicRule#iot_analytics}.
	IotAnalytics *IotTopicRuleTopicRulePayloadErrorActionIotAnalytics `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#iot_events IotTopicRule#iot_events}.
	IotEvents *IotTopicRuleTopicRulePayloadErrorActionIotEvents `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#iot_site_wise IotTopicRule#iot_site_wise}.
	IotSiteWise *IotTopicRuleTopicRulePayloadErrorActionIotSiteWise `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#kafka IotTopicRule#kafka}.
	Kafka *IotTopicRuleTopicRulePayloadErrorActionKafka `field:"optional" json:"kafka" yaml:"kafka"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#kinesis IotTopicRule#kinesis}.
	Kinesis *IotTopicRuleTopicRulePayloadErrorActionKinesis `field:"optional" json:"kinesis" yaml:"kinesis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#lambda IotTopicRule#lambda}.
	Lambda *IotTopicRuleTopicRulePayloadErrorActionLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#location IotTopicRule#location}.
	Location *IotTopicRuleTopicRulePayloadErrorActionLocation `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#open_search IotTopicRule#open_search}.
	OpenSearch *IotTopicRuleTopicRulePayloadErrorActionOpenSearch `field:"optional" json:"openSearch" yaml:"openSearch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#republish IotTopicRule#republish}.
	Republish *IotTopicRuleTopicRulePayloadErrorActionRepublish `field:"optional" json:"republish" yaml:"republish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#s3 IotTopicRule#s3}.
	S3 *IotTopicRuleTopicRulePayloadErrorActionS3 `field:"optional" json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#sns IotTopicRule#sns}.
	Sns *IotTopicRuleTopicRulePayloadErrorActionSns `field:"optional" json:"sns" yaml:"sns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#sqs IotTopicRule#sqs}.
	Sqs *IotTopicRuleTopicRulePayloadErrorActionSqs `field:"optional" json:"sqs" yaml:"sqs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#step_functions IotTopicRule#step_functions}.
	StepFunctions *IotTopicRuleTopicRulePayloadErrorActionStepFunctions `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule#timestream IotTopicRule#timestream}.
	Timestream *IotTopicRuleTopicRulePayloadErrorActionTimestream `field:"optional" json:"timestream" yaml:"timestream"`
}

