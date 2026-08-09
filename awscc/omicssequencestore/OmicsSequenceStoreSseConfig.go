// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicssequencestore


type OmicsSequenceStoreSseConfig struct {
	// An encryption key ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/omics_sequence_store#key_arn OmicsSequenceStore#key_arn}
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/omics_sequence_store#type OmicsSequenceStore#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

