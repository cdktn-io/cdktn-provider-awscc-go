// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsannotationstore


type OmicsAnnotationStoreStoreOptionsTsvStoreOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_annotation_store#annotation_type OmicsAnnotationStore#annotation_type}.
	AnnotationType *string `field:"optional" json:"annotationType" yaml:"annotationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_annotation_store#format_to_header OmicsAnnotationStore#format_to_header}.
	FormatToHeader *map[string]*string `field:"optional" json:"formatToHeader" yaml:"formatToHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_annotation_store#schema OmicsAnnotationStore#schema}.
	Schema interface{} `field:"optional" json:"schema" yaml:"schema"`
}

