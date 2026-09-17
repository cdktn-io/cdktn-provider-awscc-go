// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamappblock


type AppstreamAppBlockTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_app_block#key AppstreamAppBlock#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_app_block#tag_key AppstreamAppBlock#tag_key}.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_app_block#tag_value AppstreamAppBlock#tag_value}.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_app_block#value AppstreamAppBlock#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

