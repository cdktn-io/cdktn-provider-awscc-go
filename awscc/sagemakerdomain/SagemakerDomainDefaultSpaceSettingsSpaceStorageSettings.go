// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultSpaceSettingsSpaceStorageSettings struct {
	// Properties related to the Amazon Elastic Block Store volume.
	//
	// Must be provided if storage type is Amazon EBS and must not be provided if storage type is not Amazon EBS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_domain#default_ebs_storage_settings SagemakerDomain#default_ebs_storage_settings}
	DefaultEbsStorageSettings *SagemakerDomainDefaultSpaceSettingsSpaceStorageSettingsDefaultEbsStorageSettings `field:"optional" json:"defaultEbsStorageSettings" yaml:"defaultEbsStorageSettings"`
}

