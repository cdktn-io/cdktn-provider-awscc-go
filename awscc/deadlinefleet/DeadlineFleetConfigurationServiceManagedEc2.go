// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet


type DeadlineFleetConfigurationServiceManagedEc2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#auto_scaling_configuration DeadlineFleet#auto_scaling_configuration}.
	AutoScalingConfiguration *DeadlineFleetConfigurationServiceManagedEc2AutoScalingConfiguration `field:"optional" json:"autoScalingConfiguration" yaml:"autoScalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#instance_capabilities DeadlineFleet#instance_capabilities}.
	InstanceCapabilities *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities `field:"optional" json:"instanceCapabilities" yaml:"instanceCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#instance_market_options DeadlineFleet#instance_market_options}.
	InstanceMarketOptions *DeadlineFleetConfigurationServiceManagedEc2InstanceMarketOptions `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#persistent_volume_configuration DeadlineFleet#persistent_volume_configuration}.
	PersistentVolumeConfiguration *DeadlineFleetConfigurationServiceManagedEc2PersistentVolumeConfiguration `field:"optional" json:"persistentVolumeConfiguration" yaml:"persistentVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#storage_profile_id DeadlineFleet#storage_profile_id}.
	StorageProfileId *string `field:"optional" json:"storageProfileId" yaml:"storageProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/deadline_fleet#vpc_configuration DeadlineFleet#vpc_configuration}.
	VpcConfiguration *DeadlineFleetConfigurationServiceManagedEc2VpcConfiguration `field:"optional" json:"vpcConfiguration" yaml:"vpcConfiguration"`
}

