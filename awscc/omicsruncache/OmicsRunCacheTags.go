// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsruncache


type OmicsRunCacheTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/omics_run_cache#key OmicsRunCache#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/omics_run_cache#value OmicsRunCache#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

