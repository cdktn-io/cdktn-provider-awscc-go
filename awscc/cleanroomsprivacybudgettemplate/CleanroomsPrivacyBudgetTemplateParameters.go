// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsprivacybudgettemplate


type CleanroomsPrivacyBudgetTemplateParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_privacy_budget_template#budget_parameters CleanroomsPrivacyBudgetTemplate#budget_parameters}.
	BudgetParameters interface{} `field:"optional" json:"budgetParameters" yaml:"budgetParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_privacy_budget_template#epsilon CleanroomsPrivacyBudgetTemplate#epsilon}.
	Epsilon *float64 `field:"optional" json:"epsilon" yaml:"epsilon"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_privacy_budget_template#resource_arn CleanroomsPrivacyBudgetTemplate#resource_arn}.
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cleanrooms_privacy_budget_template#users_noise_per_query CleanroomsPrivacyBudgetTemplate#users_noise_per_query}.
	UsersNoisePerQuery *float64 `field:"optional" json:"usersNoisePerQuery" yaml:"usersNoisePerQuery"`
}

