// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupDeploymentRevision struct {
	// Specifies the location of an application revision that is stored in GitHub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#git_hub_location CodedeployDeploymentGroup#git_hub_location}
	GitHubLocation *CodedeployDeploymentGroupDeploymentRevisionGitHubLocation `field:"optional" json:"gitHubLocation" yaml:"gitHubLocation"`
	// The type of application revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#revision_type CodedeployDeploymentGroup#revision_type}
	RevisionType *string `field:"optional" json:"revisionType" yaml:"revisionType"`
	// Information about the location of application artifacts stored in Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#s3_location CodedeployDeploymentGroup#s3_location}
	S3Location *CodedeployDeploymentGroupDeploymentRevisionS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

