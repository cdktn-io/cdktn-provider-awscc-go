// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3storagelens

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3StorageLensConfig struct {
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
	// Specifies the details of Amazon S3 Storage Lens configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#storage_lens_configuration S3StorageLens#storage_lens_configuration}
	StorageLensConfiguration *S3StorageLensStorageLensConfiguration `field:"required" json:"storageLensConfiguration" yaml:"storageLensConfiguration"`
	// A set of tags (key-value pairs) for this Amazon S3 Storage Lens configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3_storage_lens#tags S3StorageLens#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

