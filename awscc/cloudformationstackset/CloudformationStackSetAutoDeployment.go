// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackset


type CloudformationStackSetAutoDeployment struct {
	// A list of StackSet ARNs that this StackSet depends on for auto-deployment operations.
	//
	// When auto-deployment is triggered, operations will be sequenced to ensure all dependencies complete successfully before this StackSet's operation begins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_stack_set#depends_on CloudformationStackSet#depends_on}
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// If set to true, StackSets automatically deploys additional stack instances to AWS Organizations accounts that are added to a target organization or organizational unit (OU) in the specified Regions.
	//
	// If an account is removed from a target organization or OU, StackSets deletes stack instances from the account in the specified Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_stack_set#enabled CloudformationStackSet#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// If set to true, stack resources are retained when an account is removed from a target organization or OU.
	//
	// If set to false, stack resources are deleted. Specify only if Enabled is set to True.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudformation_stack_set#retain_stacks_on_account_removal CloudformationStackSet#retain_stacks_on_account_removal}
	RetainStacksOnAccountRemoval interface{} `field:"optional" json:"retainStacksOnAccountRemoval" yaml:"retainStacksOnAccountRemoval"`
}

