// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinesswebexperience

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QbusinessWebExperienceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#application_id QbusinessWebExperience#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#browser_extension_configuration QbusinessWebExperience#browser_extension_configuration}.
	BrowserExtensionConfiguration *QbusinessWebExperienceBrowserExtensionConfiguration `field:"optional" json:"browserExtensionConfiguration" yaml:"browserExtensionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#customization_configuration QbusinessWebExperience#customization_configuration}.
	CustomizationConfiguration *QbusinessWebExperienceCustomizationConfiguration `field:"optional" json:"customizationConfiguration" yaml:"customizationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#identity_provider_configuration QbusinessWebExperience#identity_provider_configuration}.
	IdentityProviderConfiguration *QbusinessWebExperienceIdentityProviderConfiguration `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#origins QbusinessWebExperience#origins}.
	Origins *[]*string `field:"optional" json:"origins" yaml:"origins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#role_arn QbusinessWebExperience#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#sample_prompts_control_mode QbusinessWebExperience#sample_prompts_control_mode}.
	SamplePromptsControlMode *string `field:"optional" json:"samplePromptsControlMode" yaml:"samplePromptsControlMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#subtitle QbusinessWebExperience#subtitle}.
	Subtitle *string `field:"optional" json:"subtitle" yaml:"subtitle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#tags QbusinessWebExperience#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#title QbusinessWebExperience#title}.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/qbusiness_web_experience#welcome_message QbusinessWebExperience#welcome_message}.
	WelcomeMessage *string `field:"optional" json:"welcomeMessage" yaml:"welcomeMessage"`
}

