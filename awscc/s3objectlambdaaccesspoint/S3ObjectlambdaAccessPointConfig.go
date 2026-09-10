// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3objectlambdaaccesspoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3ObjectlambdaAccessPointConfig struct {
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
	// The Object lambda Access Point Configuration that configures transformations to be applied on the objects on specified S3 Actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3objectlambda_access_point#object_lambda_configuration S3ObjectlambdaAccessPoint#object_lambda_configuration}
	ObjectLambdaConfiguration *S3ObjectlambdaAccessPointObjectLambdaConfiguration `field:"required" json:"objectLambdaConfiguration" yaml:"objectLambdaConfiguration"`
	// The name you want to assign to this Object lambda Access Point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/s3objectlambda_access_point#name S3ObjectlambdaAccessPoint#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

