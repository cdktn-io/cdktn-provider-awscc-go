// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2channel


type Mediapackagev2ChannelOutputHeaderConfiguration struct {
	// <p>When true, AWS Elemental MediaPackage includes the MQCS in responses to the CDN.
	//
	// This setting is valid only when <code>InputType</code> is <code>CMAF</code>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mediapackagev2_channel#publish_mqcs Mediapackagev2Channel#publish_mqcs}
	PublishMqcs interface{} `field:"optional" json:"publishMqcs" yaml:"publishMqcs"`
}

