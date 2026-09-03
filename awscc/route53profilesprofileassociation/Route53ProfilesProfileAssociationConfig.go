// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53profilesprofileassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53ProfilesProfileAssociationConfig struct {
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
	// The name of an association between a  Profile and a VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53profiles_profile_association#name Route53ProfilesProfileAssociation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the  profile that you associated with the resource that is specified by ResourceId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53profiles_profile_association#profile_id Route53ProfilesProfileAssociation#profile_id}
	ProfileId *string `field:"required" json:"profileId" yaml:"profileId"`
	// The resource that you associated the  profile with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53profiles_profile_association#resource_id Route53ProfilesProfileAssociation#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The Amazon Resource Name (ARN) of the profile association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53profiles_profile_association#arn Route53ProfilesProfileAssociation#arn}
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/route53profiles_profile_association#tags Route53ProfilesProfileAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

