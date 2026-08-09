// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorebrowsercustom


type BedrockagentcoreBrowserCustomCertificatesCertificateLocation struct {
	// Secrets Manager secret ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_browser_custom#secret_arn BedrockagentcoreBrowserCustom#secret_arn}
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

