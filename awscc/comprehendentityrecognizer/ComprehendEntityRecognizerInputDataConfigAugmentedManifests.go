// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehendentityrecognizer


type ComprehendEntityRecognizerInputDataConfigAugmentedManifests struct {
	// The S3 prefix to the annotation files that are referred in the augmented manifest file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer#annotation_data_s3_uri ComprehendEntityRecognizer#annotation_data_s3_uri}
	AnnotationDataS3Uri *string `field:"optional" json:"annotationDataS3Uri" yaml:"annotationDataS3Uri"`
	// The JSON attribute that contains the annotations for your training documents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer#attribute_names ComprehendEntityRecognizer#attribute_names}
	AttributeNames *[]*string `field:"optional" json:"attributeNames" yaml:"attributeNames"`
	// The type of augmented manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer#document_type ComprehendEntityRecognizer#document_type}
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// The Amazon S3 location of the augmented manifest file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer#s3_uri ComprehendEntityRecognizer#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// The S3 prefix to the source files (PDFs) that are referred to in the augmented manifest file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer#source_documents_s3_uri ComprehendEntityRecognizer#source_documents_s3_uri}
	SourceDocumentsS3Uri *string `field:"optional" json:"sourceDocumentsS3Uri" yaml:"sourceDocumentsS3Uri"`
	// The purpose of the data you've provided in the augmented manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/comprehend_entity_recognizer#split ComprehendEntityRecognizer#split}
	Split *string `field:"optional" json:"split" yaml:"split"`
}

