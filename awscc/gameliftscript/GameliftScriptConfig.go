// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gameliftscript

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GameliftScriptConfig struct {
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
	// The location of the Amazon S3 bucket where a zipped file containing your Realtime scripts is stored.
	//
	// The storage location must specify the Amazon S3 bucket name, the zip file name (the "key"), and a role ARN that allows Amazon GameLift to access the Amazon S3 storage location. The S3 bucket must be in the same Region where you want to create a new script. By default, Amazon GameLift uploads the latest version of the zip file; if you have S3 object versioning turned on, you can use the ObjectVersion parameter to specify an earlier version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_script#storage_location GameliftScript#storage_location}
	StorageLocation *GameliftScriptStorageLocation `field:"required" json:"storageLocation" yaml:"storageLocation"`
	// A descriptive label that is associated with a script. Script names do not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_script#name GameliftScript#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Node.js version used for execution of the Realtime script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_script#node_js_version GameliftScript#node_js_version}
	NodeJsVersion *string `field:"optional" json:"nodeJsVersion" yaml:"nodeJsVersion"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_script#tags GameliftScript#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The version that is associated with a script. Version strings do not need to be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/gamelift_script#version GameliftScript#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

