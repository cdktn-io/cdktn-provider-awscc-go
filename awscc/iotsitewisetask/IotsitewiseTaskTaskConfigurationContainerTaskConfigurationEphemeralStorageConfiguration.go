// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisetask


type IotsitewiseTaskTaskConfigurationContainerTaskConfigurationEphemeralStorageConfiguration struct {
	// The storage type that determines I/O performance characteristics.
	//
	// Family name indicates workload pattern, level number indicates performance within that family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_task#storage_class IotsitewiseTask#storage_class}
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Storage volume size in GiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotsitewise_task#storage_size_in_gi_b IotsitewiseTask#storage_size_in_gi_b}
	StorageSizeInGiB *float64 `field:"optional" json:"storageSizeInGiB" yaml:"storageSizeInGiB"`
}

