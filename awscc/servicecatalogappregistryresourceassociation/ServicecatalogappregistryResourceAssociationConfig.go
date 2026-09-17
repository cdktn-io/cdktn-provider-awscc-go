// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogappregistryresourceassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogappregistryResourceAssociationConfig struct {
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
	// The name or the Id of the Application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalogappregistry_resource_association#application ServicecatalogappregistryResourceAssociation#application}
	Application *string `field:"required" json:"application" yaml:"application"`
	// The name or the Id of the Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalogappregistry_resource_association#resource ServicecatalogappregistryResourceAssociation#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The type of the CFN Resource for now it's enum CFN_STACK.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalogappregistry_resource_association#resource_type ServicecatalogappregistryResourceAssociation#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

