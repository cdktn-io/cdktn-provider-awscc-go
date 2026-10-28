// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsGatewayConfig struct {
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
	// The configuration for the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#configuration EcsGateway#configuration}
	Configuration *EcsGatewayConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The name of the gateway. Up to 255 letters (uppercase and lowercase), numbers, and hyphens are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#gateway_name EcsGateway#gateway_name}
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// The ARN of the IAM role that ECS uses to manage infrastructure for the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#infrastructure_role_arn EcsGateway#infrastructure_role_arn}
	InfrastructureRoleArn *string `field:"required" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// Whether deletion protection is enabled for the gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#deletion_protection EcsGateway#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Specifies whether to propagate the tags from the gateway to other resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#propagate_tags EcsGateway#propagate_tags}
	PropagateTags *string `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// The metadata that you apply to the gateway to help you categorize and organize them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway#tags EcsGateway#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

