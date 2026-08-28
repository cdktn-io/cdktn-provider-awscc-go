// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package aiopsinvestigationgroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AiopsInvestigationGroupConfig struct {
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
	// User friendly name for resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#name AiopsInvestigationGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of key-value pairs of notification channels to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#chatbot_notification_channels AiopsInvestigationGroup#chatbot_notification_channels}
	ChatbotNotificationChannels interface{} `field:"optional" json:"chatbotNotificationChannels" yaml:"chatbotNotificationChannels"`
	// An array of cross account configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#cross_account_configurations AiopsInvestigationGroup#cross_account_configurations}
	CrossAccountConfigurations interface{} `field:"optional" json:"crossAccountConfigurations" yaml:"crossAccountConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#encryption_config AiopsInvestigationGroup#encryption_config}.
	EncryptionConfig *AiopsInvestigationGroupEncryptionConfig `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Investigation Group policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#investigation_group_policy AiopsInvestigationGroup#investigation_group_policy}
	InvestigationGroupPolicy *string `field:"optional" json:"investigationGroupPolicy" yaml:"investigationGroupPolicy"`
	// Flag to enable cloud trail history.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#is_cloud_trail_event_history_enabled AiopsInvestigationGroup#is_cloud_trail_event_history_enabled}
	IsCloudTrailEventHistoryEnabled interface{} `field:"optional" json:"isCloudTrailEventHistoryEnabled" yaml:"isCloudTrailEventHistoryEnabled"`
	// The number of days to retain the investigation group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#retention_in_days AiopsInvestigationGroup#retention_in_days}
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// The Investigation Role's ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#role_arn AiopsInvestigationGroup#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#tag_key_boundaries AiopsInvestigationGroup#tag_key_boundaries}.
	TagKeyBoundaries *[]*string `field:"optional" json:"tagKeyBoundaries" yaml:"tagKeyBoundaries"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/aiops_investigation_group#tags AiopsInvestigationGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

