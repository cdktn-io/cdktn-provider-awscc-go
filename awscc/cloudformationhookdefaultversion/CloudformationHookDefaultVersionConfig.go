// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationhookdefaultversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudformationHookDefaultVersionConfig struct {
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
	// The name of the type being registered.
	//
	// We recommend that type names adhere to the following pattern: company_or_organization::service::type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_hook_default_version#type_name CloudformationHookDefaultVersion#type_name}
	TypeName *string `field:"optional" json:"typeName" yaml:"typeName"`
	// The Amazon Resource Name (ARN) of the type version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_hook_default_version#type_version_arn CloudformationHookDefaultVersion#type_version_arn}
	TypeVersionArn *string `field:"optional" json:"typeVersionArn" yaml:"typeVersionArn"`
	// The ID of an existing version of the hook to set as the default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_hook_default_version#version_id CloudformationHookDefaultVersion#version_id}
	VersionId *string `field:"optional" json:"versionId" yaml:"versionId"`
}

