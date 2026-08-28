// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeFilesystemConfigurationsS3FilesAccessPoint struct {
	// ARN of the S3 Files access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_runtime#access_point_arn BedrockagentcoreRuntime#access_point_arn}
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// Mount path for filesystem configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_runtime#mount_path BedrockagentcoreRuntime#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

