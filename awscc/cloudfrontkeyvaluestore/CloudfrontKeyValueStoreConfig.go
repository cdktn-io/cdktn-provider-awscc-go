// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontkeyvaluestore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudfrontKeyValueStoreConfig struct {
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
	// The name of the key value store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_key_value_store#name CloudfrontKeyValueStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A comment to describe the Key Value Store.
	//
	// Omitting ``Comment`` from the template during updates will clear the existing comment (set to empty string). To preserve an existing comment, you must explicitly include it in the template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_key_value_store#comment CloudfrontKeyValueStore#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// The import source for the key value store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_key_value_store#import_source CloudfrontKeyValueStore#import_source}
	ImportSource *CloudfrontKeyValueStoreImportSource `field:"optional" json:"importSource" yaml:"importSource"`
	// A complex type that contains zero or more ``Tag`` elements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_key_value_store#tags CloudfrontKeyValueStore#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

