// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetmetric


type IotFleetMetricAggregationType struct {
	// Fleet Indexing aggregation type names such as Statistics, Percentiles and Cardinality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_fleet_metric#name IotFleetMetric#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Fleet Indexing aggregation type values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_fleet_metric#values IotFleetMetric#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

