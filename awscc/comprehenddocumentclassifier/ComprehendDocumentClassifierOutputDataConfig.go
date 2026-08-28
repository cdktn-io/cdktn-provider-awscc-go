// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehenddocumentclassifier


type ComprehendDocumentClassifierOutputDataConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/comprehend_document_classifier#kms_key_id ComprehendDocumentClassifier#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/comprehend_document_classifier#s3_uri ComprehendDocumentClassifier#s3_uri}.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

