// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorebrowsercustom


type BedrockagentcoreBrowserCustomEnterprisePolicies struct {
	// The S3 location of the enterprise policy file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_browser_custom#location BedrockagentcoreBrowserCustom#location}
	Location *BedrockagentcoreBrowserCustomEnterprisePoliciesLocation `field:"optional" json:"location" yaml:"location"`
	// The type of browser enterprise policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bedrockagentcore_browser_custom#type BedrockagentcoreBrowserCustom#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

