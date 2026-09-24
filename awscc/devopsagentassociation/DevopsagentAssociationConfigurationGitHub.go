// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationGitHub struct {
	// Repository owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#owner DevopsagentAssociation#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Type of repository owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#owner_type DevopsagentAssociation#owner_type}
	OwnerType *string `field:"optional" json:"ownerType" yaml:"ownerType"`
	// Associated Github repo ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#repo_id DevopsagentAssociation#repo_id}
	RepoId *string `field:"optional" json:"repoId" yaml:"repoId"`
	// Associated Github repo name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#repo_name DevopsagentAssociation#repo_name}
	RepoName *string `field:"optional" json:"repoName" yaml:"repoName"`
}

