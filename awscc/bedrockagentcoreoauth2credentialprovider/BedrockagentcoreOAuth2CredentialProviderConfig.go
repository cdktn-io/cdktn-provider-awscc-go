// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreOAuth2CredentialProviderConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The vendor of the OAuth2 credential provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#credential_provider_vendor BedrockagentcoreOAuth2CredentialProvider#credential_provider_vendor}
	CredentialProviderVendor *string `field:"required" json:"credentialProviderVendor" yaml:"credentialProviderVendor"`
	// The name of the OAuth2 credential provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#name BedrockagentcoreOAuth2CredentialProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The configuration settings for the OAuth2 provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#oauth_2_provider_config_input BedrockagentcoreOAuth2CredentialProvider#oauth_2_provider_config_input}
	Oauth2ProviderConfigInput *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInput `field:"optional" json:"oauth2ProviderConfigInput" yaml:"oauth2ProviderConfigInput"`
	// Tags to assign to the OAuth2 credential provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#tags BedrockagentcoreOAuth2CredentialProvider#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

