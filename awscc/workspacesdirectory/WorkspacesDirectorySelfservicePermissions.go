// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectorySelfservicePermissions struct {
	// Specifies whether users can change the compute type (bundle) for their WorkSpace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#change_compute_type WorkspacesDirectory#change_compute_type}
	ChangeComputeType *string `field:"optional" json:"changeComputeType" yaml:"changeComputeType"`
	// Specifies whether users can increase the volume size of the drives on their WorkSpace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#increase_volume_size WorkspacesDirectory#increase_volume_size}
	IncreaseVolumeSize *string `field:"optional" json:"increaseVolumeSize" yaml:"increaseVolumeSize"`
	// Specifies whether users can rebuild the operating system of a WorkSpace to its original state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#rebuild_workspace WorkspacesDirectory#rebuild_workspace}
	RebuildWorkspace *string `field:"optional" json:"rebuildWorkspace" yaml:"rebuildWorkspace"`
	// Specifies whether users can restart their WorkSpace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#restart_workspace WorkspacesDirectory#restart_workspace}
	RestartWorkspace *string `field:"optional" json:"restartWorkspace" yaml:"restartWorkspace"`
	// Specifies whether users can switch the running mode of their WorkSpace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/workspaces_directory#switch_running_mode WorkspacesDirectory#switch_running_mode}
	SwitchRunningMode *string `field:"optional" json:"switchRunningMode" yaml:"switchRunningMode"`
}

