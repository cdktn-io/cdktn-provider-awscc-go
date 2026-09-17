// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecodeinterpretercustom


type BedrockagentcoreCodeInterpreterCustomCertificates struct {
	// Certificate location in Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_code_interpreter_custom#certificate_location BedrockagentcoreCodeInterpreterCustom#certificate_location}
	CertificateLocation *BedrockagentcoreCodeInterpreterCustomCertificatesCertificateLocation `field:"optional" json:"certificateLocation" yaml:"certificateLocation"`
}

