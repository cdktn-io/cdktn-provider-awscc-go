// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontkeygroup


type CloudfrontKeyGroupKeyGroupConfig struct {
	// A list of the identifiers of the public keys in the key group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_key_group#items CloudfrontKeyGroup#items}
	Items *[]*string `field:"required" json:"items" yaml:"items"`
	// A name to identify the key group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_key_group#name CloudfrontKeyGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment to describe the key group. The comment cannot be longer than 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cloudfront_key_group#comment CloudfrontKeyGroup#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

