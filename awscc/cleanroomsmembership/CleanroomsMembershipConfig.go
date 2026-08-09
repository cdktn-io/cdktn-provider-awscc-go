// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomsmembership

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsMembershipConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_membership#collaboration_identifier CleanroomsMembership#collaboration_identifier}.
	CollaborationIdentifier *string `field:"required" json:"collaborationIdentifier" yaml:"collaborationIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_membership#query_log_status CleanroomsMembership#query_log_status}.
	QueryLogStatus *string `field:"required" json:"queryLogStatus" yaml:"queryLogStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_membership#default_job_result_configuration CleanroomsMembership#default_job_result_configuration}.
	DefaultJobResultConfiguration *CleanroomsMembershipDefaultJobResultConfiguration `field:"optional" json:"defaultJobResultConfiguration" yaml:"defaultJobResultConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_membership#default_result_configuration CleanroomsMembership#default_result_configuration}.
	DefaultResultConfiguration *CleanroomsMembershipDefaultResultConfiguration `field:"optional" json:"defaultResultConfiguration" yaml:"defaultResultConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_membership#is_metrics_enabled CleanroomsMembership#is_metrics_enabled}.
	IsMetricsEnabled interface{} `field:"optional" json:"isMetricsEnabled" yaml:"isMetricsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_membership#job_log_status CleanroomsMembership#job_log_status}.
	JobLogStatus *string `field:"optional" json:"jobLogStatus" yaml:"jobLogStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_membership#payment_configuration CleanroomsMembership#payment_configuration}.
	PaymentConfiguration *CleanroomsMembershipPaymentConfiguration `field:"optional" json:"paymentConfiguration" yaml:"paymentConfiguration"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/cleanrooms_membership#tags CleanroomsMembership#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

