// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscapacityprovider


type EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#capacity_option_type EcsCapacityProvider#capacity_option_type}.
	CapacityOptionType *string `field:"optional" json:"capacityOptionType" yaml:"capacityOptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#capacity_reservations EcsCapacityProvider#capacity_reservations}.
	CapacityReservations *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateCapacityReservations `field:"optional" json:"capacityReservations" yaml:"capacityReservations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#ec_2_instance_profile_arn EcsCapacityProvider#ec_2_instance_profile_arn}.
	Ec2InstanceProfileArn *string `field:"optional" json:"ec2InstanceProfileArn" yaml:"ec2InstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#fips_enabled EcsCapacityProvider#fips_enabled}.
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#instance_metadata_tags_propagation EcsCapacityProvider#instance_metadata_tags_propagation}.
	InstanceMetadataTagsPropagation interface{} `field:"optional" json:"instanceMetadataTagsPropagation" yaml:"instanceMetadataTagsPropagation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#instance_requirements EcsCapacityProvider#instance_requirements}.
	InstanceRequirements *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#local_storage_configuration EcsCapacityProvider#local_storage_configuration}.
	LocalStorageConfiguration *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateLocalStorageConfiguration `field:"optional" json:"localStorageConfiguration" yaml:"localStorageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#monitoring EcsCapacityProvider#monitoring}.
	Monitoring *string `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#network_configuration EcsCapacityProvider#network_configuration}.
	NetworkConfiguration *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ecs_capacity_provider#storage_configuration EcsCapacityProvider#storage_configuration}.
	StorageConfiguration *EcsCapacityProviderManagedInstancesProviderInstanceLaunchTemplateStorageConfiguration `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

