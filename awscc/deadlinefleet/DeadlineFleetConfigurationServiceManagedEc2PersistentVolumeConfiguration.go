// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#iops DeadlineFleet#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#last_used_ttl_hours DeadlineFleet#last_used_ttl_hours}.
	LastUsedTtlHours *float64 `field:"optional" json:"lastUsedTtlHours" yaml:"lastUsedTtlHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#mount_path DeadlineFleet#mount_path}.
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#size_gi_b DeadlineFleet#size_gi_b}.
	SizeGiB *float64 `field:"optional" json:"sizeGiB" yaml:"sizeGiB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_fleet#throughput_mi_b DeadlineFleet#throughput_mi_b}.
	ThroughputMiB *float64 `field:"optional" json:"throughputMiB" yaml:"throughputMiB"`
}

