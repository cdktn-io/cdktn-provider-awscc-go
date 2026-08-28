// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraindex


type KendraIndexDocumentMetadataConfigurationsRelevance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_index#duration KendraIndex#duration}.
	Duration *string `field:"optional" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_index#freshness KendraIndex#freshness}.
	Freshness interface{} `field:"optional" json:"freshness" yaml:"freshness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_index#importance KendraIndex#importance}.
	Importance *float64 `field:"optional" json:"importance" yaml:"importance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_index#rank_order KendraIndex#rank_order}.
	RankOrder *string `field:"optional" json:"rankOrder" yaml:"rankOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/kendra_index#value_importance_items KendraIndex#value_importance_items}.
	ValueImportanceItems interface{} `field:"optional" json:"valueImportanceItems" yaml:"valueImportanceItems"`
}

