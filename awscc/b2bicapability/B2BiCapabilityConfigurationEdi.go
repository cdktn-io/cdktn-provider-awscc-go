// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bicapability


type B2BiCapabilityConfigurationEdi struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/b2bi_capability#capability_direction B2BiCapability#capability_direction}.
	CapabilityDirection *string `field:"optional" json:"capabilityDirection" yaml:"capabilityDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/b2bi_capability#input_location B2BiCapability#input_location}.
	InputLocation *B2BiCapabilityConfigurationEdiInputLocation `field:"optional" json:"inputLocation" yaml:"inputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/b2bi_capability#output_location B2BiCapability#output_location}.
	OutputLocation *B2BiCapabilityConfigurationEdiOutputLocation `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/b2bi_capability#transformer_id B2BiCapability#transformer_id}.
	TransformerId *string `field:"optional" json:"transformerId" yaml:"transformerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/b2bi_capability#type B2BiCapability#type}.
	Type *B2BiCapabilityConfigurationEdiType `field:"optional" json:"type" yaml:"type"`
}

