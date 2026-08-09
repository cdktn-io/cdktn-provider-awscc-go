// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elementalinferencefeed


type ElementalinferenceFeedOutputsOutputConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elementalinference_feed#clipping ElementalinferenceFeed#clipping}.
	Clipping *ElementalinferenceFeedOutputsOutputConfigClipping `field:"optional" json:"clipping" yaml:"clipping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elementalinference_feed#cropping ElementalinferenceFeed#cropping}.
	Cropping *string `field:"optional" json:"cropping" yaml:"cropping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elementalinference_feed#subtitling ElementalinferenceFeed#subtitling}.
	Subtitling *ElementalinferenceFeedOutputsOutputConfigSubtitling `field:"optional" json:"subtitling" yaml:"subtitling"`
}

