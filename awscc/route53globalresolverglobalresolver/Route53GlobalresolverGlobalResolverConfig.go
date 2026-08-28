// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53globalresolverglobalresolver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53GlobalresolverGlobalResolverConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53globalresolver_global_resolver#name Route53GlobalresolverGlobalResolver#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The list of regions the Global Resolver exists in.
	//
	// Regions can be added or removed on update; the order of this list is not significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53globalresolver_global_resolver#regions Route53GlobalresolverGlobalResolver#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53globalresolver_global_resolver#client_token Route53GlobalresolverGlobalResolver#client_token}.
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53globalresolver_global_resolver#description Route53GlobalresolverGlobalResolver#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53globalresolver_global_resolver#ip_address_type Route53GlobalresolverGlobalResolver#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53globalresolver_global_resolver#observability_region Route53GlobalresolverGlobalResolver#observability_region}.
	ObservabilityRegion *string `field:"optional" json:"observabilityRegion" yaml:"observabilityRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53globalresolver_global_resolver#tags Route53GlobalresolverGlobalResolver#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

