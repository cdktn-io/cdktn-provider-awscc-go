// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDomainSettingsUnifiedStudioSettings struct {
	// The ID of the AWS account that has the Amazon SageMaker Unified Studio domain.
	//
	// The default value, if you don't specify an ID, is the ID of the account that has the Amazon SageMaker AI domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#domain_account_id SagemakerDomain#domain_account_id}
	DomainAccountId *string `field:"optional" json:"domainAccountId" yaml:"domainAccountId"`
	// The ID of the Amazon SageMaker Unified Studio domain associated with this domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#domain_id SagemakerDomain#domain_id}
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// The AWS Region where the domain is located in Amazon SageMaker Unified Studio.
	//
	// The default value, if you don't specify a Region, is the Region where the Amazon SageMaker AI domain is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#domain_region SagemakerDomain#domain_region}
	DomainRegion *string `field:"optional" json:"domainRegion" yaml:"domainRegion"`
	// The ID of the environment that Amazon SageMaker Unified Studio associates with the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#environment_id SagemakerDomain#environment_id}
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// The ID of the Amazon SageMaker Unified Studio project that corresponds to the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#project_id SagemakerDomain#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// The location where Amazon S3 stores temporary execution data and other artifacts for the project that corresponds to the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#project_s3_path SagemakerDomain#project_s3_path}
	ProjectS3Path *string `field:"optional" json:"projectS3Path" yaml:"projectS3Path"`
	// Sets whether you can access the domain in Amazon SageMaker Studio:.
	//
	// ENABLED
	// You can access the domain in Amazon SageMaker Studio. If you migrate the domain to Amazon SageMaker Unified Studio, you can access it in both studio interfaces.
	// DISABLED
	// You can't access the domain in Amazon SageMaker Studio. If you migrate the domain to Amazon SageMaker Unified Studio, you can access it only in that studio interface.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_domain#studio_web_portal_access SagemakerDomain#studio_web_portal_access}
	StudioWebPortalAccess *string `field:"optional" json:"studioWebPortalAccess" yaml:"studioWebPortalAccess"`
}

