// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package amplifydomain


type AmplifyDomainSubDomainSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/amplify_domain#branch_name AmplifyDomain#branch_name}.
	BranchName *string `field:"required" json:"branchName" yaml:"branchName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/amplify_domain#prefix AmplifyDomain#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
}

