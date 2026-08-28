// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package translateparalleldata


type TranslateParallelDataEncryptionKey struct {
	// The Amazon Resource Name (ARN) of the encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/translate_parallel_data#id TranslateParallelData#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The type of encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/translate_parallel_data#type TranslateParallelData#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

