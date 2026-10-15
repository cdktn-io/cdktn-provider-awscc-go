// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchintegration


type CloudwatchIntegrationCredential struct {
	// API-key credential material that can be rotated in place on update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/cloudwatch_integration#api_key_credential CloudwatchIntegration#api_key_credential}
	ApiKeyCredential *CloudwatchIntegrationCredentialApiKeyCredential `field:"optional" json:"apiKeyCredential" yaml:"apiKeyCredential"`
}

