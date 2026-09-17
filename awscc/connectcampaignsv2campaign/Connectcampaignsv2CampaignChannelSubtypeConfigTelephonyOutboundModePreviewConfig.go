// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePreviewConfig struct {
	// Actions that can be performed by agent during preview phase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaignsv2_campaign#agent_actions Connectcampaignsv2Campaign#agent_actions}
	AgentActions *[]*string `field:"optional" json:"agentActions" yaml:"agentActions"`
	// The bandwidth allocation of a queue resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaignsv2_campaign#bandwidth_allocation Connectcampaignsv2Campaign#bandwidth_allocation}
	BandwidthAllocation *float64 `field:"optional" json:"bandwidthAllocation" yaml:"bandwidthAllocation"`
	// Timeout Config for preview contacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaignsv2_campaign#timeout_config Connectcampaignsv2Campaign#timeout_config}
	TimeoutConfig *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyOutboundModePreviewConfigTimeoutConfig `field:"optional" json:"timeoutConfig" yaml:"timeoutConfig"`
}

