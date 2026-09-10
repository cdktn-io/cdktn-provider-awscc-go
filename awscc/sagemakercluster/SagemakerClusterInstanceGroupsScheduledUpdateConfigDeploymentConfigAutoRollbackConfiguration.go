// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsScheduledUpdateConfigDeploymentConfigAutoRollbackConfiguration struct {
	// The name of the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#alarm_name SagemakerCluster#alarm_name}
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
}

