// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeFilesystemConfigurations struct {
	// Configuration for EFS access point filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_runtime#efs_access_point BedrockagentcoreRuntime#efs_access_point}
	EfsAccessPoint *BedrockagentcoreRuntimeFilesystemConfigurationsEfsAccessPoint `field:"optional" json:"efsAccessPoint" yaml:"efsAccessPoint"`
	// Configuration for S3 Files access point filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_runtime#s3_files_access_point BedrockagentcoreRuntime#s3_files_access_point}
	S3FilesAccessPoint *BedrockagentcoreRuntimeFilesystemConfigurationsS3FilesAccessPoint `field:"optional" json:"s3FilesAccessPoint" yaml:"s3FilesAccessPoint"`
	// Configuration for session storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/bedrockagentcore_runtime#session_storage BedrockagentcoreRuntime#session_storage}
	SessionStorage *BedrockagentcoreRuntimeFilesystemConfigurationsSessionStorage `field:"optional" json:"sessionStorage" yaml:"sessionStorage"`
}

