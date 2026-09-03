// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsannotationstore


type OmicsAnnotationStoreSseConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_annotation_store#key_arn OmicsAnnotationStore#key_arn}.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_annotation_store#type OmicsAnnotationStore#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

