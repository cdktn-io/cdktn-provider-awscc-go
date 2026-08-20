// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightagent


type QuicksightAgentCustomPromptInputExistingPrompt struct {
	// The identifier of the model profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_agent#model_profile_id QuicksightAgent#model_profile_id}
	ModelProfileId *string `field:"optional" json:"modelProfileId" yaml:"modelProfileId"`
	// The QBS AWS account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_agent#qbs_aws_account_id QuicksightAgent#qbs_aws_account_id}
	QbsAwsAccountId *string `field:"optional" json:"qbsAwsAccountId" yaml:"qbsAwsAccountId"`
	// The subscription identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_agent#subscription_id QuicksightAgent#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
}

