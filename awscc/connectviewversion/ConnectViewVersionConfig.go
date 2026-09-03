// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectviewversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectViewVersionConfig struct {
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
	// The Amazon Resource Name (ARN) of the view for which a version is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_view_version#view_arn ConnectViewVersion#view_arn}
	ViewArn *string `field:"required" json:"viewArn" yaml:"viewArn"`
	// The description for the view version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_view_version#version_description ConnectViewVersion#version_description}
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
	// The view content hash to be checked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/connect_view_version#view_content_sha_256 ConnectViewVersion#view_content_sha_256}
	ViewContentSha256 *string `field:"optional" json:"viewContentSha256" yaml:"viewContentSha256"`
}

