// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionidnamespace


type EntityresolutionIdNamespaceIdMappingWorkflowPropertiesRuleBasedProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#attribute_matching_model EntityresolutionIdNamespace#attribute_matching_model}.
	AttributeMatchingModel *string `field:"optional" json:"attributeMatchingModel" yaml:"attributeMatchingModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#record_matching_models EntityresolutionIdNamespace#record_matching_models}.
	RecordMatchingModels *[]*string `field:"optional" json:"recordMatchingModels" yaml:"recordMatchingModels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#rule_definition_types EntityresolutionIdNamespace#rule_definition_types}.
	RuleDefinitionTypes *[]*string `field:"optional" json:"ruleDefinitionTypes" yaml:"ruleDefinitionTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#rules EntityresolutionIdNamespace#rules}.
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

