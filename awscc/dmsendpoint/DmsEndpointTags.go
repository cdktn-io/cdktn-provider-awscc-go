// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint


type DmsEndpointTags struct {
	// A key is the required name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#key DmsEndpoint#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A value is the optional value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dms_endpoint#value DmsEndpoint#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

