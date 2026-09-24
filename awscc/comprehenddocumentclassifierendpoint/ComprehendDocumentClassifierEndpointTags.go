// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehenddocumentclassifierendpoint


type ComprehendDocumentClassifierEndpointTags struct {
	// The initial part of a key-value pair that forms a tag associated with a given resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_document_classifier_endpoint#key ComprehendDocumentClassifierEndpoint#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The second part of a key-value pair that forms a tag associated with a given resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/comprehend_document_classifier_endpoint#value ComprehendDocumentClassifierEndpoint#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

