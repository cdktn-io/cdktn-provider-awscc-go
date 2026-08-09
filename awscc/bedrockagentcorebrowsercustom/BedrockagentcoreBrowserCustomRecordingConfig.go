// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorebrowsercustom


type BedrockagentcoreBrowserCustomRecordingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_browser_custom#enabled BedrockagentcoreBrowserCustom#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// S3 Location Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrockagentcore_browser_custom#s3_location BedrockagentcoreBrowserCustom#s3_location}
	S3Location *BedrockagentcoreBrowserCustomRecordingConfigS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

