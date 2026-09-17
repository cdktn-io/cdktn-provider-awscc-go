// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedreviewtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedReviewTemplateConfig struct {
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
	// The review template description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wellarchitected_review_template#description WellarchitectedReviewTemplate#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The lenses applied to the review template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wellarchitected_review_template#lenses WellarchitectedReviewTemplate#lenses}
	Lenses *[]*string `field:"required" json:"lenses" yaml:"lenses"`
	// The name of the review template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wellarchitected_review_template#template_name WellarchitectedReviewTemplate#template_name}
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// The notes associated with the review template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wellarchitected_review_template#notes WellarchitectedReviewTemplate#notes}
	Notes *string `field:"optional" json:"notes" yaml:"notes"`
	// The tags assigned to the review template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wellarchitected_review_template#tags WellarchitectedReviewTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

