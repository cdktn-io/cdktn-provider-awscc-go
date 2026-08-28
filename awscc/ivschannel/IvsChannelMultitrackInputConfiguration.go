// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivschannel


type IvsChannelMultitrackInputConfiguration struct {
	// Indicates whether multitrack input is enabled.
	//
	// Can be set to true only if channel type is STANDARD. Setting enabled to true with any other channel type will cause an exception. If true, then policy, maximumResolution, and containerFormat are required, and containerFormat must be set to FRAGMENTED_MP4. Default: false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ivs_channel#enabled IvsChannel#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Maximum resolution for multitrack input. Required if enabled is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ivs_channel#maximum_resolution IvsChannel#maximum_resolution}
	MaximumResolution *string `field:"optional" json:"maximumResolution" yaml:"maximumResolution"`
	// Indicates whether multitrack input is allowed or required. Required if enabled is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ivs_channel#policy IvsChannel#policy}
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
}

