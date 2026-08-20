// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentComputeResourcesManagedInstancesProviderInstanceLaunchTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#capacity_option_type BatchComputeEnvironment#capacity_option_type}.
	CapacityOptionType *string `field:"optional" json:"capacityOptionType" yaml:"capacityOptionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#capacity_reservations BatchComputeEnvironment#capacity_reservations}.
	CapacityReservations *BatchComputeEnvironmentComputeResourcesManagedInstancesProviderInstanceLaunchTemplateCapacityReservations `field:"optional" json:"capacityReservations" yaml:"capacityReservations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#ec_2_instance_profile_arn BatchComputeEnvironment#ec_2_instance_profile_arn}.
	Ec2InstanceProfileArn *string `field:"optional" json:"ec2InstanceProfileArn" yaml:"ec2InstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#fips_enabled BatchComputeEnvironment#fips_enabled}.
	FipsEnabled interface{} `field:"optional" json:"fipsEnabled" yaml:"fipsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#instance_metadata_tags_propagation BatchComputeEnvironment#instance_metadata_tags_propagation}.
	InstanceMetadataTagsPropagation interface{} `field:"optional" json:"instanceMetadataTagsPropagation" yaml:"instanceMetadataTagsPropagation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#instance_requirements BatchComputeEnvironment#instance_requirements}.
	InstanceRequirements *BatchComputeEnvironmentComputeResourcesManagedInstancesProviderInstanceLaunchTemplateInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#local_storage_configuration BatchComputeEnvironment#local_storage_configuration}.
	LocalStorageConfiguration *BatchComputeEnvironmentComputeResourcesManagedInstancesProviderInstanceLaunchTemplateLocalStorageConfiguration `field:"optional" json:"localStorageConfiguration" yaml:"localStorageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#monitoring BatchComputeEnvironment#monitoring}.
	Monitoring *string `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#network_configuration BatchComputeEnvironment#network_configuration}.
	NetworkConfiguration *BatchComputeEnvironmentComputeResourcesManagedInstancesProviderInstanceLaunchTemplateNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/batch_compute_environment#storage_configuration BatchComputeEnvironment#storage_configuration}.
	StorageConfiguration *BatchComputeEnvironmentComputeResourcesManagedInstancesProviderInstanceLaunchTemplateStorageConfiguration `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
}

