// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogappregistryattributegroupassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogappregistryAttributeGroupAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalogappregistry_attribute_group_association#application ServicecatalogappregistryAttributeGroupAssociation#application}
	Application *string `field:"required" json:"application" yaml:"application"`
	// The name or the Id of the AttributeGroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalogappregistry_attribute_group_association#attribute_group ServicecatalogappregistryAttributeGroupAssociation#attribute_group}
	AttributeGroup *string `field:"required" json:"attributeGroup" yaml:"attributeGroup"`
}

