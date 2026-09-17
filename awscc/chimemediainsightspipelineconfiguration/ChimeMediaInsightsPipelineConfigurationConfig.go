// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimemediainsightspipelineconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeMediaInsightsPipelineConfigurationConfig struct {
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
	// The elements in the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#elements ChimeMediaInsightsPipelineConfiguration#elements}
	Elements interface{} `field:"required" json:"elements" yaml:"elements"`
	// The name of the media insights pipeline configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#media_insights_pipeline_configuration_name ChimeMediaInsightsPipelineConfiguration#media_insights_pipeline_configuration_name}
	MediaInsightsPipelineConfigurationName *string `field:"required" json:"mediaInsightsPipelineConfigurationName" yaml:"mediaInsightsPipelineConfigurationName"`
	// The ARN of the role used by the service to access Amazon Web Services resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#resource_access_role_arn ChimeMediaInsightsPipelineConfiguration#resource_access_role_arn}
	ResourceAccessRoleArn *string `field:"required" json:"resourceAccessRoleArn" yaml:"resourceAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#real_time_alert_configuration ChimeMediaInsightsPipelineConfiguration#real_time_alert_configuration}.
	RealTimeAlertConfiguration *ChimeMediaInsightsPipelineConfigurationRealTimeAlertConfiguration `field:"optional" json:"realTimeAlertConfiguration" yaml:"realTimeAlertConfiguration"`
	// The tags associated with the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_media_insights_pipeline_configuration#tags ChimeMediaInsightsPipelineConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

