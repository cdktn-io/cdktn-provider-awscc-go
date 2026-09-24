// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3vectorsindex


type S3VectorsIndexMetadataConfiguration struct {
	// Non-filterable metadata keys allow you to enrich vectors with additional context during storage and retrieval.
	//
	// Unlike default metadata keys, these keys cannot be used as query filters. Non-filterable metadata keys can be retrieved but cannot be searched, queried, or filtered. You can access non-filterable metadata keys of your vectors after finding the vectors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/s3vectors_index#non_filterable_metadata_keys S3VectorsIndex#non_filterable_metadata_keys}
	NonFilterableMetadataKeys *[]*string `field:"optional" json:"nonFilterableMetadataKeys" yaml:"nonFilterableMetadataKeys"`
}

