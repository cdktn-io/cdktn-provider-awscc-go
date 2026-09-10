// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetrigger


type GlueTriggerPredicateConditions struct {
	// The name of the crawler to which this condition applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#crawler_name GlueTrigger#crawler_name}
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// The state of the crawler to which this condition applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#crawl_state GlueTrigger#crawl_state}
	CrawlState *string `field:"optional" json:"crawlState" yaml:"crawlState"`
	// The name of the job whose JobRuns this condition applies to, and on which this trigger waits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#job_name GlueTrigger#job_name}
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// A logical operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#logical_operator GlueTrigger#logical_operator}
	LogicalOperator *string `field:"optional" json:"logicalOperator" yaml:"logicalOperator"`
	// The condition state. Currently, the values supported are SUCCEEDED, STOPPED, TIMEOUT, and FAILED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#state GlueTrigger#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

