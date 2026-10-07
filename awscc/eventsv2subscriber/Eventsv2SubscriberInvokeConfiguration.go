// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberInvokeConfiguration struct {
	// The ARN of the IAM role the service assumes to invoke the target.
	//
	// The role must belong to the same account as the subscriber.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#role_arn Eventsv2Subscriber#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The Amazon Resource Name (ARN) of the target that the subscriber invokes.
	//
	// For universal service integration targets, use the form arn:{partition}:events:::aws-sdk:{service}:{apiAction}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#target_arn Eventsv2Subscriber#target_arn}
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Parameters for forwarding events to another EventBridge event bus, used when TargetArn is an event bus ARN of the form arn:{partition}:events:{region}:{account}:event-busv2/{name}/{id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#event_bus_v2_parameters Eventsv2Subscriber#event_bus_v2_parameters}
	EventBusV2Parameters *Eventsv2SubscriberInvokeConfigurationEventBusV2Parameters `field:"optional" json:"eventBusV2Parameters" yaml:"eventBusV2Parameters"`
	// Parameters for invoking an HTTP endpoint target, such as an Amazon API Gateway endpoint or an EventBridge API destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#http_parameters Eventsv2Subscriber#http_parameters}
	HttpParameters *Eventsv2SubscriberInvokeConfigurationHttpParameters `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Parameters for writing events to an Amazon Kinesis Data Streams target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#kinesis_parameters Eventsv2Subscriber#kinesis_parameters}
	KinesisParameters *Eventsv2SubscriberInvokeConfigurationKinesisParameters `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// Parameters for invoking an AWS Lambda function target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#lambda_parameters Eventsv2Subscriber#lambda_parameters}
	LambdaParameters *Eventsv2SubscriberInvokeConfigurationLambdaParameters `field:"optional" json:"lambdaParameters" yaml:"lambdaParameters"`
	// Parameters for publishing events to an Amazon SNS topic target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#sns_parameters Eventsv2Subscriber#sns_parameters}
	SnsParameters *Eventsv2SubscriberInvokeConfigurationSnsParameters `field:"optional" json:"snsParameters" yaml:"snsParameters"`
	// Parameters for sending events to an Amazon SQS queue target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#sqs_parameters Eventsv2Subscriber#sqs_parameters}
	SqsParameters *Eventsv2SubscriberInvokeConfigurationSqsParameters `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
	// Parameters for starting an AWS Step Functions state machine execution target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#step_functions_parameters Eventsv2Subscriber#step_functions_parameters}
	StepFunctionsParameters *Eventsv2SubscriberInvokeConfigurationStepFunctionsParameters `field:"optional" json:"stepFunctionsParameters" yaml:"stepFunctionsParameters"`
	// Parameters for invoking an AWS service API as a universal service integration target, used when TargetArn has the form arn:{partition}:events:::aws-sdk:{service}:{apiAction}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#universal_target_parameters Eventsv2Subscriber#universal_target_parameters}
	UniversalTargetParameters *Eventsv2SubscriberInvokeConfigurationUniversalTargetParameters `field:"optional" json:"universalTargetParameters" yaml:"universalTargetParameters"`
}

