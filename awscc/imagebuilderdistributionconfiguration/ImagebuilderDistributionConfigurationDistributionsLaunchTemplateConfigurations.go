// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionsLaunchTemplateConfigurations struct {
	// The account ID that this configuration applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_distribution_configuration#account_id ImagebuilderDistributionConfiguration#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Identifies the EC2 launch template to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_distribution_configuration#launch_template_id ImagebuilderDistributionConfiguration#launch_template_id}
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Set the specified EC2 launch template as the default launch template for the specified account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_distribution_configuration#set_default_version ImagebuilderDistributionConfiguration#set_default_version}
	SetDefaultVersion interface{} `field:"optional" json:"setDefaultVersion" yaml:"setDefaultVersion"`
}

