// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackset


type CloudformationStackSetStackInstancesGroupDeploymentTargets struct {
	// The filter type you want to apply on organizational units and accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_stack_set#account_filter_type CloudformationStackSet#account_filter_type}
	AccountFilterType *string `field:"optional" json:"accountFilterType" yaml:"accountFilterType"`
	// AWS accounts that you want to create stack instances in the specified Region(s) for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_stack_set#accounts CloudformationStackSet#accounts}
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// Returns the value of the AccountsUrl property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_stack_set#accounts_url CloudformationStackSet#accounts_url}
	AccountsUrl *string `field:"optional" json:"accountsUrl" yaml:"accountsUrl"`
	// The organization root ID or organizational unit (OU) IDs to which StackSets deploys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_stack_set#organizational_unit_ids CloudformationStackSet#organizational_unit_ids}
	OrganizationalUnitIds *[]*string `field:"optional" json:"organizationalUnitIds" yaml:"organizationalUnitIds"`
}

