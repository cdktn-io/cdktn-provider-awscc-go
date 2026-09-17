// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotaccountauditconfiguration


type IotAccountAuditConfigurationAuditNotificationTargetConfigurationsSns struct {
	// True if notifications to the target are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_account_audit_configuration#enabled IotAccountAuditConfiguration#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The ARN of the role that grants permission to send notifications to the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_account_audit_configuration#role_arn IotAccountAuditConfiguration#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// The ARN of the target (SNS topic) to which audit notifications are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_account_audit_configuration#target_arn IotAccountAuditConfiguration#target_arn}
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

