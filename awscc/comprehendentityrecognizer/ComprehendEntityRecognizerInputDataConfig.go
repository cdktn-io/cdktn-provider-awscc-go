// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehendentityrecognizer


type ComprehendEntityRecognizerInputDataConfig struct {
	// The entity types in the labeled training data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/comprehend_entity_recognizer#entity_types ComprehendEntityRecognizer#entity_types}
	EntityTypes interface{} `field:"required" json:"entityTypes" yaml:"entityTypes"`
	// The S3 location of the CSV file that annotates your training documents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/comprehend_entity_recognizer#annotations ComprehendEntityRecognizer#annotations}
	Annotations *ComprehendEntityRecognizerInputDataConfigAnnotations `field:"optional" json:"annotations" yaml:"annotations"`
	// A list of augmented manifest files that provide training data for a custom model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/comprehend_entity_recognizer#augmented_manifests ComprehendEntityRecognizer#augmented_manifests}
	AugmentedManifests interface{} `field:"optional" json:"augmentedManifests" yaml:"augmentedManifests"`
	// The format of your training data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/comprehend_entity_recognizer#data_format ComprehendEntityRecognizer#data_format}
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// The S3 location of the folder that contains the training documents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/comprehend_entity_recognizer#documents ComprehendEntityRecognizer#documents}
	Documents *ComprehendEntityRecognizerInputDataConfigDocuments `field:"optional" json:"documents" yaml:"documents"`
	// The S3 location of the CSV file that has the entity list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/comprehend_entity_recognizer#entity_list ComprehendEntityRecognizer#entity_list}
	EntityList *ComprehendEntityRecognizerInputDataConfigEntityListStruct `field:"optional" json:"entityList" yaml:"entityList"`
}

