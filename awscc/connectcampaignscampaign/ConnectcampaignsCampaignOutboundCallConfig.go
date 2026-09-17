// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignscampaign


type ConnectcampaignsCampaignOutboundCallConfig struct {
	// The identifier of the contact flow for the outbound call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaigns_campaign#connect_contact_flow_arn ConnectcampaignsCampaign#connect_contact_flow_arn}
	ConnectContactFlowArn *string `field:"required" json:"connectContactFlowArn" yaml:"connectContactFlowArn"`
	// The configuration used for answering machine detection during outbound calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaigns_campaign#answer_machine_detection_config ConnectcampaignsCampaign#answer_machine_detection_config}
	AnswerMachineDetectionConfig *ConnectcampaignsCampaignOutboundCallConfigAnswerMachineDetectionConfig `field:"optional" json:"answerMachineDetectionConfig" yaml:"answerMachineDetectionConfig"`
	// The queue for the call.
	//
	// If you specify a queue, the phone displayed for caller ID is the phone number specified in the queue. If you do not specify a queue, the queue defined in the contact flow is used. If you do not specify a queue, you must specify a source phone number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaigns_campaign#connect_queue_arn ConnectcampaignsCampaign#connect_queue_arn}
	ConnectQueueArn *string `field:"optional" json:"connectQueueArn" yaml:"connectQueueArn"`
	// The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/connectcampaigns_campaign#connect_source_phone_number ConnectcampaignsCampaign#connect_source_phone_number}
	ConnectSourcePhoneNumber *string `field:"optional" json:"connectSourcePhoneNumber" yaml:"connectSourcePhoneNumber"`
}

