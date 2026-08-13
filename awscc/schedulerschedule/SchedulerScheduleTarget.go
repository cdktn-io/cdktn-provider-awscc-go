// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTarget struct {
	// The Amazon Resource Name (ARN) of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#arn SchedulerSchedule#arn}
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The Amazon Resource Name (ARN) of the IAM role to be used for this target when the schedule is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#role_arn SchedulerSchedule#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A DeadLetterConfig object that contains information about a dead-letter queue configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#dead_letter_config SchedulerSchedule#dead_letter_config}
	DeadLetterConfig *SchedulerScheduleTargetDeadLetterConfig `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// The custom parameters to be used when the target is an Amazon ECS task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#ecs_parameters SchedulerSchedule#ecs_parameters}
	EcsParameters *SchedulerScheduleTargetEcsParameters `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// EventBridge PutEvent predefined target type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#event_bridge_parameters SchedulerSchedule#event_bridge_parameters}
	EventBridgeParameters *SchedulerScheduleTargetEventBridgeParameters `field:"optional" json:"eventBridgeParameters" yaml:"eventBridgeParameters"`
	// The text, or well-formed JSON, passed to the target.
	//
	// If you are configuring a templated Lambda, AWS Step Functions, or Amazon EventBridge target, the input must be a well-formed JSON. For all other target types, a JSON is not required. If you do not specify anything for this field, EventBridge Scheduler delivers a default notification to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#input SchedulerSchedule#input}
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The custom parameter you can use to control the shard to which EventBridge Scheduler sends the event.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#kinesis_parameters SchedulerSchedule#kinesis_parameters}
	KinesisParameters *SchedulerScheduleTargetKinesisParameters `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// A RetryPolicy object that includes information about the retry policy settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#retry_policy SchedulerSchedule#retry_policy}
	RetryPolicy *SchedulerScheduleTargetRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// These are custom parameters to use when the target is a SageMaker Model Building Pipeline that starts based on AWS EventBridge Scheduler schedules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#sage_maker_pipeline_parameters SchedulerSchedule#sage_maker_pipeline_parameters}
	SageMakerPipelineParameters *SchedulerScheduleTargetSageMakerPipelineParameters `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// Contains the message group ID to use when the target is a FIFO queue.
	//
	// If you specify an SQS FIFO queue as a target, the queue must have content-based deduplication enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#sqs_parameters SchedulerSchedule#sqs_parameters}
	SqsParameters *SchedulerScheduleTargetSqsParameters `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}

