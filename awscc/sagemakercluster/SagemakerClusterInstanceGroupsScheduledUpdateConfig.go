// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsScheduledUpdateConfig struct {
	// The configuration to use when updating the AMI versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#deployment_config SagemakerCluster#deployment_config}
	DeploymentConfig *SagemakerClusterInstanceGroupsScheduledUpdateConfigDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// A cron expression that specifies the schedule that SageMaker follows when updating the AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#schedule_expression SagemakerCluster#schedule_expression}
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

