// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationAws struct {
	// AWS Account Id corresponding to provided resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#account_id DevopsagentAssociation#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Account Type 'monitor' for DevOpsAgent monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#account_type DevopsagentAssociation#account_type}
	AccountType *string `field:"optional" json:"accountType" yaml:"accountType"`
	// The IAM role that AWS DevOps Agent assumes to perform elevated directed actions in this AWS account.
	//
	// Elevated actions also require ElevatedActionsEnabled to be set in the Preferences of the AgentSpace. When not provided, only non-elevated directed actions are available. Removing this property from the template clears the role on update. Setting this role is subject to the same iam:PassRole requirement as AssumableRoleArn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#agent_elevated_role_arn DevopsagentAssociation#agent_elevated_role_arn}
	AgentElevatedRoleArn *string `field:"optional" json:"agentElevatedRoleArn" yaml:"agentElevatedRoleArn"`
	// Role ARN to be assumed by DevOpsAgent to operate on behalf of customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#assumable_role_arn DevopsagentAssociation#assumable_role_arn}
	AssumableRoleArn *string `field:"optional" json:"assumableRoleArn" yaml:"assumableRoleArn"`
	// List of AWS resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#resources DevopsagentAssociation#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// List of AWS tags as key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/devopsagent_association#tags DevopsagentAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

