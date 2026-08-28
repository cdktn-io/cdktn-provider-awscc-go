// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appstreamappblock

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppstreamAppBlockConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_app_block#name AppstreamAppBlock#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_app_block#source_s3_location AppstreamAppBlock#source_s3_location}.
	SourceS3Location *AppstreamAppBlockSourceS3Location `field:"required" json:"sourceS3Location" yaml:"sourceS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_app_block#description AppstreamAppBlock#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_app_block#display_name AppstreamAppBlock#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_app_block#packaging_type AppstreamAppBlock#packaging_type}.
	PackagingType *string `field:"optional" json:"packagingType" yaml:"packagingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_app_block#post_setup_script_details AppstreamAppBlock#post_setup_script_details}.
	PostSetupScriptDetails *AppstreamAppBlockPostSetupScriptDetails `field:"optional" json:"postSetupScriptDetails" yaml:"postSetupScriptDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_app_block#setup_script_details AppstreamAppBlock#setup_script_details}.
	SetupScriptDetails *AppstreamAppBlockSetupScriptDetails `field:"optional" json:"setupScriptDetails" yaml:"setupScriptDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/appstream_app_block#tags AppstreamAppBlock#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

