// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appintegrationseventintegration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppintegrationsEventIntegrationConfig struct {
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
	// The Amazon Eventbridge bus for the event integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_event_integration#event_bridge_bus AppintegrationsEventIntegration#event_bridge_bus}
	EventBridgeBus *string `field:"required" json:"eventBridgeBus" yaml:"eventBridgeBus"`
	// The EventFilter (source) associated with the event integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_event_integration#event_filter AppintegrationsEventIntegration#event_filter}
	EventFilter *AppintegrationsEventIntegrationEventFilter `field:"required" json:"eventFilter" yaml:"eventFilter"`
	// The name of the event integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_event_integration#name AppintegrationsEventIntegration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The event integration description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_event_integration#description AppintegrationsEventIntegration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags (keys and values) associated with the event integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appintegrations_event_integration#tags AppintegrationsEventIntegration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

