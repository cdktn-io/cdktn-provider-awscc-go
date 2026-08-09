// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmdocument

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmDocumentConfig struct {
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
	// The content for the Systems Manager document in JSON, YAML or String format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#content SsmDocument#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// A list of key and value pairs that describe attachments to a version of a document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#attachments SsmDocument#attachments}
	Attachments interface{} `field:"optional" json:"attachments" yaml:"attachments"`
	// Specify the document format for the request.
	//
	// The document format can be either JSON or YAML. JSON is the default format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#document_format SsmDocument#document_format}
	DocumentFormat *string `field:"optional" json:"documentFormat" yaml:"documentFormat"`
	// The type of document to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#document_type SsmDocument#document_type}
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// A name for the Systems Manager document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#name SsmDocument#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of SSM documents required by a document. For example, an ApplicationConfiguration document requires an ApplicationConfigurationSchema document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#requires SsmDocument#requires}
	Requires interface{} `field:"optional" json:"requires" yaml:"requires"`
	// Optional metadata that you assign to a resource.
	//
	// Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#tags SsmDocument#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specify a target type to define the kinds of resources the document can run on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#target_type SsmDocument#target_type}
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// Update method - when set to 'Replace', the update will replace the existing document;
	//
	// when set to 'NewVersion', the update will create a new version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#update_method SsmDocument#update_method}
	UpdateMethod *string `field:"optional" json:"updateMethod" yaml:"updateMethod"`
	// An optional field specifying the version of the artifact you are creating with the document.
	//
	// This value is unique across all versions of a document, and cannot be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ssm_document#version_name SsmDocument#version_name}
	VersionName *string `field:"optional" json:"versionName" yaml:"versionName"`
}

