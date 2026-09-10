// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointinapptemplate


type PinpointInAppTemplateContent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pinpoint_in_app_template#background_color PinpointInAppTemplate#background_color}.
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pinpoint_in_app_template#body_config PinpointInAppTemplate#body_config}.
	BodyConfig *PinpointInAppTemplateContentBodyConfig `field:"optional" json:"bodyConfig" yaml:"bodyConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pinpoint_in_app_template#header_config PinpointInAppTemplate#header_config}.
	HeaderConfig *PinpointInAppTemplateContentHeaderConfig `field:"optional" json:"headerConfig" yaml:"headerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pinpoint_in_app_template#image_url PinpointInAppTemplate#image_url}.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pinpoint_in_app_template#primary_btn PinpointInAppTemplate#primary_btn}.
	PrimaryBtn *PinpointInAppTemplateContentPrimaryBtn `field:"optional" json:"primaryBtn" yaml:"primaryBtn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pinpoint_in_app_template#secondary_btn PinpointInAppTemplate#secondary_btn}.
	SecondaryBtn *PinpointInAppTemplateContentSecondaryBtn `field:"optional" json:"secondaryBtn" yaml:"secondaryBtn"`
}

