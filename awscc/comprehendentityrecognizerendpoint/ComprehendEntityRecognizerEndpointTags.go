// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehendentityrecognizerendpoint


type ComprehendEntityRecognizerEndpointTags struct {
	// The initial part of a key-value pair that forms a tag associated with a given resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/comprehend_entity_recognizer_endpoint#key ComprehendEntityRecognizerEndpoint#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The second part of a key-value pair that forms a tag associated with a given resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/comprehend_entity_recognizer_endpoint#value ComprehendEntityRecognizerEndpoint#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

