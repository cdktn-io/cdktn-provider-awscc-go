// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codebuildfleet

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CodebuildFleetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#base_capacity CodebuildFleet#base_capacity}.
	BaseCapacity *float64 `field:"optional" json:"baseCapacity" yaml:"baseCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#compute_configuration CodebuildFleet#compute_configuration}.
	ComputeConfiguration *CodebuildFleetComputeConfiguration `field:"optional" json:"computeConfiguration" yaml:"computeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#compute_type CodebuildFleet#compute_type}.
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#environment_type CodebuildFleet#environment_type}.
	EnvironmentType *string `field:"optional" json:"environmentType" yaml:"environmentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#fleet_proxy_configuration CodebuildFleet#fleet_proxy_configuration}.
	FleetProxyConfiguration *CodebuildFleetFleetProxyConfiguration `field:"optional" json:"fleetProxyConfiguration" yaml:"fleetProxyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#fleet_service_role CodebuildFleet#fleet_service_role}.
	FleetServiceRole *string `field:"optional" json:"fleetServiceRole" yaml:"fleetServiceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#fleet_vpc_config CodebuildFleet#fleet_vpc_config}.
	FleetVpcConfig *CodebuildFleetFleetVpcConfig `field:"optional" json:"fleetVpcConfig" yaml:"fleetVpcConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#image_id CodebuildFleet#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#name CodebuildFleet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#overflow_behavior CodebuildFleet#overflow_behavior}.
	OverflowBehavior *string `field:"optional" json:"overflowBehavior" yaml:"overflowBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#scaling_configuration CodebuildFleet#scaling_configuration}.
	ScalingConfiguration *CodebuildFleetScalingConfiguration `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codebuild_fleet#tags CodebuildFleet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

