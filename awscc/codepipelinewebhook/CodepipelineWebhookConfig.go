// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinewebhook

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CodepipelineWebhookConfig struct {
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
	// Supported options are GITHUB_HMAC, IP, and UNAUTHENTICATED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_webhook#authentication CodepipelineWebhook#authentication}
	Authentication *string `field:"required" json:"authentication" yaml:"authentication"`
	// Properties that configure the authentication applied to incoming webhook trigger requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_webhook#authentication_configuration CodepipelineWebhook#authentication_configuration}
	AuthenticationConfiguration *CodepipelineWebhookAuthenticationConfiguration `field:"required" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// A list of rules applied to the body/payload sent in the POST request to a webhook URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_webhook#filters CodepipelineWebhook#filters}
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
	// The name of the action in a pipeline you want to connect to the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_webhook#target_action CodepipelineWebhook#target_action}
	TargetAction *string `field:"required" json:"targetAction" yaml:"targetAction"`
	// The name of the pipeline you want to connect to the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_webhook#target_pipeline CodepipelineWebhook#target_pipeline}
	TargetPipeline *string `field:"required" json:"targetPipeline" yaml:"targetPipeline"`
	// The name of the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_webhook#name CodepipelineWebhook#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configures a connection between the webhook that was created and the external tool with events to be detected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_webhook#register_with_third_party CodepipelineWebhook#register_with_third_party}
	RegisterWithThirdParty interface{} `field:"optional" json:"registerWithThirdParty" yaml:"registerWithThirdParty"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_webhook#tags CodepipelineWebhook#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The version number of the pipeline to be connected to the trigger request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codepipeline_webhook#target_pipeline_version CodepipelineWebhook#target_pipeline_version}
	TargetPipelineVersion *float64 `field:"optional" json:"targetPipelineVersion" yaml:"targetPipelineVersion"`
}

