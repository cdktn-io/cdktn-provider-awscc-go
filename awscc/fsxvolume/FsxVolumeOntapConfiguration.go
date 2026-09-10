// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#aggregate_configuration FsxVolume#aggregate_configuration}.
	AggregateConfiguration *FsxVolumeOntapConfigurationAggregateConfiguration `field:"optional" json:"aggregateConfiguration" yaml:"aggregateConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#copy_tags_to_backups FsxVolume#copy_tags_to_backups}.
	CopyTagsToBackups *string `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#junction_path FsxVolume#junction_path}.
	JunctionPath *string `field:"optional" json:"junctionPath" yaml:"junctionPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#ontap_volume_type FsxVolume#ontap_volume_type}.
	OntapVolumeType *string `field:"optional" json:"ontapVolumeType" yaml:"ontapVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#security_style FsxVolume#security_style}.
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#size_in_bytes FsxVolume#size_in_bytes}.
	SizeInBytes *string `field:"optional" json:"sizeInBytes" yaml:"sizeInBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#size_in_megabytes FsxVolume#size_in_megabytes}.
	SizeInMegabytes *string `field:"optional" json:"sizeInMegabytes" yaml:"sizeInMegabytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#snaplock_configuration FsxVolume#snaplock_configuration}.
	SnaplockConfiguration *FsxVolumeOntapConfigurationSnaplockConfiguration `field:"optional" json:"snaplockConfiguration" yaml:"snaplockConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#snapshot_policy FsxVolume#snapshot_policy}.
	SnapshotPolicy *string `field:"optional" json:"snapshotPolicy" yaml:"snapshotPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#storage_efficiency_enabled FsxVolume#storage_efficiency_enabled}.
	StorageEfficiencyEnabled *string `field:"optional" json:"storageEfficiencyEnabled" yaml:"storageEfficiencyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#storage_virtual_machine_id FsxVolume#storage_virtual_machine_id}.
	StorageVirtualMachineId *string `field:"optional" json:"storageVirtualMachineId" yaml:"storageVirtualMachineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#tiering_policy FsxVolume#tiering_policy}.
	TieringPolicy *FsxVolumeOntapConfigurationTieringPolicy `field:"optional" json:"tieringPolicy" yaml:"tieringPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#volume_style FsxVolume#volume_style}.
	VolumeStyle *string `field:"optional" json:"volumeStyle" yaml:"volumeStyle"`
}

