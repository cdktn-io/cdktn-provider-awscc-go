// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionTasksConnectorOperator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#marketo CustomerprofilesIntegration#marketo}.
	Marketo *string `field:"optional" json:"marketo" yaml:"marketo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#s3 CustomerprofilesIntegration#s3}.
	S3 *string `field:"optional" json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#salesforce CustomerprofilesIntegration#salesforce}.
	Salesforce *string `field:"optional" json:"salesforce" yaml:"salesforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#service_now CustomerprofilesIntegration#service_now}.
	ServiceNow *string `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#zendesk CustomerprofilesIntegration#zendesk}.
	Zendesk *string `field:"optional" json:"zendesk" yaml:"zendesk"`
}

