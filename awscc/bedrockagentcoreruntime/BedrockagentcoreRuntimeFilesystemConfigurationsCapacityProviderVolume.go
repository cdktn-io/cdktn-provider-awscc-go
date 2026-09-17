// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeFilesystemConfigurationsCapacityProviderVolume struct {
	// Mount path for filesystem configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#mount_path BedrockagentcoreRuntime#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
	// Name of the capacity provider volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#volume_name BedrockagentcoreRuntime#volume_name}
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
}

