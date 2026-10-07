// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package translateterminology

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TranslateTerminologyConfig struct {
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
	// The name of the custom terminology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/translate_terminology#name TranslateTerminology#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the custom terminology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/translate_terminology#description TranslateTerminology#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The encryption key for the custom terminology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/translate_terminology#encryption_key TranslateTerminology#encryption_key}
	EncryptionKey *TranslateTerminologyEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The merge strategy for the custom terminology. Currently only OVERWRITE is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/translate_terminology#merge_strategy TranslateTerminology#merge_strategy}
	MergeStrategy *string `field:"optional" json:"mergeStrategy" yaml:"mergeStrategy"`
	// Tags associated with the terminology.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/translate_terminology#tags TranslateTerminology#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The terminology data for the custom terminology being imported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/translate_terminology#terminology_data TranslateTerminology#terminology_data}
	TerminologyData *TranslateTerminologyTerminologyData `field:"optional" json:"terminologyData" yaml:"terminologyData"`
}

