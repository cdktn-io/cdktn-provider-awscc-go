// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogportfolioshare

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogPortfolioShareConfig struct {
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
	// The AWS account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_portfolio_share#account_id ServicecatalogPortfolioShare#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The portfolio identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_portfolio_share#portfolio_id ServicecatalogPortfolioShare#portfolio_id}
	PortfolioId *string `field:"required" json:"portfolioId" yaml:"portfolioId"`
	// The language code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_portfolio_share#accept_language ServicecatalogPortfolioShare#accept_language}
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// Enables or disables TagOptions sharing when creating the portfolio share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_portfolio_share#share_tag_options ServicecatalogPortfolioShare#share_tag_options}
	ShareTagOptions interface{} `field:"optional" json:"shareTagOptions" yaml:"shareTagOptions"`
}

