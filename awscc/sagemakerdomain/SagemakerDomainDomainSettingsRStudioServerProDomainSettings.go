// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainDomainSettingsRStudioServerProDomainSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_domain#default_resource_spec SagemakerDomain#default_resource_spec}.
	DefaultResourceSpec *SagemakerDomainDomainSettingsRStudioServerProDomainSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// The ARN of the execution role for the RStudioServerPro Domain-level app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_domain#domain_execution_role_arn SagemakerDomain#domain_execution_role_arn}
	DomainExecutionRoleArn *string `field:"optional" json:"domainExecutionRoleArn" yaml:"domainExecutionRoleArn"`
	// A URL pointing to an RStudio Connect server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_domain#r_studio_connect_url SagemakerDomain#r_studio_connect_url}
	RStudioConnectUrl *string `field:"optional" json:"rStudioConnectUrl" yaml:"rStudioConnectUrl"`
	// A URL pointing to an RStudio Package Manager server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_domain#r_studio_package_manager_url SagemakerDomain#r_studio_package_manager_url}
	RStudioPackageManagerUrl *string `field:"optional" json:"rStudioPackageManagerUrl" yaml:"rStudioPackageManagerUrl"`
}

