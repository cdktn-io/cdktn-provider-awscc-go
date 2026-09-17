// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#accelerator_capabilities DeadlineFleet#accelerator_capabilities}.
	AcceleratorCapabilities *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesAcceleratorCapabilities `field:"optional" json:"acceleratorCapabilities" yaml:"acceleratorCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#allowed_instance_types DeadlineFleet#allowed_instance_types}.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#cpu_architecture_type DeadlineFleet#cpu_architecture_type}.
	CpuArchitectureType *string `field:"optional" json:"cpuArchitectureType" yaml:"cpuArchitectureType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#custom_amounts DeadlineFleet#custom_amounts}.
	CustomAmounts interface{} `field:"optional" json:"customAmounts" yaml:"customAmounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#custom_attributes DeadlineFleet#custom_attributes}.
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#excluded_instance_types DeadlineFleet#excluded_instance_types}.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#memory_mi_b DeadlineFleet#memory_mi_b}.
	MemoryMiB *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesMemoryMiB `field:"optional" json:"memoryMiB" yaml:"memoryMiB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#os_family DeadlineFleet#os_family}.
	OsFamily *string `field:"optional" json:"osFamily" yaml:"osFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#root_ebs_volume DeadlineFleet#root_ebs_volume}.
	RootEbsVolume *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesRootEbsVolume `field:"optional" json:"rootEbsVolume" yaml:"rootEbsVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#v_cpu_count DeadlineFleet#v_cpu_count}.
	VCpuCount *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesVCpuCount `field:"optional" json:"vCpuCount" yaml:"vCpuCount"`
}

