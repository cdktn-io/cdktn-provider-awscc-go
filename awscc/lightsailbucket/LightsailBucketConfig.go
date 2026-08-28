// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailbucket

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailBucketConfig struct {
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
	// The name for the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_bucket#bucket_name LightsailBucket#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The ID of the bundle to use for the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_bucket#bundle_id LightsailBucket#bundle_id}
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// An object that sets the public accessibility of objects in the specified bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_bucket#access_rules LightsailBucket#access_rules}
	AccessRules *LightsailBucketAccessRules `field:"optional" json:"accessRules" yaml:"accessRules"`
	// Specifies whether to enable or disable versioning of objects in the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_bucket#object_versioning LightsailBucket#object_versioning}
	ObjectVersioning interface{} `field:"optional" json:"objectVersioning" yaml:"objectVersioning"`
	// An array of strings to specify the AWS account IDs that can access the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_bucket#read_only_access_accounts LightsailBucket#read_only_access_accounts}
	ReadOnlyAccessAccounts *[]*string `field:"optional" json:"readOnlyAccessAccounts" yaml:"readOnlyAccessAccounts"`
	// The names of the Lightsail resources for which to set bucket access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_bucket#resources_receiving_access LightsailBucket#resources_receiving_access}
	ResourcesReceivingAccess *[]*string `field:"optional" json:"resourcesReceivingAccess" yaml:"resourcesReceivingAccess"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_bucket#tags LightsailBucket#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

