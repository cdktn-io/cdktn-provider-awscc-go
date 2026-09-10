// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbinstance


type RdsDbInstanceAdditionalStorageVolumes struct {
	// The amount of storage allocated for the additional storage volume, in gibibytes (GiB).
	//
	// The minimum is 20 GiB. The maximum is 65,536 GiB (64 TiB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_db_instance#allocated_storage RdsDbInstance#allocated_storage}
	AllocatedStorage *string `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// The number of I/O operations per second (IOPS) provisioned for the additional storage volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_db_instance#iops RdsDbInstance#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The upper limit in gibibytes (GiB) to which RDS can automatically scale the storage of the additional storage volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_db_instance#max_allocated_storage RdsDbInstance#max_allocated_storage}
	MaxAllocatedStorage *float64 `field:"optional" json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_db_instance#storage_operation_percent_progress RdsDbInstance#storage_operation_percent_progress}.
	StorageOperationPercentProgress *float64 `field:"optional" json:"storageOperationPercentProgress" yaml:"storageOperationPercentProgress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_db_instance#storage_operation_status RdsDbInstance#storage_operation_status}.
	StorageOperationStatus *string `field:"optional" json:"storageOperationStatus" yaml:"storageOperationStatus"`
	// The storage throughput value for the additional storage volume, in mebibytes per second (MiBps).
	//
	// This setting applies only to the General Purpose SSD (``gp3``) storage type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_db_instance#storage_throughput RdsDbInstance#storage_throughput}
	StorageThroughput *float64 `field:"optional" json:"storageThroughput" yaml:"storageThroughput"`
	// The storage type for the additional storage volume.  Valid Values: ``GP3 | IO2``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_db_instance#storage_type RdsDbInstance#storage_type}
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// The name of the additional storage volume.  Valid Values: ``RDSDBDATA2 | RDSDBDATA3 | RDSDBDATA4``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_db_instance#volume_name RdsDbInstance#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

