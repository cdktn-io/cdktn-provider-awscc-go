// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package servicecatalogportfolioprincipalassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ServicecatalogPortfolioPrincipalAssociationConfig struct {
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
	// The principal type.
	//
	// The supported value is IAM if you use a fully defined Amazon Resource Name (ARN), or IAM_PATTERN if you use an ARN with no accountID, with or without wildcard characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_portfolio_principal_association#principal_type ServicecatalogPortfolioPrincipalAssociation#principal_type}
	PrincipalType *string `field:"required" json:"principalType" yaml:"principalType"`
	// The language code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_portfolio_principal_association#accept_language ServicecatalogPortfolioPrincipalAssociation#accept_language}
	AcceptLanguage *string `field:"optional" json:"acceptLanguage" yaml:"acceptLanguage"`
	// The portfolio identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_portfolio_principal_association#portfolio_id ServicecatalogPortfolioPrincipalAssociation#portfolio_id}
	PortfolioId *string `field:"optional" json:"portfolioId" yaml:"portfolioId"`
	// The ARN of the principal (user, role, or group).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/servicecatalog_portfolio_principal_association#principal_arn ServicecatalogPortfolioPrincipalAssociation#principal_arn}
	PrincipalArn *string `field:"optional" json:"principalArn" yaml:"principalArn"`
}

