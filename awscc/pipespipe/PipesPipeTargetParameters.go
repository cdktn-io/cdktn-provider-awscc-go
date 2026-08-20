// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#batch_job_parameters PipesPipe#batch_job_parameters}.
	BatchJobParameters *PipesPipeTargetParametersBatchJobParameters `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#cloudwatch_logs_parameters PipesPipe#cloudwatch_logs_parameters}.
	CloudwatchLogsParameters *PipesPipeTargetParametersCloudwatchLogsParameters `field:"optional" json:"cloudwatchLogsParameters" yaml:"cloudwatchLogsParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#ecs_task_parameters PipesPipe#ecs_task_parameters}.
	EcsTaskParameters *PipesPipeTargetParametersEcsTaskParameters `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#event_bridge_event_bus_parameters PipesPipe#event_bridge_event_bus_parameters}.
	EventBridgeEventBusParameters *PipesPipeTargetParametersEventBridgeEventBusParameters `field:"optional" json:"eventBridgeEventBusParameters" yaml:"eventBridgeEventBusParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#http_parameters PipesPipe#http_parameters}.
	HttpParameters *PipesPipeTargetParametersHttpParameters `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#input_template PipesPipe#input_template}.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#kinesis_stream_parameters PipesPipe#kinesis_stream_parameters}.
	KinesisStreamParameters *PipesPipeTargetParametersKinesisStreamParameters `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#lambda_function_parameters PipesPipe#lambda_function_parameters}.
	LambdaFunctionParameters *PipesPipeTargetParametersLambdaFunctionParameters `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#redshift_data_parameters PipesPipe#redshift_data_parameters}.
	RedshiftDataParameters *PipesPipeTargetParametersRedshiftDataParameters `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#sage_maker_pipeline_parameters PipesPipe#sage_maker_pipeline_parameters}.
	SageMakerPipelineParameters *PipesPipeTargetParametersSageMakerPipelineParameters `field:"optional" json:"sageMakerPipelineParameters" yaml:"sageMakerPipelineParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#sqs_queue_parameters PipesPipe#sqs_queue_parameters}.
	SqsQueueParameters *PipesPipeTargetParametersSqsQueueParameters `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#step_function_state_machine_parameters PipesPipe#step_function_state_machine_parameters}.
	StepFunctionStateMachineParameters *PipesPipeTargetParametersStepFunctionStateMachineParameters `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pipes_pipe#timestream_parameters PipesPipe#timestream_parameters}.
	TimestreamParameters *PipesPipeTargetParametersTimestreamParameters `field:"optional" json:"timestreamParameters" yaml:"timestreamParameters"`
}

