// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizecampaign

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PersonalizeCampaignConfig struct {
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
	// The name of the campaign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/personalize_campaign#name PersonalizeCampaign#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ARN of the solution version to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/personalize_campaign#solution_version_arn PersonalizeCampaign#solution_version_arn}
	SolutionVersionArn *string `field:"required" json:"solutionVersionArn" yaml:"solutionVersionArn"`
	// The configuration details of a campaign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/personalize_campaign#campaign_config PersonalizeCampaign#campaign_config}
	CampaignConfig *PersonalizeCampaignCampaignConfig `field:"optional" json:"campaignConfig" yaml:"campaignConfig"`
	// Specifies the requested minimum provisioned transactions per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/personalize_campaign#min_provisioned_tps PersonalizeCampaign#min_provisioned_tps}
	MinProvisionedTps *float64 `field:"optional" json:"minProvisionedTps" yaml:"minProvisionedTps"`
	// Tags to associate with the campaign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/personalize_campaign#tags PersonalizeCampaign#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

