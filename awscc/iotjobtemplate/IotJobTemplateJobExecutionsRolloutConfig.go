// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjobtemplate


type IotJobTemplateJobExecutionsRolloutConfig struct {
	// The rate of increase for a job rollout.
	//
	// This parameter allows you to define an exponential rate for a job rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_job_template#exponential_rollout_rate IotJobTemplate#exponential_rollout_rate}
	ExponentialRolloutRate *IotJobTemplateJobExecutionsRolloutConfigExponentialRolloutRate `field:"optional" json:"exponentialRolloutRate" yaml:"exponentialRolloutRate"`
	// The maximum number of things that will be notified of a pending job, per minute.
	//
	// This parameter allows you to create a staged rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_job_template#maximum_per_minute IotJobTemplate#maximum_per_minute}
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

