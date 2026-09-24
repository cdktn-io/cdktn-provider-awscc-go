// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DevopsagentAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The unique identifier of the AgentSpace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#agent_space_id DevopsagentAssociation#agent_space_id}
	AgentSpaceId *string `field:"required" json:"agentSpaceId" yaml:"agentSpaceId"`
	// The configuration that directs how AgentSpace interacts with the given service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#configuration DevopsagentAssociation#configuration}
	Configuration *DevopsagentAssociationConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The identifier for the associated service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#service_id DevopsagentAssociation#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Set of linked association IDs for parent-child relationships.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_association#linked_association_ids DevopsagentAssociation#linked_association_ids}
	LinkedAssociationIds *[]*string `field:"optional" json:"linkedAssociationIds" yaml:"linkedAssociationIds"`
}

