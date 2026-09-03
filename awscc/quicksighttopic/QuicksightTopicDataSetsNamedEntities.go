// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksighttopic


type QuicksightTopicDataSetsNamedEntities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#definition QuicksightTopic#definition}.
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#entity_description QuicksightTopic#entity_description}.
	EntityDescription *string `field:"optional" json:"entityDescription" yaml:"entityDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#entity_name QuicksightTopic#entity_name}.
	EntityName *string `field:"optional" json:"entityName" yaml:"entityName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#entity_synonyms QuicksightTopic#entity_synonyms}.
	EntitySynonyms *[]*string `field:"optional" json:"entitySynonyms" yaml:"entitySynonyms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/quicksight_topic#semantic_entity_type QuicksightTopic#semantic_entity_type}.
	SemanticEntityType *QuicksightTopicDataSetsNamedEntitiesSemanticEntityType `field:"optional" json:"semanticEntityType" yaml:"semanticEntityType"`
}

