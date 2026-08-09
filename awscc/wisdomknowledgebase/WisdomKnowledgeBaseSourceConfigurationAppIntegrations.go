// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomknowledgebase


type WisdomKnowledgeBaseSourceConfigurationAppIntegrations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_knowledge_base#app_integration_arn WisdomKnowledgeBase#app_integration_arn}.
	AppIntegrationArn *string `field:"optional" json:"appIntegrationArn" yaml:"appIntegrationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/wisdom_knowledge_base#object_fields WisdomKnowledgeBase#object_fields}.
	ObjectFields *[]*string `field:"optional" json:"objectFields" yaml:"objectFields"`
}

