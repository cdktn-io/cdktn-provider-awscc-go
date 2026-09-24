// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscluster


type EcsClusterConfiguration struct {
	// The details of the execute command configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_cluster#execute_command_configuration EcsCluster#execute_command_configuration}
	ExecuteCommandConfiguration *EcsClusterConfigurationExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// The details of the managed storage configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ecs_cluster#managed_storage_configuration EcsCluster#managed_storage_configuration}
	ManagedStorageConfiguration *EcsClusterConfigurationManagedStorageConfiguration `field:"optional" json:"managedStorageConfiguration" yaml:"managedStorageConfiguration"`
}

