// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2spotfleet


type Ec2SpotFleetSpotFleetRequestConfigData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#iam_fleet_role Ec2SpotFleet#iam_fleet_role}.
	IamFleetRole *string `field:"required" json:"iamFleetRole" yaml:"iamFleetRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#target_capacity Ec2SpotFleet#target_capacity}.
	TargetCapacity *float64 `field:"required" json:"targetCapacity" yaml:"targetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#allocation_strategy Ec2SpotFleet#allocation_strategy}.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#context Ec2SpotFleet#context}.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#excess_capacity_termination_policy Ec2SpotFleet#excess_capacity_termination_policy}.
	ExcessCapacityTerminationPolicy *string `field:"optional" json:"excessCapacityTerminationPolicy" yaml:"excessCapacityTerminationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#instance_interruption_behavior Ec2SpotFleet#instance_interruption_behavior}.
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#instance_pools_to_use_count Ec2SpotFleet#instance_pools_to_use_count}.
	InstancePoolsToUseCount *float64 `field:"optional" json:"instancePoolsToUseCount" yaml:"instancePoolsToUseCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#launch_specifications Ec2SpotFleet#launch_specifications}.
	LaunchSpecifications interface{} `field:"optional" json:"launchSpecifications" yaml:"launchSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#launch_template_configs Ec2SpotFleet#launch_template_configs}.
	LaunchTemplateConfigs interface{} `field:"optional" json:"launchTemplateConfigs" yaml:"launchTemplateConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#load_balancers_config Ec2SpotFleet#load_balancers_config}.
	LoadBalancersConfig *Ec2SpotFleetSpotFleetRequestConfigDataLoadBalancersConfig `field:"optional" json:"loadBalancersConfig" yaml:"loadBalancersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#on_demand_allocation_strategy Ec2SpotFleet#on_demand_allocation_strategy}.
	OnDemandAllocationStrategy *string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#on_demand_max_total_price Ec2SpotFleet#on_demand_max_total_price}.
	OnDemandMaxTotalPrice *string `field:"optional" json:"onDemandMaxTotalPrice" yaml:"onDemandMaxTotalPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#on_demand_target_capacity Ec2SpotFleet#on_demand_target_capacity}.
	OnDemandTargetCapacity *float64 `field:"optional" json:"onDemandTargetCapacity" yaml:"onDemandTargetCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#replace_unhealthy_instances Ec2SpotFleet#replace_unhealthy_instances}.
	ReplaceUnhealthyInstances interface{} `field:"optional" json:"replaceUnhealthyInstances" yaml:"replaceUnhealthyInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#spot_maintenance_strategies Ec2SpotFleet#spot_maintenance_strategies}.
	SpotMaintenanceStrategies *Ec2SpotFleetSpotFleetRequestConfigDataSpotMaintenanceStrategies `field:"optional" json:"spotMaintenanceStrategies" yaml:"spotMaintenanceStrategies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#spot_max_total_price Ec2SpotFleet#spot_max_total_price}.
	SpotMaxTotalPrice *string `field:"optional" json:"spotMaxTotalPrice" yaml:"spotMaxTotalPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#spot_price Ec2SpotFleet#spot_price}.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#tag_specifications Ec2SpotFleet#tag_specifications}.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#target_capacity_unit_type Ec2SpotFleet#target_capacity_unit_type}.
	TargetCapacityUnitType *string `field:"optional" json:"targetCapacityUnitType" yaml:"targetCapacityUnitType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#terminate_instances_with_expiration Ec2SpotFleet#terminate_instances_with_expiration}.
	TerminateInstancesWithExpiration interface{} `field:"optional" json:"terminateInstancesWithExpiration" yaml:"terminateInstancesWithExpiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#type Ec2SpotFleet#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#valid_from Ec2SpotFleet#valid_from}.
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ec2_spot_fleet#valid_until Ec2SpotFleet#valid_until}.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}

