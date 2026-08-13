// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsSpaceStorageSettingsEbsStorageSettings struct {
	// Size of the Amazon EBS volume in Gb.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_space#ebs_volume_size_in_gb SagemakerSpace#ebs_volume_size_in_gb}
	EbsVolumeSizeInGb *float64 `field:"optional" json:"ebsVolumeSizeInGb" yaml:"ebsVolumeSizeInGb"`
}

