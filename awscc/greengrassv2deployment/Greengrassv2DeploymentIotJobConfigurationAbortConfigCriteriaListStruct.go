// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment


type Greengrassv2DeploymentIotJobConfigurationAbortConfigCriteriaListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#action Greengrassv2Deployment#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#failure_type Greengrassv2Deployment#failure_type}.
	FailureType *string `field:"optional" json:"failureType" yaml:"failureType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#min_number_of_executed_things Greengrassv2Deployment#min_number_of_executed_things}.
	MinNumberOfExecutedThings *float64 `field:"optional" json:"minNumberOfExecutedThings" yaml:"minNumberOfExecutedThings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#threshold_percentage Greengrassv2Deployment#threshold_percentage}.
	ThresholdPercentage *float64 `field:"optional" json:"thresholdPercentage" yaml:"thresholdPercentage"`
}

