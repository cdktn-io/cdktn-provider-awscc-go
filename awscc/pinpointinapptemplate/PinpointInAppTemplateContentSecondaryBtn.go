// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pinpointinapptemplate


type PinpointInAppTemplateContentSecondaryBtn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pinpoint_in_app_template#android PinpointInAppTemplate#android}.
	Android *PinpointInAppTemplateContentSecondaryBtnAndroid `field:"optional" json:"android" yaml:"android"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pinpoint_in_app_template#default_config PinpointInAppTemplate#default_config}.
	DefaultConfig *PinpointInAppTemplateContentSecondaryBtnDefaultConfig `field:"optional" json:"defaultConfig" yaml:"defaultConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pinpoint_in_app_template#ios PinpointInAppTemplate#ios}.
	Ios *PinpointInAppTemplateContentSecondaryBtnIos `field:"optional" json:"ios" yaml:"ios"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pinpoint_in_app_template#web PinpointInAppTemplate#web}.
	Web *PinpointInAppTemplateContentSecondaryBtnWeb `field:"optional" json:"web" yaml:"web"`
}

