// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elementalinferencefeed

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElementalinferenceFeedConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elementalinference_feed#name ElementalinferenceFeed#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elementalinference_feed#outputs ElementalinferenceFeed#outputs}.
	Outputs interface{} `field:"required" json:"outputs" yaml:"outputs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elementalinference_feed#access_role_arn ElementalinferenceFeed#access_role_arn}.
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elementalinference_feed#tags ElementalinferenceFeed#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

