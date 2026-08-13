// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncapi

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppsyncApiConfig struct {
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
	// The name of the AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_api#name AppsyncApi#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The configuration for an Event Api.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_api#event_config AppsyncApi#event_config}
	EventConfig *AppsyncApiEventConfig `field:"optional" json:"eventConfig" yaml:"eventConfig"`
	// The owner contact information for an API resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_api#owner_contact AppsyncApi#owner_contact}
	OwnerContact *string `field:"optional" json:"ownerContact" yaml:"ownerContact"`
	// An arbitrary set of tags (key-value pairs) for this AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/appsync_api#tags AppsyncApi#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

