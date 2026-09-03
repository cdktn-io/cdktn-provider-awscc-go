// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerproject


type SagemakerProjectTemplateProviderDetails struct {
	// CloudFormation template provider details for a SageMaker project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_project#cfn_template_provider_detail SagemakerProject#cfn_template_provider_detail}
	CfnTemplateProviderDetail *SagemakerProjectTemplateProviderDetailsCfnTemplateProviderDetail `field:"optional" json:"cfnTemplateProviderDetail" yaml:"cfnTemplateProviderDetail"`
}

