// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetrigger


type GlueTriggerActions struct {
	// The job arguments used when this trigger fires.
	//
	// For this job run, they replace the default arguments set in the job definition itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#arguments GlueTrigger#arguments}
	Arguments *string `field:"optional" json:"arguments" yaml:"arguments"`
	// The name of the crawler to be used with this action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#crawler_name GlueTrigger#crawler_name}
	CrawlerName *string `field:"optional" json:"crawlerName" yaml:"crawlerName"`
	// The name of a job to be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#job_name GlueTrigger#job_name}
	JobName *string `field:"optional" json:"jobName" yaml:"jobName"`
	// Specifies configuration properties of a job run notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#notification_property GlueTrigger#notification_property}
	NotificationProperty *GlueTriggerActionsNotificationProperty `field:"optional" json:"notificationProperty" yaml:"notificationProperty"`
	// The name of the SecurityConfiguration structure to be used with this action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#security_configuration GlueTrigger#security_configuration}
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The JobRun timeout in minutes.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status. The default is 2,880 minutes (48 hours). This overrides the timeout value set in the parent job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_trigger#timeout GlueTrigger#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

