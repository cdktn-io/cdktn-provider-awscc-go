// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfighostedconfigurationversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppconfigHostedConfigurationVersionConfig struct {
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
	// The application ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_hosted_configuration_version#application_id AppconfigHostedConfigurationVersion#application_id}
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// The configuration profile ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_hosted_configuration_version#configuration_profile_id AppconfigHostedConfigurationVersion#configuration_profile_id}
	ConfigurationProfileId *string `field:"required" json:"configurationProfileId" yaml:"configurationProfileId"`
	// The content of the configuration or the configuration data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_hosted_configuration_version#content AppconfigHostedConfigurationVersion#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// A standard MIME type describing the format of the configuration content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_hosted_configuration_version#content_type AppconfigHostedConfigurationVersion#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// A description of the hosted configuration version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_hosted_configuration_version#description AppconfigHostedConfigurationVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An optional locking token used to prevent race conditions from overwriting configuration updates when creating a new version.
	//
	// To ensure your data is not overwritten when creating multiple hosted configuration versions in rapid succession, specify the version number of the latest hosted configuration version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_hosted_configuration_version#latest_version_number AppconfigHostedConfigurationVersion#latest_version_number}
	LatestVersionNumber *float64 `field:"optional" json:"latestVersionNumber" yaml:"latestVersionNumber"`
	// A user-defined label for an AWS AppConfig hosted configuration version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appconfig_hosted_configuration_version#version_label AppconfigHostedConfigurationVersion#version_label}
	VersionLabel *string `field:"optional" json:"versionLabel" yaml:"versionLabel"`
}

