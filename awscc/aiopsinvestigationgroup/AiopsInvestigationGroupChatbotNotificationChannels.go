// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aiopsinvestigationgroup


type AiopsInvestigationGroupChatbotNotificationChannels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/aiops_investigation_group#chat_configuration_arns AiopsInvestigationGroup#chat_configuration_arns}.
	ChatConfigurationArns *[]*string `field:"optional" json:"chatConfigurationArns" yaml:"chatConfigurationArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/aiops_investigation_group#sns_topic_arn AiopsInvestigationGroup#sns_topic_arn}.
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

