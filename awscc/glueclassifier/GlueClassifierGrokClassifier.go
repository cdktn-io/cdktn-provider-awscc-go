// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueclassifier


type GlueClassifierGrokClassifier struct {
	// An identifier of the data format that the classifier matches, such as Twitter, JSON, Omniture logs, and so on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_classifier#classification GlueClassifier#classification}
	Classification *string `field:"optional" json:"classification" yaml:"classification"`
	// Optional custom grok patterns defined by this classifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_classifier#custom_patterns GlueClassifier#custom_patterns}
	CustomPatterns *string `field:"optional" json:"customPatterns" yaml:"customPatterns"`
	// The grok pattern applied to a data store by this classifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_classifier#grok_pattern GlueClassifier#grok_pattern}
	GrokPattern *string `field:"optional" json:"grokPattern" yaml:"grokPattern"`
	// The name of the classifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/glue_classifier#name GlueClassifier#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

