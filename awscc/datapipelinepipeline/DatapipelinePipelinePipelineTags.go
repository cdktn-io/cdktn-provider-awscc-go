// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DatapipelinePipelinePipelineTags struct {
	// The key name of a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datapipeline_pipeline#key DatapipelinePipeline#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value to associate with the key name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datapipeline_pipeline#value DatapipelinePipeline#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

