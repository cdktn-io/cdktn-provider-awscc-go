// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediapackagev2channel


type Mediapackagev2ChannelInputSwitchConfiguration struct {
	// <p>When true, AWS Elemental MediaPackage performs input switching based on the MQCS.
	//
	// Default is false. This setting is valid only when <code>InputType</code> is <code>CMAF</code>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediapackagev2_channel#mqcs_input_switching Mediapackagev2Channel#mqcs_input_switching}
	MqcsInputSwitching interface{} `field:"optional" json:"mqcsInputSwitching" yaml:"mqcsInputSwitching"`
	// <p>For CMAF inputs, indicates which input MediaPackage should prefer when both inputs have equal MQCS scores.
	//
	// Select <code>1</code> to prefer the first ingest endpoint, or <code>2</code> to prefer the second ingest endpoint. If you don't specify a preferred input, MediaPackage uses its default switching behavior when MQCS scores are equal.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediapackagev2_channel#preferred_input Mediapackagev2Channel#preferred_input}
	PreferredInput *float64 `field:"optional" json:"preferredInput" yaml:"preferredInput"`
}

