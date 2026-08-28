// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectmetric


type ConnectMetricMetricCalculation struct {
	// The calculation formula.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#calculation ConnectMetric#calculation}
	Calculation *string `field:"optional" json:"calculation" yaml:"calculation"`
	// The calculation components for the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_metric#calculation_components ConnectMetric#calculation_components}
	CalculationComponents interface{} `field:"optional" json:"calculationComponents" yaml:"calculationComponents"`
}

