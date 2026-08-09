// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codepipelinewebhook


type CodepipelineWebhookAuthenticationConfiguration struct {
	// The property used to configure acceptance of webhooks in an IP address range.
	//
	// For IP, only the AllowedIPRange property must be set. This property must be set to a valid CIDR range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_webhook#allowed_ip_range CodepipelineWebhook#allowed_ip_range}
	AllowedIpRange *string `field:"optional" json:"allowedIpRange" yaml:"allowedIpRange"`
	// The property used to configure GitHub authentication. For GITHUB_HMAC, only the SecretToken property must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/codepipeline_webhook#secret_token CodepipelineWebhook#secret_token}
	SecretToken *string `field:"optional" json:"secretToken" yaml:"secretToken"`
}

