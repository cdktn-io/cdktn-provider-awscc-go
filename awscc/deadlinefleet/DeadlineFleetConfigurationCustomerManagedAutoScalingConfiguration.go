// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfigurationCustomerManagedAutoScalingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/deadline_fleet#scale_out_workers_per_minute DeadlineFleet#scale_out_workers_per_minute}.
	ScaleOutWorkersPerMinute *float64 `field:"optional" json:"scaleOutWorkersPerMinute" yaml:"scaleOutWorkersPerMinute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/deadline_fleet#standby_worker_count DeadlineFleet#standby_worker_count}.
	StandbyWorkerCount *float64 `field:"optional" json:"standbyWorkerCount" yaml:"standbyWorkerCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/deadline_fleet#worker_idle_duration_seconds DeadlineFleet#worker_idle_duration_seconds}.
	WorkerIdleDurationSeconds *float64 `field:"optional" json:"workerIdleDurationSeconds" yaml:"workerIdleDurationSeconds"`
}

