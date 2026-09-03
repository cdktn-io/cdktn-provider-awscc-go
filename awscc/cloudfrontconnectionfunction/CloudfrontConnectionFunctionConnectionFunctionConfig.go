// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontconnectionfunction


type CloudfrontConnectionFunctionConnectionFunctionConfig struct {
	// A comment to describe the function. The comment cannot be longer than 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudfront_connection_function#comment CloudfrontConnectionFunction#comment}
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// The function's runtime environment version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudfront_connection_function#runtime CloudfrontConnectionFunction#runtime}
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// The configuration for the key value store associations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/cloudfront_connection_function#key_value_store_associations CloudfrontConnectionFunction#key_value_store_associations}
	KeyValueStoreAssociations interface{} `field:"optional" json:"keyValueStoreAssociations" yaml:"keyValueStoreAssociations"`
}

