// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsCustomFileSystemConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_user_profile#efs_file_system_config SagemakerUserProfile#efs_file_system_config}.
	EfsFileSystemConfig *SagemakerUserProfileUserSettingsCustomFileSystemConfigsEfsFileSystemConfig `field:"optional" json:"efsFileSystemConfig" yaml:"efsFileSystemConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_user_profile#fsx_lustre_file_system_config SagemakerUserProfile#fsx_lustre_file_system_config}.
	FsxLustreFileSystemConfig *SagemakerUserProfileUserSettingsCustomFileSystemConfigsFsxLustreFileSystemConfig `field:"optional" json:"fsxLustreFileSystemConfig" yaml:"fsxLustreFileSystemConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_user_profile#s3_file_system_config SagemakerUserProfile#s3_file_system_config}.
	S3FileSystemConfig *SagemakerUserProfileUserSettingsCustomFileSystemConfigsS3FileSystemConfig `field:"optional" json:"s3FileSystemConfig" yaml:"s3FileSystemConfig"`
}

