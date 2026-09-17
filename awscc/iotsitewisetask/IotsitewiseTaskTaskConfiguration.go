// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewisetask


type IotsitewiseTaskTaskConfiguration struct {
	// Configuration for running a custom container image on managed compute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotsitewise_task#container_task_configuration IotsitewiseTask#container_task_configuration}
	ContainerTaskConfiguration *IotsitewiseTaskTaskConfigurationContainerTaskConfiguration `field:"required" json:"containerTaskConfiguration" yaml:"containerTaskConfiguration"`
}

