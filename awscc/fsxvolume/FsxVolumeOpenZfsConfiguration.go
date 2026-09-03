// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOpenZfsConfiguration struct {
	// A Boolean value indicating whether tags for the volume should be copied to snapshots.
	//
	// This value defaults to false. If this value is set to true, and you do not specify any tags, all tags for the original volume are copied over to snapshots. If this value is set to true, and you do specify one or more tags, only the specified tags for the original volume are copied over to snapshots. If you specify one or more tags when creating a new snapshot, no tags are copied over from the original volume, regardless of this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#copy_tags_to_snapshots FsxVolume#copy_tags_to_snapshots}
	CopyTagsToSnapshots interface{} `field:"optional" json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Specifies the method used to compress the data on the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#data_compression_type FsxVolume#data_compression_type}
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// The configuration object for mounting a Network File System (NFS) file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#nfs_exports FsxVolume#nfs_exports}
	NfsExports interface{} `field:"optional" json:"nfsExports" yaml:"nfsExports"`
	// The configuration object for mounting a Network File System (NFS) file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#options FsxVolume#options}
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// The configuration of an Amazon FSx for OpenZFS volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#origin_snapshot FsxVolume#origin_snapshot}
	OriginSnapshot *FsxVolumeOpenZfsConfigurationOriginSnapshot `field:"optional" json:"originSnapshot" yaml:"originSnapshot"`
	// The ID of the volume to use as the parent volume of the volume that you are creating.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#parent_volume_id FsxVolume#parent_volume_id}
	ParentVolumeId *string `field:"optional" json:"parentVolumeId" yaml:"parentVolumeId"`
	// A Boolean value indicating whether the volume is read-only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#read_only FsxVolume#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Specifies the suggested block size for a volume in a ZFS dataset, in kibibytes (KiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#record_size_ki_b FsxVolume#record_size_ki_b}
	RecordSizeKiB *float64 `field:"optional" json:"recordSizeKiB" yaml:"recordSizeKiB"`
	// Sets the maximum storage size in gibibytes (GiB) for the volume.
	//
	// You can specify a quota that is larger than the storage on the parent volume. A volume quota limits the amount of storage that the volume can consume to the configured amount, but does not guarantee the space will be available on the parent volume. To guarantee quota space, you must also set StorageCapacityReservationGiB. To not specify a storage capacity quota, set this to -1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#storage_capacity_quota_gi_b FsxVolume#storage_capacity_quota_gi_b}
	StorageCapacityQuotaGiB *float64 `field:"optional" json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// Specifies the amount of storage in gibibytes (GiB) to reserve from the parent volume.
	//
	// Setting StorageCapacityReservationGiB guarantees that the specified amount of storage space on the parent volume will always be available for the volume. You can't reserve more storage than the parent volume has. To not specify a storage capacity reservation, set this to 0 or -1. For more information, see Volume properties in the Amazon FSx for OpenZFS User Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#storage_capacity_reservation_gi_b FsxVolume#storage_capacity_reservation_gi_b}
	StorageCapacityReservationGiB *float64 `field:"optional" json:"storageCapacityReservationGiB" yaml:"storageCapacityReservationGiB"`
	// Configures how much storage users and groups can use on the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#user_and_group_quotas FsxVolume#user_and_group_quotas}
	UserAndGroupQuotas interface{} `field:"optional" json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
}

