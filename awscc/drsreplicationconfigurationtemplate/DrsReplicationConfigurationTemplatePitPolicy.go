// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package drsreplicationconfigurationtemplate


type DrsReplicationConfigurationTemplatePitPolicy struct {
	// How often, in the chosen units, a snapshot should be taken.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_replication_configuration_template#interval DrsReplicationConfigurationTemplate#interval}
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// The duration to retain a snapshot for, in the chosen units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_replication_configuration_template#retention_duration DrsReplicationConfigurationTemplate#retention_duration}
	RetentionDuration *float64 `field:"required" json:"retentionDuration" yaml:"retentionDuration"`
	// The units used to measure the interval and retentionDuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_replication_configuration_template#units DrsReplicationConfigurationTemplate#units}
	Units *string `field:"required" json:"units" yaml:"units"`
	// Whether this rule is enabled or not.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_replication_configuration_template#enabled DrsReplicationConfigurationTemplate#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ID of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/drs_replication_configuration_template#rule_id DrsReplicationConfigurationTemplate#rule_id}
	RuleId *float64 `field:"optional" json:"ruleId" yaml:"ruleId"`
}

