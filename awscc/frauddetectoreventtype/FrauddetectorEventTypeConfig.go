// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package frauddetectoreventtype

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FrauddetectorEventTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_event_type#entity_types FrauddetectorEventType#entity_types}.
	EntityTypes interface{} `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_event_type#event_variables FrauddetectorEventType#event_variables}.
	EventVariables interface{} `field:"required" json:"eventVariables" yaml:"eventVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_event_type#labels FrauddetectorEventType#labels}.
	Labels interface{} `field:"required" json:"labels" yaml:"labels"`
	// The name for the event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_event_type#name FrauddetectorEventType#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_event_type#description FrauddetectorEventType#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Tags associated with this event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/frauddetector_event_type#tags FrauddetectorEventType#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

