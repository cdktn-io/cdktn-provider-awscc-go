// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupDeploymentRevisionS3Location struct {
	// The name of the Amazon S3 bucket where the application revision is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#bucket CodedeployDeploymentGroup#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The file type of the application revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#bundle_type CodedeployDeploymentGroup#bundle_type}
	BundleType *string `field:"optional" json:"bundleType" yaml:"bundleType"`
	// The ETag of the Amazon S3 object that represents the bundled artifacts for the application revision.
	//
	// If the ETag is not specified as an input parameter, ETag validation of the object is skipped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#e_tag CodedeployDeploymentGroup#e_tag}
	ETag *string `field:"optional" json:"eTag" yaml:"eTag"`
	// The name of the Amazon S3 object that represents the bundled artifacts for the application revision.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#key CodedeployDeploymentGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A specific version of the Amazon S3 object that represents the bundled artifacts for the application revision.
	//
	// If the version is not specified, the system uses the most recent version by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_group#version CodedeployDeploymentGroup#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

