// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicediscoverypublicdnsnamespace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicediscoveryPublicDnsNamespaceConfig struct {
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
	// The name that you want to assign to this namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicediscovery_public_dns_namespace#name ServicediscoveryPublicDnsNamespace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description for the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicediscovery_public_dns_namespace#description ServicediscoveryPublicDnsNamespace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Properties for the public DNS namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicediscovery_public_dns_namespace#properties ServicediscoveryPublicDnsNamespace#properties}
	Properties *ServicediscoveryPublicDnsNamespaceProperties `field:"optional" json:"properties" yaml:"properties"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicediscovery_public_dns_namespace#tags ServicediscoveryPublicDnsNamespace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

