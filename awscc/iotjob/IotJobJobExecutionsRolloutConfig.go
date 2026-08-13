// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobJobExecutionsRolloutConfig struct {
	// Allows you to create an exponential rate of rollout for a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job#exponential_rate IotJob#exponential_rate}
	ExponentialRate *IotJobJobExecutionsRolloutConfigExponentialRate `field:"optional" json:"exponentialRate" yaml:"exponentialRate"`
	// The maximum number of things that will be notified of a pending job, per minute.
	//
	// This parameter allows you to create a staged rollout.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_job#maximum_per_minute IotJob#maximum_per_minute}
	MaximumPerMinute *float64 `field:"optional" json:"maximumPerMinute" yaml:"maximumPerMinute"`
}

