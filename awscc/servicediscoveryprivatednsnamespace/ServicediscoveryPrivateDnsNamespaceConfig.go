// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicediscoveryprivatednsnamespace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicediscoveryPrivateDnsNamespaceConfig struct {
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
	// When you create a private DNS namespace, AWS Cloud Map automatically creates an Amazon Route 53 private hosted zone that has the same name as the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/servicediscovery_private_dns_namespace#name ServicediscoveryPrivateDnsNamespace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of the namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/servicediscovery_private_dns_namespace#description ServicediscoveryPrivateDnsNamespace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Properties of the private DNS namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/servicediscovery_private_dns_namespace#properties ServicediscoveryPrivateDnsNamespace#properties}
	Properties *ServicediscoveryPrivateDnsNamespaceProperties `field:"optional" json:"properties" yaml:"properties"`
	// The tags for the namespace.
	//
	// Each tag consists of a key and an optional value, both of which you define. Tag keys can have a maximum character length of 128 characters, and tag values can have a maximum length of 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/servicediscovery_private_dns_namespace#tags ServicediscoveryPrivateDnsNamespace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the Amazon VPC that you want to associate the namespace with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/servicediscovery_private_dns_namespace#vpc ServicediscoveryPrivateDnsNamespace#vpc}
	Vpc *string `field:"optional" json:"vpc" yaml:"vpc"`
}

