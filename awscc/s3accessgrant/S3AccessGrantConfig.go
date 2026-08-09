// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3accessgrant

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3AccessGrantConfig struct {
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
	// The custom S3 location to be accessed by the grantee.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_grant#access_grants_location_id S3AccessGrant#access_grants_location_id}
	AccessGrantsLocationId *string `field:"required" json:"accessGrantsLocationId" yaml:"accessGrantsLocationId"`
	// The principal who will be granted permission to access S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_grant#grantee S3AccessGrant#grantee}
	Grantee *S3AccessGrantGrantee `field:"required" json:"grantee" yaml:"grantee"`
	// The level of access to be afforded to the grantee.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_grant#permission S3AccessGrant#permission}
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// The configuration options of the grant location, which is the S3 path to the data to which you are granting access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_grant#access_grants_location_configuration S3AccessGrant#access_grants_location_configuration}
	AccessGrantsLocationConfiguration *S3AccessGrantAccessGrantsLocationConfiguration `field:"optional" json:"accessGrantsLocationConfiguration" yaml:"accessGrantsLocationConfiguration"`
	// The ARN of the application grantees will use to access the location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_grant#application_arn S3AccessGrant#application_arn}
	ApplicationArn *string `field:"optional" json:"applicationArn" yaml:"applicationArn"`
	// The type of S3SubPrefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_grant#s3_prefix_type S3AccessGrant#s3_prefix_type}
	S3PrefixType *string `field:"optional" json:"s3PrefixType" yaml:"s3PrefixType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/s3_access_grant#tags S3AccessGrant#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

