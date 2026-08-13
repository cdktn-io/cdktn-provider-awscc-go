// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraindex


type KendraIndexDocumentMetadataConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_index#name KendraIndex#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_index#relevance KendraIndex#relevance}.
	Relevance *KendraIndexDocumentMetadataConfigurationsRelevance `field:"optional" json:"relevance" yaml:"relevance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_index#search KendraIndex#search}.
	Search *KendraIndexDocumentMetadataConfigurationsSearch `field:"optional" json:"search" yaml:"search"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/kendra_index#type KendraIndex#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

