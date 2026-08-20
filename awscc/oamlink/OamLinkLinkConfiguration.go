// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package oamlink


type OamLinkLinkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/oam_link#log_group_configuration OamLink#log_group_configuration}.
	LogGroupConfiguration *OamLinkLinkConfigurationLogGroupConfiguration `field:"optional" json:"logGroupConfiguration" yaml:"logGroupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/oam_link#metric_configuration OamLink#metric_configuration}.
	MetricConfiguration *OamLinkLinkConfigurationMetricConfiguration `field:"optional" json:"metricConfiguration" yaml:"metricConfiguration"`
}

