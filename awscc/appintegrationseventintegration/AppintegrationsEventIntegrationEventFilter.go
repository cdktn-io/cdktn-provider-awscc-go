// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationseventintegration


type AppintegrationsEventIntegrationEventFilter struct {
	// The source of the events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/appintegrations_event_integration#source AppintegrationsEventIntegration#source}
	Source *string `field:"required" json:"source" yaml:"source"`
}

