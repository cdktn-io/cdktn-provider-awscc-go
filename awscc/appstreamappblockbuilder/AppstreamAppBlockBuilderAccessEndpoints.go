// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamappblockbuilder


type AppstreamAppBlockBuilderAccessEndpoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appstream_app_block_builder#endpoint_type AppstreamAppBlockBuilder#endpoint_type}.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appstream_app_block_builder#vpce_id AppstreamAppBlockBuilder#vpce_id}.
	VpceId *string `field:"optional" json:"vpceId" yaml:"vpceId"`
}

