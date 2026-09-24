// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimechannel


type ChimeChannelElasticChannelConfiguration struct {
	// The maximum number of SubChannels allowed in the elastic channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#maximum_sub_channels ChimeChannel#maximum_sub_channels}
	MaximumSubChannels *float64 `field:"optional" json:"maximumSubChannels" yaml:"maximumSubChannels"`
	// The minimum allowed percentage of TargetMembershipsPerSubChannel users, used to balance members across SubChannels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#minimum_membership_percentage ChimeChannel#minimum_membership_percentage}
	MinimumMembershipPercentage *float64 `field:"optional" json:"minimumMembershipPercentage" yaml:"minimumMembershipPercentage"`
	// The maximum number of members allowed in a SubChannel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#target_memberships_per_sub_channel ChimeChannel#target_memberships_per_sub_channel}
	TargetMembershipsPerSubChannel *float64 `field:"optional" json:"targetMembershipsPerSubChannel" yaml:"targetMembershipsPerSubChannel"`
}

