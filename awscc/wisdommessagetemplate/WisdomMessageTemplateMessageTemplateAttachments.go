// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdommessagetemplate


type WisdomMessageTemplateMessageTemplateAttachments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_message_template#attachment_id WisdomMessageTemplate#attachment_id}.
	AttachmentId *string `field:"optional" json:"attachmentId" yaml:"attachmentId"`
	// The name of the attachment file being uploaded. The name should include the file extension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_message_template#attachment_name WisdomMessageTemplate#attachment_name}
	AttachmentName *string `field:"optional" json:"attachmentName" yaml:"attachmentName"`
	// The S3 Presigned URL for the attachment file.
	//
	// When generating the PreSignedUrl, please ensure that the expires-in time is set to 30 minutes. The URL can be generated through the AWS Console or through the AWS CLI (https://docs.aws.amazon.com/AmazonS3/latest/userguide/ShareObjectPreSignedURL.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/wisdom_message_template#s3_presigned_url WisdomMessageTemplate#s3_presigned_url}
	S3PresignedUrl *string `field:"optional" json:"s3PresignedUrl" yaml:"s3PresignedUrl"`
}

