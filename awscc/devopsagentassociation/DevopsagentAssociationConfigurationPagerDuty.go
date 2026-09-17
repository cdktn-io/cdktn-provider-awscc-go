// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentassociation


type DevopsagentAssociationConfigurationPagerDuty struct {
	// Email to be used in PagerDuty API header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_association#customer_email DevopsagentAssociation#customer_email}
	CustomerEmail *string `field:"optional" json:"customerEmail" yaml:"customerEmail"`
	// When set to true, enables the Agent Space to create and update webhooks for receiving notifications and events from the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_association#enable_webhook_updates DevopsagentAssociation#enable_webhook_updates}
	EnableWebhookUpdates interface{} `field:"optional" json:"enableWebhookUpdates" yaml:"enableWebhookUpdates"`
	// List of PagerDuty service IDs available for the association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/devopsagent_association#services DevopsagentAssociation#services}
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
}

