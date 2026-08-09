// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package finspaceenvironment


type FinspaceEnvironmentFederationParameters struct {
	// SAML metadata URL to link with the Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/finspace_environment#application_call_back_url FinspaceEnvironment#application_call_back_url}
	ApplicationCallBackUrl *string `field:"optional" json:"applicationCallBackUrl" yaml:"applicationCallBackUrl"`
	// Attribute map for SAML configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/finspace_environment#attribute_map FinspaceEnvironment#attribute_map}
	AttributeMap interface{} `field:"optional" json:"attributeMap" yaml:"attributeMap"`
	// Federation provider name to link with the Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/finspace_environment#federation_provider_name FinspaceEnvironment#federation_provider_name}
	FederationProviderName *string `field:"optional" json:"federationProviderName" yaml:"federationProviderName"`
	// SAML metadata URL to link with the Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/finspace_environment#federation_urn FinspaceEnvironment#federation_urn}
	FederationUrn *string `field:"optional" json:"federationUrn" yaml:"federationUrn"`
	// SAML metadata document to link the federation provider to the Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/finspace_environment#saml_metadata_document FinspaceEnvironment#saml_metadata_document}
	SamlMetadataDocument *string `field:"optional" json:"samlMetadataDocument" yaml:"samlMetadataDocument"`
	// SAML metadata URL to link with the Environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/finspace_environment#saml_metadata_url FinspaceEnvironment#saml_metadata_url}
	SamlMetadataUrl *string `field:"optional" json:"samlMetadataUrl" yaml:"samlMetadataUrl"`
}

