// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectmetric


type ConnectMetricMetricCalculationCalculationComponentsMetricFiltersStringCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_metric#comparison ConnectMetric#comparison}.
	Comparison *string `field:"optional" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_metric#values ConnectMetric#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

