// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightSpaceConfig struct {
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
	// The ID of the Amazon Web Services account where the space is being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_space#aws_account_id QuicksightSpace#aws_account_id}
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The display name of the space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_space#name QuicksightSpace#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The unique identifier for the space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_space#space_id QuicksightSpace#space_id}
	SpaceId *string `field:"required" json:"spaceId" yaml:"spaceId"`
	// A description of the space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_space#description QuicksightSpace#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of permissions granted on the space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_space#permissions QuicksightSpace#permissions}
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// A list of QuickSight resources attached to the space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_space#resources QuicksightSpace#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// A list of key-value pairs to associate with the space resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_space#tags QuicksightSpace#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

