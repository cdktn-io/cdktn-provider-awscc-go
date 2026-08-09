// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appconfigextension


type AppconfigExtensionActions struct {
	// The description of the extension Action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appconfig_extension#description AppconfigExtension#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the extension action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appconfig_extension#name AppconfigExtension#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ARN of the role for invoking the extension action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appconfig_extension#role_arn AppconfigExtension#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The URI of the extension action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/appconfig_extension#uri AppconfigExtension#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

