// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfiguration struct {
	// Used to specify the configuration options for an FSx for ONTAP volume's storage aggregate or aggregates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#aggregate_configuration FsxVolume#aggregate_configuration}
	AggregateConfiguration *FsxVolumeOntapConfigurationAggregateConfiguration `field:"optional" json:"aggregateConfiguration" yaml:"aggregateConfiguration"`
	// A boolean flag indicating whether tags for the volume should be copied to backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#copy_tags_to_backups FsxVolume#copy_tags_to_backups}
	CopyTagsToBackups *string `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Specifies the location in the SVM's namespace where the volume is mounted.
	//
	// This parameter is required. The JunctionPath must have a leading forward slash, such as /vol3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#junction_path FsxVolume#junction_path}
	JunctionPath *string `field:"optional" json:"junctionPath" yaml:"junctionPath"`
	// Specifies the type of volume you are creating. Valid values are the following: RW or DP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#ontap_volume_type FsxVolume#ontap_volume_type}
	OntapVolumeType *string `field:"optional" json:"ontapVolumeType" yaml:"ontapVolumeType"`
	// Specifies the security style for the volume.
	//
	// If a volume's security style is not specified, it is automatically set to the root volume's security style.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#security_style FsxVolume#security_style}
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// Specifies the configured size of the volume, in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#size_in_bytes FsxVolume#size_in_bytes}
	SizeInBytes *string `field:"optional" json:"sizeInBytes" yaml:"sizeInBytes"`
	// Use SizeInBytes instead. Specifies the size of the volume, in megabytes (MB), that you are creating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#size_in_megabytes FsxVolume#size_in_megabytes}
	SizeInMegabytes *string `field:"optional" json:"sizeInMegabytes" yaml:"sizeInMegabytes"`
	// The SnapLock configuration object for an FSx for ONTAP SnapLock volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#snaplock_configuration FsxVolume#snaplock_configuration}
	SnaplockConfiguration *FsxVolumeOntapConfigurationSnaplockConfiguration `field:"optional" json:"snaplockConfiguration" yaml:"snaplockConfiguration"`
	// Specifies the snapshot policy for the volume. There are three built-in snapshot policies: default, default-1weekly, none.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#snapshot_policy FsxVolume#snapshot_policy}
	SnapshotPolicy *string `field:"optional" json:"snapshotPolicy" yaml:"snapshotPolicy"`
	// Set to true to enable deduplication, compression, and compaction storage efficiency features on the volume, or set to false to disable them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#storage_efficiency_enabled FsxVolume#storage_efficiency_enabled}
	StorageEfficiencyEnabled *string `field:"optional" json:"storageEfficiencyEnabled" yaml:"storageEfficiencyEnabled"`
	// Specifies the ONTAP SVM in which to create the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#storage_virtual_machine_id FsxVolume#storage_virtual_machine_id}
	StorageVirtualMachineId *string `field:"optional" json:"storageVirtualMachineId" yaml:"storageVirtualMachineId"`
	// Describes the data tiering policy for an ONTAP volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#tiering_policy FsxVolume#tiering_policy}
	TieringPolicy *FsxVolumeOntapConfigurationTieringPolicy `field:"optional" json:"tieringPolicy" yaml:"tieringPolicy"`
	// Use to specify the style of an ONTAP volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#volume_style FsxVolume#volume_style}
	VolumeStyle *string `field:"optional" json:"volumeStyle" yaml:"volumeStyle"`
}

