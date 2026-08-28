// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package locationapikey


type LocationApiKeyRestrictions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#allow_actions LocationApiKey#allow_actions}.
	AllowActions *[]*string `field:"required" json:"allowActions" yaml:"allowActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#allow_resources LocationApiKey#allow_resources}.
	AllowResources *[]*string `field:"required" json:"allowResources" yaml:"allowResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#allow_android_apps LocationApiKey#allow_android_apps}.
	AllowAndroidApps interface{} `field:"optional" json:"allowAndroidApps" yaml:"allowAndroidApps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#allow_apple_apps LocationApiKey#allow_apple_apps}.
	AllowAppleApps interface{} `field:"optional" json:"allowAppleApps" yaml:"allowAppleApps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/location_api_key#allow_referers LocationApiKey#allow_referers}.
	AllowReferers *[]*string `field:"optional" json:"allowReferers" yaml:"allowReferers"`
}

