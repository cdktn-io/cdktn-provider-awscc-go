// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cleanroomscollaboration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CleanroomsCollaborationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#creator_display_name CleanroomsCollaboration#creator_display_name}.
	CreatorDisplayName *string `field:"required" json:"creatorDisplayName" yaml:"creatorDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#description CleanroomsCollaboration#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#name CleanroomsCollaboration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#query_log_status CleanroomsCollaboration#query_log_status}.
	QueryLogStatus *string `field:"required" json:"queryLogStatus" yaml:"queryLogStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#allowed_result_regions CleanroomsCollaboration#allowed_result_regions}.
	AllowedResultRegions *[]*string `field:"optional" json:"allowedResultRegions" yaml:"allowedResultRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#analytics_engine CleanroomsCollaboration#analytics_engine}.
	AnalyticsEngine *string `field:"optional" json:"analyticsEngine" yaml:"analyticsEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#auto_approved_change_types CleanroomsCollaboration#auto_approved_change_types}.
	AutoApprovedChangeTypes *[]*string `field:"optional" json:"autoApprovedChangeTypes" yaml:"autoApprovedChangeTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#creator_member_abilities CleanroomsCollaboration#creator_member_abilities}.
	CreatorMemberAbilities *[]*string `field:"optional" json:"creatorMemberAbilities" yaml:"creatorMemberAbilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#creator_ml_member_abilities CleanroomsCollaboration#creator_ml_member_abilities}.
	CreatorMlMemberAbilities *CleanroomsCollaborationCreatorMlMemberAbilities `field:"optional" json:"creatorMlMemberAbilities" yaml:"creatorMlMemberAbilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#creator_payment_configuration CleanroomsCollaboration#creator_payment_configuration}.
	CreatorPaymentConfiguration *CleanroomsCollaborationCreatorPaymentConfiguration `field:"optional" json:"creatorPaymentConfiguration" yaml:"creatorPaymentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#data_encryption_metadata CleanroomsCollaboration#data_encryption_metadata}.
	DataEncryptionMetadata *CleanroomsCollaborationDataEncryptionMetadata `field:"optional" json:"dataEncryptionMetadata" yaml:"dataEncryptionMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#is_metrics_enabled CleanroomsCollaboration#is_metrics_enabled}.
	IsMetricsEnabled interface{} `field:"optional" json:"isMetricsEnabled" yaml:"isMetricsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#job_log_status CleanroomsCollaboration#job_log_status}.
	JobLogStatus *string `field:"optional" json:"jobLogStatus" yaml:"jobLogStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#members CleanroomsCollaboration#members}.
	Members interface{} `field:"optional" json:"members" yaml:"members"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cleanrooms_collaboration#tags CleanroomsCollaboration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

