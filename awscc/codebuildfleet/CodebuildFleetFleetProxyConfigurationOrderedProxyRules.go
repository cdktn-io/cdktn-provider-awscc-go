// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codebuildfleet


type CodebuildFleetFleetProxyConfigurationOrderedProxyRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codebuild_fleet#effect CodebuildFleet#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codebuild_fleet#entities CodebuildFleet#entities}.
	Entities *[]*string `field:"optional" json:"entities" yaml:"entities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codebuild_fleet#type CodebuildFleet#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

