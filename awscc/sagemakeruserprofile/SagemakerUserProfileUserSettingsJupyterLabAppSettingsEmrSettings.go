// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsJupyterLabAppSettingsEmrSettings struct {
	// An array of Amazon Resource Names (ARNs) of the IAM roles that the execution role of SageMaker can assume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_user_profile#assumable_role_arns SagemakerUserProfile#assumable_role_arns}
	AssumableRoleArns *[]*string `field:"optional" json:"assumableRoleArns" yaml:"assumableRoleArns"`
	// An array of ARNs of IAM roles used by EMR cluster instances or job execution environments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_user_profile#execution_role_arns SagemakerUserProfile#execution_role_arns}
	ExecutionRoleArns *[]*string `field:"optional" json:"executionRoleArns" yaml:"executionRoleArns"`
}

