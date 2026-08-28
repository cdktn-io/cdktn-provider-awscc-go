// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupDeploymentRevisionGitHubLocation struct {
	// The SHA1 commit ID of the GitHub commit that represents the bundled artifacts for the application revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codedeploy_deployment_group#commit_id CodedeployDeploymentGroup#commit_id}
	CommitId *string `field:"optional" json:"commitId" yaml:"commitId"`
	// The GitHub account and repository pair that stores the application revision to be deployed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/codedeploy_deployment_group#repository CodedeployDeploymentGroup#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

