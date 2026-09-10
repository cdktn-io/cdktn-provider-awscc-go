// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsAmazonQProperties struct {
	// The authentication mode of the connection's AmazonQ properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#auth_mode DatazoneConnection#auth_mode}
	AuthMode *string `field:"optional" json:"authMode" yaml:"authMode"`
	// Specifies whether Amazon Q is enabled for the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#is_enabled DatazoneConnection#is_enabled}
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#profile_arn DatazoneConnection#profile_arn}.
	ProfileArn *string `field:"optional" json:"profileArn" yaml:"profileArn"`
}

