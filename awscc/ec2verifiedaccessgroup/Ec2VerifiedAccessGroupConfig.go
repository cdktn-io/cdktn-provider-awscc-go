// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VerifiedAccessGroupConfig struct {
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
	// The ID of the AWS Verified Access instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_group#verified_access_instance_id Ec2VerifiedAccessGroup#verified_access_instance_id}
	VerifiedAccessInstanceId *string `field:"required" json:"verifiedAccessInstanceId" yaml:"verifiedAccessInstanceId"`
	// A description for the AWS Verified Access group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_group#description Ec2VerifiedAccessGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS Verified Access policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_group#policy_document Ec2VerifiedAccessGroup#policy_document}
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The status of the Verified Access policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_group#policy_enabled Ec2VerifiedAccessGroup#policy_enabled}
	PolicyEnabled interface{} `field:"optional" json:"policyEnabled" yaml:"policyEnabled"`
	// The configuration options for customer provided KMS encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_group#sse_specification Ec2VerifiedAccessGroup#sse_specification}
	SseSpecification *Ec2VerifiedAccessGroupSseSpecification `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_group#tags Ec2VerifiedAccessGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

