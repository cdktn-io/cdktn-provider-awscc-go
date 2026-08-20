// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogresourceupdateconstraint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogResourceUpdateConstraintConfig struct {
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
	// The portfolio identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_resource_update_constraint#portfolio_id ServicecatalogResourceUpdateConstraint#portfolio_id}
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_resource_update_constraint#product_id ServicecatalogResourceUpdateConstraint#product_id}
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// ALLOWED or NOT_ALLOWED, to permit or prevent changes to the tags on provisioned instances of the specified portfolio / product combination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_resource_update_constraint#tag_update_on_provisioned_product ServicecatalogResourceUpdateConstraint#tag_update_on_provisioned_product}
	TagUpdateOnProvisionedProduct *string `field:"required" json:"tagUpdateOnProvisionedProduct" yaml:"tagUpdateOnProvisionedProduct"`
	// The language code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_resource_update_constraint#accept_language ServicecatalogResourceUpdateConstraint#accept_language}
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The description of the constraint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/servicecatalog_resource_update_constraint#description ServicecatalogResourceUpdateConstraint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

