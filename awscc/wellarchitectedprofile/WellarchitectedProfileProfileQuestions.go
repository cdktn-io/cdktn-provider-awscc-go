// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedprofile


type WellarchitectedProfileProfileQuestions struct {
	// The ID of the question.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_profile#question_id WellarchitectedProfile#question_id}
	QuestionId *string `field:"optional" json:"questionId" yaml:"questionId"`
	// The selected choices.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/wellarchitected_profile#selected_choice_ids WellarchitectedProfile#selected_choice_ids}
	SelectedChoiceIds *[]*string `field:"optional" json:"selectedChoiceIds" yaml:"selectedChoiceIds"`
}

