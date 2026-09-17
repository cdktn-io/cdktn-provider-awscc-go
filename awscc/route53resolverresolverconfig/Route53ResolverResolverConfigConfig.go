// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53resolverresolverconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53ResolverResolverConfigConfig struct {
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
	// Represents the desired status of AutodefinedReverse.
	//
	// The only supported value on creation is DISABLE. Deletion of this resource will return AutodefinedReverse to its default value (ENABLED).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53resolver_resolver_config#autodefined_reverse_flag Route53ResolverResolverConfig#autodefined_reverse_flag}
	AutodefinedReverseFlag *string `field:"required" json:"autodefinedReverseFlag" yaml:"autodefinedReverseFlag"`
	// ResourceId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/route53resolver_resolver_config#resource_id Route53ResolverResolverConfig#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
}

