// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheparametergroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticacheParameterGroupConfig struct {
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
	// The name of the cache parameter group family that this cache parameter group is compatible with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticache_parameter_group#cache_parameter_group_family ElasticacheParameterGroup#cache_parameter_group_family}
	CacheParameterGroupFamily *string `field:"required" json:"cacheParameterGroupFamily" yaml:"cacheParameterGroupFamily"`
	// The description for this cache parameter group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticache_parameter_group#description ElasticacheParameterGroup#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// A comma-delimited list of parameter name/value pairs. For more information see ModifyCacheParameterGroup in the Amazon ElastiCache API Reference Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticache_parameter_group#properties ElasticacheParameterGroup#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Tags are composed of a Key/Value pair.
	//
	// You can use tags to categorize and track each parameter group. The tag value null is permitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elasticache_parameter_group#tags ElasticacheParameterGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

