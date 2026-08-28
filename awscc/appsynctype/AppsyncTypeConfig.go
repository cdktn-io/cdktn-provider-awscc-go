// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsynctype

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncTypeConfig struct {
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
	// The API ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_type#api_id AppsyncType#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The type definition, in GraphQL Schema Definition Language (SDL) format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_type#definition AppsyncType#definition}
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// The type format: SDL or JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appsync_type#format AppsyncType#format}
	Format *string `field:"required" json:"format" yaml:"format"`
}

