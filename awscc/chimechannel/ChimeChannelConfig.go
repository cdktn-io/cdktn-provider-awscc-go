// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimechannel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeChannelConfig struct {
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
	// The ARN of the AppInstance that contains the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#app_instance_arn ChimeChannel#app_instance_arn}
	AppInstanceArn *string `field:"required" json:"appInstanceArn" yaml:"appInstanceArn"`
	// The ARN of the AppInstanceUser or AppInstanceBot that performs every operation on this channel.
	//
	// Whichever of the two creates a channel automatically becomes one of its moderators, so the same ARN can subsequently read, update and delete the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#chime_bearer ChimeChannel#chime_bearer}
	ChimeBearer *string `field:"required" json:"chimeBearer" yaml:"chimeBearer"`
	// The name of the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#name ChimeChannel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the channel. When omitted, the service generates a UUID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#channel_id ChimeChannel#channel_id}
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// The attributes required to configure and create an elastic channel.
	//
	// An elastic channel must use RESTRICTED mode, cannot be created with MemberArns, and is available only in some regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#elastic_channel_configuration ChimeChannel#elastic_channel_configuration}
	ElasticChannelConfiguration *ChimeChannelElasticChannelConfiguration `field:"optional" json:"elasticChannelConfiguration" yaml:"elasticChannelConfiguration"`
	// Settings that control the interval after which the channel is automatically deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#expiration_settings ChimeChannel#expiration_settings}
	ExpirationSettings *ChimeChannelExpirationSettings `field:"optional" json:"expirationSettings" yaml:"expirationSettings"`
	// The ARNs of the AppInstanceUsers to add to the channel as members when it is created.
	//
	// Cannot be combined with ElasticChannelConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#member_arns ChimeChannel#member_arns}
	MemberArns *[]*string `field:"optional" json:"memberArns" yaml:"memberArns"`
	// The metadata of the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#metadata ChimeChannel#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// The channel mode.
	//
	// In an UNRESTRICTED channel, members can add themselves and other members; in a RESTRICTED channel, only administrators and moderators can add members. An elastic channel must be RESTRICTED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#mode ChimeChannel#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The ARNs of the AppInstanceUsers to add to the channel as moderators when it is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#moderator_arns ChimeChannel#moderator_arns}
	ModeratorArns *[]*string `field:"optional" json:"moderatorArns" yaml:"moderatorArns"`
	// The channel's privacy level.
	//
	// A PUBLIC channel is discoverable by anyone in the AppInstance; a PRIVATE channel is not. Privacy cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#privacy ChimeChannel#privacy}
	Privacy *string `field:"optional" json:"privacy" yaml:"privacy"`
	// The tags for the channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/chime_channel#tags ChimeChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

