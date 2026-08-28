// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#marketo CustomerprofilesIntegration#marketo}.
	Marketo *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesMarketo `field:"optional" json:"marketo" yaml:"marketo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#s3 CustomerprofilesIntegration#s3}.
	S3 *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesS3 `field:"optional" json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#salesforce CustomerprofilesIntegration#salesforce}.
	Salesforce *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesSalesforce `field:"optional" json:"salesforce" yaml:"salesforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#service_now CustomerprofilesIntegration#service_now}.
	ServiceNow *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/customerprofiles_integration#zendesk CustomerprofilesIntegration#zendesk}.
	Zendesk *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesZendesk `field:"optional" json:"zendesk" yaml:"zendesk"`
}

