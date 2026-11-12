// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricappbundle

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppfabricAppBundleConfig struct {
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
	// The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS) key to use to encrypt the application data.
	//
	// If this is not specified, an AWS owned key is used for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/appfabric_app_bundle#customer_managed_key_identifier AppfabricAppBundle#customer_managed_key_identifier}
	CustomerManagedKeyIdentifier *string `field:"optional" json:"customerManagedKeyIdentifier" yaml:"customerManagedKeyIdentifier"`
	// A map of the key-value pairs of the tag or tags to assign to the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/appfabric_app_bundle#tags AppfabricAppBundle#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

