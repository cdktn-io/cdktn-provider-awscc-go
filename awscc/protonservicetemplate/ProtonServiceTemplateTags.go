// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package protonservicetemplate


type ProtonServiceTemplateTags struct {
	// <p>The key of the resource tag.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/proton_service_template#key ProtonServiceTemplate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// <p>The value of the resource tag.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/proton_service_template#value ProtonServiceTemplate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

