// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package finspaceenvironment


type FinspaceEnvironmentSuperuserParameters struct {
	// Email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#email_address FinspaceEnvironment#email_address}
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// First name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#first_name FinspaceEnvironment#first_name}
	FirstName *string `field:"optional" json:"firstName" yaml:"firstName"`
	// Last name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/finspace_environment#last_name FinspaceEnvironment#last_name}
	LastName *string `field:"optional" json:"lastName" yaml:"lastName"`
}

