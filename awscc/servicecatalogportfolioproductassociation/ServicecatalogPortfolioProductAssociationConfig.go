// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogportfolioproductassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogPortfolioProductAssociationConfig struct {
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
	// The language code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_portfolio_product_association#accept_language ServicecatalogPortfolioProductAssociation#accept_language}
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The portfolio identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_portfolio_product_association#portfolio_id ServicecatalogPortfolioProductAssociation#portfolio_id}
	PortfolioId *string `field:"optional" json:"portfolioId" yaml:"portfolioId"`
	// The product identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_portfolio_product_association#product_id ServicecatalogPortfolioProductAssociation#product_id}
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// The identifier of the source portfolio.
	//
	// The source portfolio must be a portfolio imported from a different account than the one creating the association. This account must have previously shared this portfolio with the account creating the association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/servicecatalog_portfolio_product_association#source_portfolio_id ServicecatalogPortfolioProductAssociation#source_portfolio_id}
	SourcePortfolioId *string `field:"optional" json:"sourcePortfolioId" yaml:"sourcePortfolioId"`
}

