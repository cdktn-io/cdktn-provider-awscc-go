// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsSpaceStorageSettings struct {
	// Properties related to the Amazon Elastic Block Store volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_user_profile#default_ebs_storage_settings SagemakerUserProfile#default_ebs_storage_settings}
	DefaultEbsStorageSettings *SagemakerUserProfileUserSettingsSpaceStorageSettingsDefaultEbsStorageSettings `field:"optional" json:"defaultEbsStorageSettings" yaml:"defaultEbsStorageSettings"`
}

