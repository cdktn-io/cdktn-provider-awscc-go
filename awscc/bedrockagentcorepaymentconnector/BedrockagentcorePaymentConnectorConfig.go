// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorepaymentconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockagentcorePaymentConnectorConfig struct {
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
	// The name of the payment connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_payment_connector#connector_name BedrockagentcorePaymentConnector#connector_name}
	ConnectorName *string `field:"required" json:"connectorName" yaml:"connectorName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_payment_connector#connector_type BedrockagentcorePaymentConnector#connector_type}.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// The identifier of the parent payment manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_payment_connector#payment_manager_id BedrockagentcorePaymentConnector#payment_manager_id}
	PaymentManagerId *string `field:"required" json:"paymentManagerId" yaml:"paymentManagerId"`
	// The credential provider configurations for the connector.
	//
	// Required when ProvisionMode is MANUAL or not specified. Empty for QUICK_CREATE until provisioning completes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_payment_connector#credential_provider_configurations BedrockagentcorePaymentConnector#credential_provider_configurations}
	CredentialProviderConfigurations interface{} `field:"optional" json:"credentialProviderConfigurations" yaml:"credentialProviderConfigurations"`
	// A description of the payment connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_payment_connector#description BedrockagentcorePaymentConnector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The provision mode for creating the connector. MANUAL requires CredentialProviderConfigurations; QUICK_CREATE orchestrates OAuth consent and credential provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_payment_connector#provision_mode BedrockagentcorePaymentConnector#provision_mode}
	ProvisionMode *string `field:"optional" json:"provisionMode" yaml:"provisionMode"`
}

