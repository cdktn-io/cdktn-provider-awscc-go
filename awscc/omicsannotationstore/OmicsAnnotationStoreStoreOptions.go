// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsannotationstore


type OmicsAnnotationStoreStoreOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_annotation_store#tsv_store_options OmicsAnnotationStore#tsv_store_options}.
	TsvStoreOptions *OmicsAnnotationStoreStoreOptionsTsvStoreOptions `field:"optional" json:"tsvStoreOptions" yaml:"tsvStoreOptions"`
}

