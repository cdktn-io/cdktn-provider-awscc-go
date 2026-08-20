// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueclassifier


type GlueClassifierJsonClassifier struct {
	// A JsonPath string defining the JSON data for the classifier to classify.
	//
	// AWS Glue supports a subset of JsonPath, as described in Writing JsonPath Custom Classifiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_classifier#json_path GlueClassifier#json_path}
	JsonPath *string `field:"optional" json:"jsonPath" yaml:"jsonPath"`
	// The name of the classifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/glue_classifier#name GlueClassifier#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

