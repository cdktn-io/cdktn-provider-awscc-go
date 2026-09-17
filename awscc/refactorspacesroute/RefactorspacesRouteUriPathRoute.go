// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package refactorspacesroute


type RefactorspacesRouteUriPathRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/refactorspaces_route#activation_state RefactorspacesRoute#activation_state}.
	ActivationState *string `field:"optional" json:"activationState" yaml:"activationState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/refactorspaces_route#append_source_path RefactorspacesRoute#append_source_path}.
	AppendSourcePath interface{} `field:"optional" json:"appendSourcePath" yaml:"appendSourcePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/refactorspaces_route#include_child_paths RefactorspacesRoute#include_child_paths}.
	IncludeChildPaths interface{} `field:"optional" json:"includeChildPaths" yaml:"includeChildPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/refactorspaces_route#methods RefactorspacesRoute#methods}.
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/refactorspaces_route#source_path RefactorspacesRoute#source_path}.
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

