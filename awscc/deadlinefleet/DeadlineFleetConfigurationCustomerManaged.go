// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfigurationCustomerManaged struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#auto_scaling_configuration DeadlineFleet#auto_scaling_configuration}.
	AutoScalingConfiguration *DeadlineFleetConfigurationCustomerManagedAutoScalingConfiguration `field:"optional" json:"autoScalingConfiguration" yaml:"autoScalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#mode DeadlineFleet#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#storage_profile_id DeadlineFleet#storage_profile_id}.
	StorageProfileId *string `field:"optional" json:"storageProfileId" yaml:"storageProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#tag_propagation_mode DeadlineFleet#tag_propagation_mode}.
	TagPropagationMode *string `field:"optional" json:"tagPropagationMode" yaml:"tagPropagationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#worker_capabilities DeadlineFleet#worker_capabilities}.
	WorkerCapabilities *DeadlineFleetConfigurationCustomerManagedWorkerCapabilities `field:"optional" json:"workerCapabilities" yaml:"workerCapabilities"`
}

