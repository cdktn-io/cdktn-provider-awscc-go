// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcampaignsv2campaign


type Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfig struct {
	// The configuration used for answering machine detection during outbound calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#answer_machine_detection_config Connectcampaignsv2Campaign#answer_machine_detection_config}
	AnswerMachineDetectionConfig *Connectcampaignsv2CampaignChannelSubtypeConfigTelephonyDefaultOutboundConfigAnswerMachineDetectionConfig `field:"optional" json:"answerMachineDetectionConfig" yaml:"answerMachineDetectionConfig"`
	// The identifier of the contact flow for the outbound call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#connect_contact_flow_id Connectcampaignsv2Campaign#connect_contact_flow_id}
	ConnectContactFlowId *string `field:"optional" json:"connectContactFlowId" yaml:"connectContactFlowId"`
	// The phone number associated with the Amazon Connect instance, in E.164 format. If you do not specify a source phone number, you must specify a queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#connect_source_phone_number Connectcampaignsv2Campaign#connect_source_phone_number}
	ConnectSourcePhoneNumber *string `field:"optional" json:"connectSourcePhoneNumber" yaml:"connectSourcePhoneNumber"`
	// Maximum ring time for outbound calls in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connectcampaignsv2_campaign#ring_timeout Connectcampaignsv2Campaign#ring_timeout}
	RingTimeout *float64 `field:"optional" json:"ringTimeout" yaml:"ringTimeout"`
}

