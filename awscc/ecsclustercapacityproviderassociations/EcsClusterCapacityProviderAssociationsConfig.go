// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsclustercapacityproviderassociations

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsClusterCapacityProviderAssociationsConfig struct {
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
	// The name of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_cluster_capacity_provider_associations#cluster EcsClusterCapacityProviderAssociations#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// List of capacity providers to associate with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_cluster_capacity_provider_associations#default_capacity_provider_strategy EcsClusterCapacityProviderAssociations#default_capacity_provider_strategy}
	DefaultCapacityProviderStrategy interface{} `field:"required" json:"defaultCapacityProviderStrategy" yaml:"defaultCapacityProviderStrategy"`
	// List of capacity providers to associate with the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ecs_cluster_capacity_provider_associations#capacity_providers EcsClusterCapacityProviderAssociations#capacity_providers}
	CapacityProviders *[]*string `field:"optional" json:"capacityProviders" yaml:"capacityProviders"`
}

