// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueclassifier

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueClassifierConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A classifier for comma-separated values (CSV).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_classifier#csv_classifier GlueClassifier#csv_classifier}
	CsvClassifier *GlueClassifierCsvClassifier `field:"optional" json:"csvClassifier" yaml:"csvClassifier"`
	// A classifier that uses grok.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_classifier#grok_classifier GlueClassifier#grok_classifier}
	GrokClassifier *GlueClassifierGrokClassifier `field:"optional" json:"grokClassifier" yaml:"grokClassifier"`
	// A classifier for JSON content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_classifier#json_classifier GlueClassifier#json_classifier}
	JsonClassifier *GlueClassifierJsonClassifier `field:"optional" json:"jsonClassifier" yaml:"jsonClassifier"`
	// A classifier for XML content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_classifier#xml_classifier GlueClassifier#xml_classifier}
	XmlClassifier *GlueClassifierXmlClassifier `field:"optional" json:"xmlClassifier" yaml:"xmlClassifier"`
}

