// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsrule


type EventsRuleTargetsRedshiftDataParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_rule#database EventsRule#database}.
	Database *string `field:"optional" json:"database" yaml:"database"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_rule#db_user EventsRule#db_user}.
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_rule#secret_manager_arn EventsRule#secret_manager_arn}.
	SecretManagerArn *string `field:"optional" json:"secretManagerArn" yaml:"secretManagerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_rule#sql EventsRule#sql}.
	Sql *string `field:"optional" json:"sql" yaml:"sql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_rule#sqls EventsRule#sqls}.
	Sqls *[]*string `field:"optional" json:"sqls" yaml:"sqls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_rule#statement_name EventsRule#statement_name}.
	StatementName *string `field:"optional" json:"statementName" yaml:"statementName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/events_rule#with_event EventsRule#with_event}.
	WithEvent interface{} `field:"optional" json:"withEvent" yaml:"withEvent"`
}

