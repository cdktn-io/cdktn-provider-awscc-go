// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsrule


type EventsRuleTargets struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#app_sync_parameters EventsRule#app_sync_parameters}.
	AppSyncParameters *EventsRuleTargetsAppSyncParameters `field:"optional" json:"appSyncParameters" yaml:"appSyncParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#arn EventsRule#arn}.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#batch_parameters EventsRule#batch_parameters}.
	BatchParameters *EventsRuleTargetsBatchParameters `field:"optional" json:"batchParameters" yaml:"batchParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#dead_letter_config EventsRule#dead_letter_config}.
	DeadLetterConfig *EventsRuleTargetsDeadLetterConfig `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#ecs_parameters EventsRule#ecs_parameters}.
	EcsParameters *EventsRuleTargetsEcsParameters `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#http_parameters EventsRule#http_parameters}.
	HttpParameters *EventsRuleTargetsHttpParameters `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#id EventsRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#input EventsRule#input}.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#input_path EventsRule#input_path}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#input_transformer EventsRule#input_transformer}.
	InputTransformer *EventsRuleTargetsInputTransformer `field:"optional" json:"inputTransformer" yaml:"inputTransformer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#kinesis_parameters EventsRule#kinesis_parameters}.
	KinesisParameters *EventsRuleTargetsKinesisParameters `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#redshift_data_parameters EventsRule#redshift_data_parameters}.
	RedshiftDataParameters *EventsRuleTargetsRedshiftDataParameters `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#retry_policy EventsRule#retry_policy}.
	RetryPolicy *EventsRuleTargetsRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#role_arn EventsRule#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#run_command_parameters EventsRule#run_command_parameters}.
	RunCommandParameters *EventsRuleTargetsRunCommandParameters `field:"optional" json:"runCommandParameters" yaml:"runCommandParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#sage_maker_pipeline_parameters EventsRule#sage_maker_pipeline_parameters}.
	SageMakerPipelineParameters *EventsRuleTargetsSageMakerPipelineParameters `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/events_rule#sqs_parameters EventsRule#sqs_parameters}.
	SqsParameters *EventsRuleTargetsSqsParameters `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

