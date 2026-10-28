// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityagentIntegrationConfig struct {
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
	// The display name for the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#display_name SecurityagentIntegration#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The integration provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#provider_name SecurityagentIntegration#provider_name}
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The identifier of the AWS KMS key to use for encrypting data associated with the integration.
	//
	// Can be a key ID, key ARN, alias name, or alias ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#kms_key_id SecurityagentIntegration#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The name of an active private connection used to reach a self-hosted provider instance over private networking.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#private_connection_name SecurityagentIntegration#private_connection_name}
	PrivateConnectionName *string `field:"optional" json:"privateConnectionName" yaml:"privateConnectionName"`
	// The provider-specific input required to create the integration. Exactly one of the provider input objects must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#provider_input SecurityagentIntegration#provider_input}
	ProviderInput *SecurityagentIntegrationProviderInput `field:"optional" json:"providerInput" yaml:"providerInput"`
	// The tags to associate with the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/securityagent_integration#tags SecurityagentIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

