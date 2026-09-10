// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentComputeResources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#allocation_strategy BatchComputeEnvironment#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#bid_percentage BatchComputeEnvironment#bid_percentage}.
	BidPercentage *float64 `field:"optional" json:"bidPercentage" yaml:"bidPercentage"`
	// Capacity-level tags for compute environments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#capacity_tags BatchComputeEnvironment#capacity_tags}
	CapacityTags *map[string]*string `field:"optional" json:"capacityTags" yaml:"capacityTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#desiredv_cpus BatchComputeEnvironment#desiredv_cpus}.
	DesiredvCpus *float64 `field:"optional" json:"desiredvCpus" yaml:"desiredvCpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#ec_2_configuration BatchComputeEnvironment#ec_2_configuration}.
	Ec2Configuration interface{} `field:"optional" json:"ec2Configuration" yaml:"ec2Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#ec_2_key_pair BatchComputeEnvironment#ec_2_key_pair}.
	Ec2KeyPair *string `field:"optional" json:"ec2KeyPair" yaml:"ec2KeyPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#image_id BatchComputeEnvironment#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#instance_role BatchComputeEnvironment#instance_role}.
	InstanceRole *string `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#instance_types BatchComputeEnvironment#instance_types}.
	InstanceTypes *[]*string `field:"optional" json:"instanceTypes" yaml:"instanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#launch_template BatchComputeEnvironment#launch_template}.
	LaunchTemplate *BatchComputeEnvironmentComputeResourcesLaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#managed_instances_provider BatchComputeEnvironment#managed_instances_provider}.
	ManagedInstancesProvider *BatchComputeEnvironmentComputeResourcesManagedInstancesProvider `field:"optional" json:"managedInstancesProvider" yaml:"managedInstancesProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#maxv_cpus BatchComputeEnvironment#maxv_cpus}.
	MaxvCpus *float64 `field:"optional" json:"maxvCpus" yaml:"maxvCpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#minv_cpus BatchComputeEnvironment#minv_cpus}.
	MinvCpus *float64 `field:"optional" json:"minvCpus" yaml:"minvCpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#placement_group BatchComputeEnvironment#placement_group}.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#scaling_policy BatchComputeEnvironment#scaling_policy}.
	ScalingPolicy *BatchComputeEnvironmentComputeResourcesScalingPolicy `field:"optional" json:"scalingPolicy" yaml:"scalingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#security_group_ids BatchComputeEnvironment#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#spot_iam_fleet_role BatchComputeEnvironment#spot_iam_fleet_role}.
	SpotIamFleetRole *string `field:"optional" json:"spotIamFleetRole" yaml:"spotIamFleetRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#subnets BatchComputeEnvironment#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#tags BatchComputeEnvironment#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#type BatchComputeEnvironment#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_compute_environment#update_to_latest_image_version BatchComputeEnvironment#update_to_latest_image_version}.
	UpdateToLatestImageVersion interface{} `field:"optional" json:"updateToLatestImageVersion" yaml:"updateToLatestImageVersion"`
}

