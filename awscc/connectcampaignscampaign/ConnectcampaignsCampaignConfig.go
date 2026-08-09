// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignscampaign

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectcampaignsCampaignConfig struct {
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
	// Amazon Connect Instance Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaigns_campaign#connect_instance_arn ConnectcampaignsCampaign#connect_instance_arn}
	ConnectInstanceArn *string `field:"required" json:"connectInstanceArn" yaml:"connectInstanceArn"`
	// The possible types of dialer config parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaigns_campaign#dialer_config ConnectcampaignsCampaign#dialer_config}
	DialerConfig *ConnectcampaignsCampaignDialerConfig `field:"required" json:"dialerConfig" yaml:"dialerConfig"`
	// Amazon Connect Campaign Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaigns_campaign#name ConnectcampaignsCampaign#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The configuration used for outbound calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaigns_campaign#outbound_call_config ConnectcampaignsCampaign#outbound_call_config}
	OutboundCallConfig *ConnectcampaignsCampaignOutboundCallConfig `field:"required" json:"outboundCallConfig" yaml:"outboundCallConfig"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connectcampaigns_campaign#tags ConnectcampaignsCampaign#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

