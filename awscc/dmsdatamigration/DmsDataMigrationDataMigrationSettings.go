// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsdatamigration


type DmsDataMigrationDataMigrationSettings struct {
	// The property specifies whether to enable the CloudWatch log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#cloudwatch_logs_enabled DmsDataMigration#cloudwatch_logs_enabled}
	CloudwatchLogsEnabled interface{} `field:"optional" json:"cloudwatchLogsEnabled" yaml:"cloudwatchLogsEnabled"`
	// The number of parallel jobs that trigger parallel threads to unload the tables from the source, and then load them to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#number_of_jobs DmsDataMigration#number_of_jobs}
	NumberOfJobs *float64 `field:"optional" json:"numberOfJobs" yaml:"numberOfJobs"`
	// The property specifies the rules of selecting objects for data migration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dms_data_migration#selection_rules DmsDataMigration#selection_rules}
	SelectionRules *string `field:"optional" json:"selectionRules" yaml:"selectionRules"`
}

