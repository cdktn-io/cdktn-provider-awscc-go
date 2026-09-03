// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob


type IotJobTimeoutConfig struct {
	// Specifies the amount of time, in minutes, this device has to finish execution of this job.
	//
	// The timeout interval can be anywhere between 1 minute and 7 days (1 to 10080 minutes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_job#in_progress_timeout_in_minutes IotJob#in_progress_timeout_in_minutes}
	InProgressTimeoutInMinutes *float64 `field:"optional" json:"inProgressTimeoutInMinutes" yaml:"inProgressTimeoutInMinutes"`
}

