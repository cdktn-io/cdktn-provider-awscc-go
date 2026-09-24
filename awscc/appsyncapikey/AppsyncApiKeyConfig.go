// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncapikey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncApiKeyConfig struct {
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
	// Unique AWS AppSync GraphQL API ID for this API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_api_key#api_id AppsyncApiKey#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Unique description of your API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_api_key#description AppsyncApiKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The time after which the API key expires.
	//
	// The date is represented as seconds since the epoch, rounded down to the nearest hour.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_api_key#expires AppsyncApiKey#expires}
	Expires *float64 `field:"optional" json:"expires" yaml:"expires"`
}

