// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetSageMakerPipelineParametersPipelineParameterListStruct struct {
	// Name of parameter to start execution of a SageMaker Model Building Pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#name SchedulerSchedule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value of parameter to start execution of a SageMaker Model Building Pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/scheduler_schedule#value SchedulerSchedule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

