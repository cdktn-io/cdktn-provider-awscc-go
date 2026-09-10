// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_domain#efs_file_system_config SagemakerDomain#efs_file_system_config}.
	EfsFileSystemConfig *SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsEfsFileSystemConfig `field:"optional" json:"efsFileSystemConfig" yaml:"efsFileSystemConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_domain#fsx_lustre_file_system_config SagemakerDomain#fsx_lustre_file_system_config}.
	FsxLustreFileSystemConfig *SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsFsxLustreFileSystemConfig `field:"optional" json:"fsxLustreFileSystemConfig" yaml:"fsxLustreFileSystemConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_domain#s3_file_system_config SagemakerDomain#s3_file_system_config}.
	S3FileSystemConfig *SagemakerDomainDefaultSpaceSettingsCustomFileSystemConfigsS3FileSystemConfig `field:"optional" json:"s3FileSystemConfig" yaml:"s3FileSystemConfig"`
}

