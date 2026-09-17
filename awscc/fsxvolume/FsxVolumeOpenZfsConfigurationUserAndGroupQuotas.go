// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOpenZfsConfigurationUserAndGroupQuotas struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fsx_volume#id FsxVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *float64 `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fsx_volume#storage_capacity_quota_gi_b FsxVolume#storage_capacity_quota_gi_b}.
	StorageCapacityQuotaGiB *float64 `field:"optional" json:"storageCapacityQuotaGiB" yaml:"storageCapacityQuotaGiB"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fsx_volume#type FsxVolume#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

