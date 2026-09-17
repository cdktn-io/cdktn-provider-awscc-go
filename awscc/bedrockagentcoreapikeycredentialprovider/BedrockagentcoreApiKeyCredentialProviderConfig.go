// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreapikeycredentialprovider

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcoreApiKeyCredentialProviderConfig struct {
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
	// The name of the API key credential provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_api_key_credential_provider#name BedrockagentcoreApiKeyCredentialProvider#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The API key to use for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_api_key_credential_provider#api_key BedrockagentcoreApiKeyCredentialProvider#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// Configuration for a customer-provided secret containing the API key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_api_key_credential_provider#api_key_secret_config BedrockagentcoreApiKeyCredentialProvider#api_key_secret_config}
	ApiKeySecretConfig *BedrockagentcoreApiKeyCredentialProviderApiKeySecretConfig `field:"optional" json:"apiKeySecretConfig" yaml:"apiKeySecretConfig"`
	// The source of the API key secret. Use MANAGED for service-managed secrets or EXTERNAL for customer-provided secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_api_key_credential_provider#api_key_secret_source BedrockagentcoreApiKeyCredentialProvider#api_key_secret_source}
	ApiKeySecretSource *string `field:"optional" json:"apiKeySecretSource" yaml:"apiKeySecretSource"`
	// Tags to assign to the API key credential provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_api_key_credential_provider#tags BedrockagentcoreApiKeyCredentialProvider#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

