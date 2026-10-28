// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsgatewayclusterassociations

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsGatewayClusterAssociationsConfig struct {
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
	// The name or full Amazon Resource Name (ARN) of the cluster to associate the gateway with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway_cluster_associations#cluster EcsGatewayClusterAssociations#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The name or full Amazon Resource Name (ARN) of the gateway to associate with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ecs_gateway_cluster_associations#gateway EcsGatewayClusterAssociations#gateway}
	Gateway *string `field:"required" json:"gateway" yaml:"gateway"`
}

