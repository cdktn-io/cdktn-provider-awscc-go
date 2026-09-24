// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncchannelnamespace


type AppsyncChannelNamespaceTags struct {
	// A string used to identify this tag. You can specify a maximum of 128 characters for a tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_channel_namespace#key AppsyncChannelNamespace#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A string containing the value for this tag.
	//
	// You can specify a maximum of 256 characters for a tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/appsync_channel_namespace#value AppsyncChannelNamespace#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

