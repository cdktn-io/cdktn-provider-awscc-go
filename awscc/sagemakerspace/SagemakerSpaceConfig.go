// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerSpaceConfig struct {
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
	// The ID of the associated Domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_space#domain_id SagemakerSpace#domain_id}
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// A name for the Space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_space#space_name SagemakerSpace#space_name}
	SpaceName *string `field:"required" json:"spaceName" yaml:"spaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_space#ownership_settings SagemakerSpace#ownership_settings}.
	OwnershipSettings *SagemakerSpaceOwnershipSettings `field:"optional" json:"ownershipSettings" yaml:"ownershipSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_space#space_display_name SagemakerSpace#space_display_name}.
	SpaceDisplayName *string `field:"optional" json:"spaceDisplayName" yaml:"spaceDisplayName"`
	// A collection of settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_space#space_settings SagemakerSpace#space_settings}
	SpaceSettings *SagemakerSpaceSpaceSettings `field:"optional" json:"spaceSettings" yaml:"spaceSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_space#space_sharing_settings SagemakerSpace#space_sharing_settings}.
	SpaceSharingSettings *SagemakerSpaceSpaceSharingSettings `field:"optional" json:"spaceSharingSettings" yaml:"spaceSharingSettings"`
	// A list of tags to apply to the space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_space#tags SagemakerSpace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

