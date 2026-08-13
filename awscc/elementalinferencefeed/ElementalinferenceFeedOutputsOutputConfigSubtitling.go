// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elementalinferencefeed


type ElementalinferenceFeedOutputsOutputConfigSubtitling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elementalinference_feed#aspect_ratio ElementalinferenceFeed#aspect_ratio}.
	AspectRatio *ElementalinferenceFeedOutputsOutputConfigSubtitlingAspectRatio `field:"optional" json:"aspectRatio" yaml:"aspectRatio"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elementalinference_feed#dictionary ElementalinferenceFeed#dictionary}.
	Dictionary *string `field:"optional" json:"dictionary" yaml:"dictionary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elementalinference_feed#language ElementalinferenceFeed#language}.
	Language *string `field:"optional" json:"language" yaml:"language"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/elementalinference_feed#profanity_filter ElementalinferenceFeed#profanity_filter}.
	ProfanityFilter *string `field:"optional" json:"profanityFilter" yaml:"profanityFilter"`
}

