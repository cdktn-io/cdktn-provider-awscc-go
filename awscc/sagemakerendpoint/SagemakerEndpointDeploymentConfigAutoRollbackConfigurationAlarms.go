// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint


type SagemakerEndpointDeploymentConfigAutoRollbackConfigurationAlarms struct {
	// The name of the CloudWatch alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#alarm_name SagemakerEndpoint#alarm_name}
	AlarmName *string `field:"optional" json:"alarmName" yaml:"alarmName"`
}

