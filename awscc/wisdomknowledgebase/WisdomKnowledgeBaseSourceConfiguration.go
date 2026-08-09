// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomknowledgebase


type WisdomKnowledgeBaseSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_knowledge_base#app_integrations WisdomKnowledgeBase#app_integrations}.
	AppIntegrations *WisdomKnowledgeBaseSourceConfigurationAppIntegrations `field:"optional" json:"appIntegrations" yaml:"appIntegrations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_knowledge_base#managed_source_configuration WisdomKnowledgeBase#managed_source_configuration}.
	ManagedSourceConfiguration *WisdomKnowledgeBaseSourceConfigurationManagedSourceConfiguration `field:"optional" json:"managedSourceConfiguration" yaml:"managedSourceConfiguration"`
}

