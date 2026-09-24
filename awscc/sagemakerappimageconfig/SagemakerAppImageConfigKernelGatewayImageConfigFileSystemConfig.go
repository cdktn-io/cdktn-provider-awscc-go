// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerappimageconfig


type SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig struct {
	// The default POSIX group ID (GID). If not specified, defaults to 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_app_image_config#default_gid SagemakerAppImageConfig#default_gid}
	DefaultGid *float64 `field:"optional" json:"defaultGid" yaml:"defaultGid"`
	// The default POSIX user ID (UID). If not specified, defaults to 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_app_image_config#default_uid SagemakerAppImageConfig#default_uid}
	DefaultUid *float64 `field:"optional" json:"defaultUid" yaml:"defaultUid"`
	// The path within the image to mount the user's EFS home directory.
	//
	// The directory should be empty. If not specified, defaults to /home/sagemaker-user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_app_image_config#mount_path SagemakerAppImageConfig#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

