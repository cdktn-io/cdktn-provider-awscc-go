// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinesswebexperience


type QbusinessWebExperienceCustomizationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_web_experience#custom_css_url QbusinessWebExperience#custom_css_url}.
	CustomCssUrl *string `field:"optional" json:"customCssUrl" yaml:"customCssUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_web_experience#favicon_url QbusinessWebExperience#favicon_url}.
	FaviconUrl *string `field:"optional" json:"faviconUrl" yaml:"faviconUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_web_experience#font_url QbusinessWebExperience#font_url}.
	FontUrl *string `field:"optional" json:"fontUrl" yaml:"fontUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/qbusiness_web_experience#logo_url QbusinessWebExperience#logo_url}.
	LogoUrl *string `field:"optional" json:"logoUrl" yaml:"logoUrl"`
}

