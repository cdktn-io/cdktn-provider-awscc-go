// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfigurationCustomerManagedWorkerCapabilities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#accelerator_count DeadlineFleet#accelerator_count}.
	AcceleratorCount *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorCount `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#accelerator_total_memory_mi_b DeadlineFleet#accelerator_total_memory_mi_b}.
	AcceleratorTotalMemoryMiB *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesAcceleratorTotalMemoryMiB `field:"optional" json:"acceleratorTotalMemoryMiB" yaml:"acceleratorTotalMemoryMiB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#accelerator_types DeadlineFleet#accelerator_types}.
	AcceleratorTypes *[]*string `field:"optional" json:"acceleratorTypes" yaml:"acceleratorTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#cpu_architecture_type DeadlineFleet#cpu_architecture_type}.
	CpuArchitectureType *string `field:"optional" json:"cpuArchitectureType" yaml:"cpuArchitectureType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#custom_amounts DeadlineFleet#custom_amounts}.
	CustomAmounts interface{} `field:"optional" json:"customAmounts" yaml:"customAmounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#custom_attributes DeadlineFleet#custom_attributes}.
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#memory_mi_b DeadlineFleet#memory_mi_b}.
	MemoryMiB *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesMemoryMiB `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#os_family DeadlineFleet#os_family}.
	OsFamily *string `field:"optional" json:"osFamily" yaml:"osFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_fleet#v_cpu_count DeadlineFleet#v_cpu_count}.
	VCpuCount *DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesVCpuCount `field:"optional" json:"vCpuCount" yaml:"vCpuCount"`
}

