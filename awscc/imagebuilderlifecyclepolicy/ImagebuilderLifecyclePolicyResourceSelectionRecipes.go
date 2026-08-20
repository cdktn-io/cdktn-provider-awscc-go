// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyResourceSelectionRecipes struct {
	// The recipe name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_lifecycle_policy#name ImagebuilderLifecyclePolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The recipe version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_lifecycle_policy#semantic_version ImagebuilderLifecyclePolicy#semantic_version}
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
}

