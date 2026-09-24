// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedagentprofile


type WellarchitectedAgentProfileAggregationConfiguration struct {
	// The ARN of the IAM role used to access resources in this account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#access_role_arn WellarchitectedAgentProfile#access_role_arn}
	AccessRoleArn *string `field:"required" json:"accessRoleArn" yaml:"accessRoleArn"`
	// The target AWS account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#account_id WellarchitectedAgentProfile#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The target regions in the account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_agent_profile#regions WellarchitectedAgentProfile#regions}
	Regions *[]*string `field:"required" json:"regions" yaml:"regions"`
}

