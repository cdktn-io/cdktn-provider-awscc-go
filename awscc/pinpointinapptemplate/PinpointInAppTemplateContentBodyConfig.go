// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointinapptemplate


type PinpointInAppTemplateContentBodyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pinpoint_in_app_template#alignment PinpointInAppTemplate#alignment}.
	Alignment *string `field:"optional" json:"alignment" yaml:"alignment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pinpoint_in_app_template#body PinpointInAppTemplate#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pinpoint_in_app_template#text_color PinpointInAppTemplate#text_color}.
	TextColor *string `field:"optional" json:"textColor" yaml:"textColor"`
}

