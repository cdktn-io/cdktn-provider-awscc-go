// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamstack


type AppstreamStackUserSettings struct {
	// The action that is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#action AppstreamStack#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Specifies the number of characters that can be copied by end users from the local device to the remote session, and to the local device from the remote session.
	//
	// This can be specified only for the CLIPBOARD_COPY_FROM_LOCAL_DEVICE and CLIPBOARD_COPY_TO_LOCAL_DEVICE actions. This defaults to 20,971,520 (20 MB) when unspecified and the permission is ENABLED. This can't be specified when the permission is DISABLED. The value can be between 1 and 20,971,520 (20 MB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#maximum_length AppstreamStack#maximum_length}
	MaximumLength *float64 `field:"optional" json:"maximumLength" yaml:"maximumLength"`
	// Indicates whether the action is enabled or disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appstream_stack#permission AppstreamStack#permission}
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

