// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsSpaceStorageSettings struct {
	// Properties related to the space's Amazon Elastic Block Store volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_space#ebs_storage_settings SagemakerSpace#ebs_storage_settings}
	EbsStorageSettings *SagemakerSpaceSpaceSettingsSpaceStorageSettingsEbsStorageSettings `field:"optional" json:"ebsStorageSettings" yaml:"ebsStorageSettings"`
}

