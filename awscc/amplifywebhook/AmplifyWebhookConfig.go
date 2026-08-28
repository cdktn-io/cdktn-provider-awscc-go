// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package amplifywebhook

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AmplifyWebhookConfig struct {
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
	// The name for a branch that is part of an Amplify app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_webhook#branch_name AmplifyWebhook#branch_name}
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// The unique ID for an Amplify app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_webhook#app_id AmplifyWebhook#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The description for a webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_webhook#description AmplifyWebhook#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags for the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/amplify_webhook#tags AmplifyWebhook#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

