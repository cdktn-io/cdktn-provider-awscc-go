// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationAws struct {
	// AWS Account Id corresponding to provided resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#account_id DevopsagentAssociation#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Account Type 'monitor' for DevOpsAgent monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#account_type DevopsagentAssociation#account_type}
	AccountType *string `field:"optional" json:"accountType" yaml:"accountType"`
	// Role ARN to be assumed by DevOpsAgent to operate on behalf of customer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#assumable_role_arn DevopsagentAssociation#assumable_role_arn}
	AssumableRoleArn *string `field:"optional" json:"assumableRoleArn" yaml:"assumableRoleArn"`
	// List of AWS resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#resources DevopsagentAssociation#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// List of AWS tags as key-value pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#tags DevopsagentAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

