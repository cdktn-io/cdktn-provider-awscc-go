// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontkeyvaluestore


type CloudfrontKeyValueStoreImportSource struct {
	// The Amazon Resource Name (ARN) of the import source for the key value store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudfront_key_value_store#source_arn CloudfrontKeyValueStore#source_arn}
	SourceArn *string `field:"optional" json:"sourceArn" yaml:"sourceArn"`
	// The source type of the import source for the key value store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudfront_key_value_store#source_type CloudfrontKeyValueStore#source_type}
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
}

