// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclink


type RtbfabricLinkModuleConfigurationListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_link#depends_on RtbfabricLink#depends_on}.
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_link#module_parameters RtbfabricLink#module_parameters}.
	ModuleParameters *RtbfabricLinkModuleConfigurationListModuleParameters `field:"optional" json:"moduleParameters" yaml:"moduleParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_link#name RtbfabricLink#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/rtbfabric_link#version RtbfabricLink#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

