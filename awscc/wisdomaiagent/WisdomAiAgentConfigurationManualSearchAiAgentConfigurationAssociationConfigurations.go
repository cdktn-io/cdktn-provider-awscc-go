// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent


type WisdomAiAgentConfigurationManualSearchAiAgentConfigurationAssociationConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#association_configuration_data WisdomAiAgent#association_configuration_data}.
	AssociationConfigurationData *WisdomAiAgentConfigurationManualSearchAiAgentConfigurationAssociationConfigurationsAssociationConfigurationData `field:"optional" json:"associationConfigurationData" yaml:"associationConfigurationData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#association_id WisdomAiAgent#association_id}.
	AssociationId *string `field:"optional" json:"associationId" yaml:"associationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_ai_agent#association_type WisdomAiAgent#association_type}.
	AssociationType *string `field:"optional" json:"associationType" yaml:"associationType"`
}

