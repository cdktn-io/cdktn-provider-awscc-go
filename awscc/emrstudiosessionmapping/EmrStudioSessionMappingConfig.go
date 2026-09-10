// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrstudiosessionmapping

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EmrStudioSessionMappingConfig struct {
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
	// The name of the user or group.
	//
	// For more information, see UserName and DisplayName in the AWS SSO Identity Store API Reference. Either IdentityName or IdentityId must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_studio_session_mapping#identity_name EmrStudioSessionMapping#identity_name}
	IdentityName *string `field:"required" json:"identityName" yaml:"identityName"`
	// Specifies whether the identity to map to the Studio is a user or a group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_studio_session_mapping#identity_type EmrStudioSessionMapping#identity_type}
	IdentityType *string `field:"required" json:"identityType" yaml:"identityType"`
	// The Amazon Resource Name (ARN) for the session policy that will be applied to the user or group.
	//
	// Session policies refine Studio user permissions without the need to use multiple IAM user roles.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_studio_session_mapping#session_policy_arn EmrStudioSessionMapping#session_policy_arn}
	SessionPolicyArn *string `field:"required" json:"sessionPolicyArn" yaml:"sessionPolicyArn"`
	// The ID of the Amazon EMR Studio to which the user or group will be mapped.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emr_studio_session_mapping#studio_id EmrStudioSessionMapping#studio_id}
	StudioId *string `field:"required" json:"studioId" yaml:"studioId"`
}

