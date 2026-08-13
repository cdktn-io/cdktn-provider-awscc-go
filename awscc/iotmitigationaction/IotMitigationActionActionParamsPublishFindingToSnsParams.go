// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotmitigationaction


type IotMitigationActionActionParamsPublishFindingToSnsParams struct {
	// The ARN of the topic to which you want to publish the findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iot_mitigation_action#topic_arn IotMitigationAction#topic_arn}
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

