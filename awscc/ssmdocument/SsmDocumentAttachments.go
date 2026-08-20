// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmdocument


type SsmDocumentAttachments struct {
	// The key of a key-value pair that identifies the location of an attachment to a document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssm_document#key SsmDocument#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The name of the document attachment file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssm_document#name SsmDocument#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of a key-value pair that identifies the location of an attachment to a document.
	//
	// The format for Value depends on the type of key you specify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssm_document#values SsmDocument#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

