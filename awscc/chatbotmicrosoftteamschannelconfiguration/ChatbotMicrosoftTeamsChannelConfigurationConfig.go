// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chatbotmicrosoftteamschannelconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChatbotMicrosoftTeamsChannelConfigurationConfig struct {
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
	// The name of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#configuration_name ChatbotMicrosoftTeamsChannelConfiguration#configuration_name}
	ConfigurationName *string `field:"required" json:"configurationName" yaml:"configurationName"`
	// The ARN of the IAM role that defines the permissions for AWS Chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#iam_role_arn ChatbotMicrosoftTeamsChannelConfiguration#iam_role_arn}
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The id of the Microsoft Teams team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#team_id ChatbotMicrosoftTeamsChannelConfiguration#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// The id of the Microsoft Teams channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#teams_channel_id ChatbotMicrosoftTeamsChannelConfiguration#teams_channel_id}
	TeamsChannelId *string `field:"required" json:"teamsChannelId" yaml:"teamsChannelId"`
	// The id of the Microsoft Teams tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#teams_tenant_id ChatbotMicrosoftTeamsChannelConfiguration#teams_tenant_id}
	TeamsTenantId *string `field:"required" json:"teamsTenantId" yaml:"teamsTenantId"`
	// ARNs of Custom Actions to associate with notifications in the provided chat channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#customization_resource_arns ChatbotMicrosoftTeamsChannelConfiguration#customization_resource_arns}
	CustomizationResourceArns *[]*string `field:"optional" json:"customizationResourceArns" yaml:"customizationResourceArns"`
	// The list of IAM policy ARNs that are applied as channel guardrails.
	//
	// The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#guardrail_policies ChatbotMicrosoftTeamsChannelConfiguration#guardrail_policies}
	GuardrailPolicies *[]*string `field:"optional" json:"guardrailPolicies" yaml:"guardrailPolicies"`
	// Specifies the logging level for this configuration:ERROR,INFO or NONE.
	//
	// This property affects the log entries pushed to Amazon CloudWatch logs
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#logging_level ChatbotMicrosoftTeamsChannelConfiguration#logging_level}
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#sns_topic_arns ChatbotMicrosoftTeamsChannelConfiguration#sns_topic_arns}
	SnsTopicArns *[]*string `field:"optional" json:"snsTopicArns" yaml:"snsTopicArns"`
	// The tags to add to the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#tags ChatbotMicrosoftTeamsChannelConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the Microsoft Teams channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#teams_channel_name ChatbotMicrosoftTeamsChannelConfiguration#teams_channel_name}
	TeamsChannelName *string `field:"optional" json:"teamsChannelName" yaml:"teamsChannelName"`
	// Enables use of a user role requirement in your chat configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/chatbot_microsoft_teams_channel_configuration#user_role_required ChatbotMicrosoftTeamsChannelConfiguration#user_role_required}
	UserRoleRequired interface{} `field:"optional" json:"userRoleRequired" yaml:"userRoleRequired"`
}

