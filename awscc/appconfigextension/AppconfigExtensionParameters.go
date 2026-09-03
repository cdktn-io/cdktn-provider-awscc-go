// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigextension


type AppconfigExtensionParameters struct {
	// The description of the extension Parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_extension#description AppconfigExtension#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_extension#dynamic AppconfigExtension#dynamic}.
	Dynamic interface{} `field:"optional" json:"dynamic" yaml:"dynamic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appconfig_extension#required AppconfigExtension#required}.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}

