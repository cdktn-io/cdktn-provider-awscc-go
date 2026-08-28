// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorebrowsercustom


type BedrockagentcoreBrowserCustomCertificates struct {
	// Certificate location in Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_browser_custom#certificate_location BedrockagentcoreBrowserCustom#certificate_location}
	CertificateLocation *BedrockagentcoreBrowserCustomCertificatesCertificateLocation `field:"optional" json:"certificateLocation" yaml:"certificateLocation"`
}

