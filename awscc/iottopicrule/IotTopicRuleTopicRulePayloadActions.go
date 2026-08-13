// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#cloudwatch_alarm IotTopicRule#cloudwatch_alarm}.
	CloudwatchAlarm *IotTopicRuleTopicRulePayloadActionsCloudwatchAlarm `field:"optional" json:"cloudwatchAlarm" yaml:"cloudwatchAlarm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#cloudwatch_logs IotTopicRule#cloudwatch_logs}.
	CloudwatchLogs *IotTopicRuleTopicRulePayloadActionsCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#cloudwatch_metric IotTopicRule#cloudwatch_metric}.
	CloudwatchMetric *IotTopicRuleTopicRulePayloadActionsCloudwatchMetric `field:"optional" json:"cloudwatchMetric" yaml:"cloudwatchMetric"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#dynamo_db IotTopicRule#dynamo_db}.
	DynamoDb *IotTopicRuleTopicRulePayloadActionsDynamoDb `field:"optional" json:"dynamoDb" yaml:"dynamoDb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#dynamo_d_bv_2 IotTopicRule#dynamo_d_bv_2}.
	DynamoDBv2 *IotTopicRuleTopicRulePayloadActionsDynamoDBv2 `field:"optional" json:"dynamoDBv2" yaml:"dynamoDBv2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#elasticsearch IotTopicRule#elasticsearch}.
	Elasticsearch *IotTopicRuleTopicRulePayloadActionsElasticsearch `field:"optional" json:"elasticsearch" yaml:"elasticsearch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#firehose IotTopicRule#firehose}.
	Firehose *IotTopicRuleTopicRulePayloadActionsFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#http IotTopicRule#http}.
	Http *IotTopicRuleTopicRulePayloadActionsHttp `field:"optional" json:"http" yaml:"http"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#iot_analytics IotTopicRule#iot_analytics}.
	IotAnalytics *IotTopicRuleTopicRulePayloadActionsIotAnalytics `field:"optional" json:"iotAnalytics" yaml:"iotAnalytics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#iot_events IotTopicRule#iot_events}.
	IotEvents *IotTopicRuleTopicRulePayloadActionsIotEvents `field:"optional" json:"iotEvents" yaml:"iotEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#iot_site_wise IotTopicRule#iot_site_wise}.
	IotSiteWise *IotTopicRuleTopicRulePayloadActionsIotSiteWise `field:"optional" json:"iotSiteWise" yaml:"iotSiteWise"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#kafka IotTopicRule#kafka}.
	Kafka *IotTopicRuleTopicRulePayloadActionsKafka `field:"optional" json:"kafka" yaml:"kafka"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#kinesis IotTopicRule#kinesis}.
	Kinesis *IotTopicRuleTopicRulePayloadActionsKinesis `field:"optional" json:"kinesis" yaml:"kinesis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#lambda IotTopicRule#lambda}.
	Lambda *IotTopicRuleTopicRulePayloadActionsLambda `field:"optional" json:"lambda" yaml:"lambda"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#location IotTopicRule#location}.
	Location *IotTopicRuleTopicRulePayloadActionsLocation `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#open_search IotTopicRule#open_search}.
	OpenSearch *IotTopicRuleTopicRulePayloadActionsOpenSearch `field:"optional" json:"openSearch" yaml:"openSearch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#republish IotTopicRule#republish}.
	Republish *IotTopicRuleTopicRulePayloadActionsRepublish `field:"optional" json:"republish" yaml:"republish"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#s3 IotTopicRule#s3}.
	S3 *IotTopicRuleTopicRulePayloadActionsS3 `field:"optional" json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#sns IotTopicRule#sns}.
	Sns *IotTopicRuleTopicRulePayloadActionsSns `field:"optional" json:"sns" yaml:"sns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#sqs IotTopicRule#sqs}.
	Sqs *IotTopicRuleTopicRulePayloadActionsSqs `field:"optional" json:"sqs" yaml:"sqs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#step_functions IotTopicRule#step_functions}.
	StepFunctions *IotTopicRuleTopicRulePayloadActionsStepFunctions `field:"optional" json:"stepFunctions" yaml:"stepFunctions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_topic_rule#timestream IotTopicRule#timestream}.
	Timestream *IotTopicRuleTopicRulePayloadActionsTimestream `field:"optional" json:"timestream" yaml:"timestream"`
}

