// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentagentspace


type SecurityagentAgentSpaceIntegratedResourcesProviderResourcesConfluenceCapabilities struct {
	// Enables creation of new Confluence documents in the same space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#create_document SecurityagentAgentSpace#create_document}
	CreateDocument interface{} `field:"optional" json:"createDocument" yaml:"createDocument"`
	// Enables read access to the document content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#fetch_document SecurityagentAgentSpace#fetch_document}
	FetchDocument interface{} `field:"optional" json:"fetchDocument" yaml:"fetchDocument"`
	// Enables updates to the document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/securityagent_agent_space#update_document SecurityagentAgentSpace#update_document}
	UpdateDocument interface{} `field:"optional" json:"updateDocument" yaml:"updateDocument"`
}

