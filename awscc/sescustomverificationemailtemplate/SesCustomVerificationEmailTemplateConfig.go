// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sescustomverificationemailtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesCustomVerificationEmailTemplateConfig struct {
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
	// The URL that the recipient of the verification email is sent to if his or her address is not successfully verified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_custom_verification_email_template#failure_redirection_url SesCustomVerificationEmailTemplate#failure_redirection_url}
	FailureRedirectionUrl *string `field:"required" json:"failureRedirectionUrl" yaml:"failureRedirectionUrl"`
	// The email address that the custom verification email is sent from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_custom_verification_email_template#from_email_address SesCustomVerificationEmailTemplate#from_email_address}
	FromEmailAddress *string `field:"required" json:"fromEmailAddress" yaml:"fromEmailAddress"`
	// The URL that the recipient of the verification email is sent to if his or her address is successfully verified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_custom_verification_email_template#success_redirection_url SesCustomVerificationEmailTemplate#success_redirection_url}
	SuccessRedirectionUrl *string `field:"required" json:"successRedirectionUrl" yaml:"successRedirectionUrl"`
	// The content of the custom verification email.
	//
	// The total size of the email must be less than 10 MB. The message body may contain HTML, with some limitations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_custom_verification_email_template#template_content SesCustomVerificationEmailTemplate#template_content}
	TemplateContent *string `field:"required" json:"templateContent" yaml:"templateContent"`
	// The name of the custom verification email template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_custom_verification_email_template#template_name SesCustomVerificationEmailTemplate#template_name}
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// The subject line of the custom verification email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_custom_verification_email_template#template_subject SesCustomVerificationEmailTemplate#template_subject}
	TemplateSubject *string `field:"required" json:"templateSubject" yaml:"templateSubject"`
	// The tags (keys and values) associated with the tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_custom_verification_email_template#tags SesCustomVerificationEmailTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

