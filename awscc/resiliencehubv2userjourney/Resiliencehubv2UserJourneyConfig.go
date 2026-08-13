// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2userjourney

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Resiliencehubv2UserJourneyConfig struct {
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
	// The name of the user journey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/resiliencehubv2_user_journey#name Resiliencehubv2UserJourney#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The system ARN or system ID that owns this user journey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/resiliencehubv2_user_journey#system_identifier Resiliencehubv2UserJourney#system_identifier}
	SystemIdentifier *string `field:"required" json:"systemIdentifier" yaml:"systemIdentifier"`
	// The description of the user journey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/resiliencehubv2_user_journey#description Resiliencehubv2UserJourney#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ARN of the resilience policy to associate with this user journey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/resiliencehubv2_user_journey#policy_arn Resiliencehubv2UserJourney#policy_arn}
	PolicyArn *string `field:"optional" json:"policyArn" yaml:"policyArn"`
}

