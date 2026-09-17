// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetSageMakerPipelineParameters struct {
	// List of Parameter names and values for SageMaker Model Building Pipeline execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/scheduler_schedule#pipeline_parameter_list SchedulerSchedule#pipeline_parameter_list}
	PipelineParameterList interface{} `field:"optional" json:"pipelineParameterList" yaml:"pipelineParameterList"`
}

