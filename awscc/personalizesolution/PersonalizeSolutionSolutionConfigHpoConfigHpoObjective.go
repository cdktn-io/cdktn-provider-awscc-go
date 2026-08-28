// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizesolution


type PersonalizeSolutionSolutionConfigHpoConfigHpoObjective struct {
	// The name of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/personalize_solution#metric_name PersonalizeSolution#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A regular expression for finding the metric in the training job logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/personalize_solution#metric_regex PersonalizeSolution#metric_regex}
	MetricRegex *string `field:"optional" json:"metricRegex" yaml:"metricRegex"`
	// The type of the metric. Valid values are Maximize and Minimize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/personalize_solution#type PersonalizeSolution#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

