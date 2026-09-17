// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package frauddetectordetector


type FrauddetectorDetectorEventType struct {
	// The description of the event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/frauddetector_detector#description FrauddetectorDetector#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/frauddetector_detector#entity_types FrauddetectorDetector#entity_types}.
	EntityTypes interface{} `field:"optional" json:"entityTypes" yaml:"entityTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/frauddetector_detector#event_variables FrauddetectorDetector#event_variables}.
	EventVariables interface{} `field:"optional" json:"eventVariables" yaml:"eventVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/frauddetector_detector#inline FrauddetectorDetector#inline}.
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/frauddetector_detector#labels FrauddetectorDetector#labels}.
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// The name for the event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/frauddetector_detector#name FrauddetectorDetector#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags associated with this event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/frauddetector_detector#tags FrauddetectorDetector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

