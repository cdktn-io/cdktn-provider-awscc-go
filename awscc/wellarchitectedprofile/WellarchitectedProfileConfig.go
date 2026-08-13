// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedprofile

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedProfileConfig struct {
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
	// The profile description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wellarchitected_profile#profile_description WellarchitectedProfile#profile_description}
	ProfileDescription *string `field:"required" json:"profileDescription" yaml:"profileDescription"`
	// The name of the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wellarchitected_profile#profile_name WellarchitectedProfile#profile_name}
	ProfileName *string `field:"required" json:"profileName" yaml:"profileName"`
	// The profile questions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wellarchitected_profile#profile_questions WellarchitectedProfile#profile_questions}
	ProfileQuestions interface{} `field:"required" json:"profileQuestions" yaml:"profileQuestions"`
	// The tags assigned to the profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wellarchitected_profile#tags WellarchitectedProfile#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

