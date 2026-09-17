// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerproject


type SagemakerProjectTemplateProviderDetailsCfnTemplateProviderDetail struct {
	// A list of parameters used in the CloudFormation template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_project#parameters SagemakerProject#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The Amazon Resource Name (ARN) of the IAM role used by the template provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_project#role_arn SagemakerProject#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The name of the template used for the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_project#template_name SagemakerProject#template_name}
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// The URL of the CloudFormation template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_project#template_url SagemakerProject#template_url}
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
}

