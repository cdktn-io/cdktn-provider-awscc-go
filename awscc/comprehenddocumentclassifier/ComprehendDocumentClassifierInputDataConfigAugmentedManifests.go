// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehenddocumentclassifier


type ComprehendDocumentClassifierInputDataConfigAugmentedManifests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/comprehend_document_classifier#attribute_names ComprehendDocumentClassifier#attribute_names}.
	AttributeNames *[]*string `field:"optional" json:"attributeNames" yaml:"attributeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/comprehend_document_classifier#s3_uri ComprehendDocumentClassifier#s3_uri}.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/comprehend_document_classifier#split ComprehendDocumentClassifier#split}.
	Split *string `field:"optional" json:"split" yaml:"split"`
}

