// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialiveinput


type MedialiveInputMulticastSettingsSources struct {
	// This represents the ip address of the device sending the multicast stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#source_ip MedialiveInput#source_ip}
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// This represents the customer's source URL where multicast stream is pulled from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/medialive_input#url MedialiveInput#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

