// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoiceregistrationattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SmsvoiceRegistrationAttachmentConfig struct {
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
	// The registration file to upload.
	//
	// The maximum file size is 1500KB and valid file extensions are PDF, JPEG and PNG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/smsvoice_registration_attachment#attachment_body SmsvoiceRegistrationAttachment#attachment_body}
	AttachmentBody *string `field:"optional" json:"attachmentBody" yaml:"attachmentBody"`
	// A URL to the required registration file. For example, the URL to an MMS/shortcode form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/smsvoice_registration_attachment#attachment_url SmsvoiceRegistrationAttachment#attachment_url}
	AttachmentUrl *string `field:"optional" json:"attachmentUrl" yaml:"attachmentUrl"`
	// An array of tags (key and value pairs) to associate with the registration attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/smsvoice_registration_attachment#tags SmsvoiceRegistrationAttachment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

