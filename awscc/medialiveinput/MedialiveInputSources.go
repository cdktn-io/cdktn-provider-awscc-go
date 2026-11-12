// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputSources struct {
	// The key used to extract the password from EC2 Parameter store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#password_param MedialiveInput#password_param}
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// This represents the customer's source URL where stream is pulled from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#url MedialiveInput#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The username for the input source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#username MedialiveInput#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

