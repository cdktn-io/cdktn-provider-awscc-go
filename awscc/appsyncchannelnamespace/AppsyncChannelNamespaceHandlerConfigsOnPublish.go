// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncchannelnamespace


type AppsyncChannelNamespaceHandlerConfigsOnPublish struct {
	// Integration behavior for a handler configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appsync_channel_namespace#behavior AppsyncChannelNamespace#behavior}
	Behavior *string `field:"optional" json:"behavior" yaml:"behavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appsync_channel_namespace#integration AppsyncChannelNamespace#integration}.
	Integration *AppsyncChannelNamespaceHandlerConfigsOnPublishIntegration `field:"optional" json:"integration" yaml:"integration"`
}

