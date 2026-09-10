// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package b2bipartnership


type B2BiPartnershipCapabilityOptionsOutboundEdiX12CommonInterchangeControlHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/b2bi_partnership#acknowledgment_requested_code B2BiPartnership#acknowledgment_requested_code}.
	AcknowledgmentRequestedCode *string `field:"optional" json:"acknowledgmentRequestedCode" yaml:"acknowledgmentRequestedCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/b2bi_partnership#receiver_id B2BiPartnership#receiver_id}.
	ReceiverId *string `field:"optional" json:"receiverId" yaml:"receiverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/b2bi_partnership#receiver_id_qualifier B2BiPartnership#receiver_id_qualifier}.
	ReceiverIdQualifier *string `field:"optional" json:"receiverIdQualifier" yaml:"receiverIdQualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/b2bi_partnership#repetition_separator B2BiPartnership#repetition_separator}.
	RepetitionSeparator *string `field:"optional" json:"repetitionSeparator" yaml:"repetitionSeparator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/b2bi_partnership#sender_id B2BiPartnership#sender_id}.
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/b2bi_partnership#sender_id_qualifier B2BiPartnership#sender_id_qualifier}.
	SenderIdQualifier *string `field:"optional" json:"senderIdQualifier" yaml:"senderIdQualifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/b2bi_partnership#usage_indicator_code B2BiPartnership#usage_indicator_code}.
	UsageIndicatorCode *string `field:"optional" json:"usageIndicatorCode" yaml:"usageIndicatorCode"`
}

