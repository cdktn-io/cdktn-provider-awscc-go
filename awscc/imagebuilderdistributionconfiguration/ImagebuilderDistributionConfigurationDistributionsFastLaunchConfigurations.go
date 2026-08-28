// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionsFastLaunchConfigurations struct {
	// The owner account ID for the fast-launch enabled Windows AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#account_id ImagebuilderDistributionConfiguration#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// A Boolean that represents the current state of faster launching for the Windows AMI.
	//
	// Set to true to start using Windows faster launching, or false to stop using it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#enabled ImagebuilderDistributionConfiguration#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The launch template that the fast-launch enabled Windows AMI uses when it launches Windows instances to create pre-provisioned snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#launch_template ImagebuilderDistributionConfiguration#launch_template}
	LaunchTemplate *ImagebuilderDistributionConfigurationDistributionsFastLaunchConfigurationsLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// The maximum number of parallel instances that are launched for creating resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#max_parallel_launches ImagebuilderDistributionConfiguration#max_parallel_launches}
	MaxParallelLaunches *float64 `field:"optional" json:"maxParallelLaunches" yaml:"maxParallelLaunches"`
	// Configuration settings for managing the number of snapshots that are created from pre-provisioned instances for the Windows AMI when faster launching is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_distribution_configuration#snapshot_configuration ImagebuilderDistributionConfiguration#snapshot_configuration}
	SnapshotConfiguration *ImagebuilderDistributionConfigurationDistributionsFastLaunchConfigurationsSnapshotConfiguration `field:"optional" json:"snapshotConfiguration" yaml:"snapshotConfiguration"`
}

