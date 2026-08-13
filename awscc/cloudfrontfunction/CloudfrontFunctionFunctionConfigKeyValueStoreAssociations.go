// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontfunction


type CloudfrontFunctionFunctionConfigKeyValueStoreAssociations struct {
	// The Amazon Resource Name (ARN) of the key value store association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudfront_function#key_value_store_arn CloudfrontFunction#key_value_store_arn}
	KeyValueStoreArn *string `field:"optional" json:"keyValueStoreArn" yaml:"keyValueStoreArn"`
}

