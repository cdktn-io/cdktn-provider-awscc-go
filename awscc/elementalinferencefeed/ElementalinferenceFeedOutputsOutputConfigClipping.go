// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elementalinferencefeed


type ElementalinferenceFeedOutputsOutputConfigClipping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elementalinference_feed#callback_metadata ElementalinferenceFeed#callback_metadata}.
	CallbackMetadata *string `field:"optional" json:"callbackMetadata" yaml:"callbackMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elementalinference_feed#data_source_configuration ElementalinferenceFeed#data_source_configuration}.
	DataSourceConfiguration *ElementalinferenceFeedOutputsOutputConfigClippingDataSourceConfiguration `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
}

