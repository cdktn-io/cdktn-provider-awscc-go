// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationGitLab struct {
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_association#enable_webhook_updates DevopsagentAssociation#enable_webhook_updates}
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// GitLab instance identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_association#instance_identifier DevopsagentAssociation#instance_identifier}
	InstanceIdentifier *string `field:"optional" json:"instanceIdentifier" yaml:"instanceIdentifier"`
	// GitLab numeric project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_association#project_id DevopsagentAssociation#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Full GitLab project path (e.g., namespace/project-name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_association#project_path DevopsagentAssociation#project_path}
	ProjectPath *string `field:"optional" json:"projectPath" yaml:"projectPath"`
}

