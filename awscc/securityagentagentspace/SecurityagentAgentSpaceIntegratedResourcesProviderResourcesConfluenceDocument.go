// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceDocument struct {
	// Customer-supplied logical name for the Confluence document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityagent_agent_space#name SecurityagentAgentSpace#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Confluence page identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityagent_agent_space#page_id SecurityagentAgentSpace#page_id}
	PageId *string `field:"optional" json:"pageId" yaml:"pageId"`
	// Confluence space key containing the document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityagent_agent_space#space_key SecurityagentAgentSpace#space_key}
	SpaceKey *string `field:"optional" json:"spaceKey" yaml:"spaceKey"`
	// Read-only human-readable title of the containing space, populated from service-side metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityagent_agent_space#space_title SecurityagentAgentSpace#space_title}
	SpaceTitle *string `field:"optional" json:"spaceTitle" yaml:"spaceTitle"`
	// Read-only human-readable title of the page, populated from service-side metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/securityagent_agent_space#title SecurityagentAgentSpace#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

