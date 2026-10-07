// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesemailidentitycertificate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesEmailIdentityCertificateConfig struct {
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
	// The ARN of the AWS Certificate Manager certificate to associate with the sender.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/ses_email_identity_certificate#certificate_arn SesEmailIdentityCertificate#certificate_arn}
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// The email identity that owns the sender the certificate is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/ses_email_identity_certificate#email_identity SesEmailIdentityCertificate#email_identity}
	EmailIdentity *string `field:"required" json:"emailIdentity" yaml:"emailIdentity"`
	// The sender the certificate signs for. For an email address identity this is the identity itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/ses_email_identity_certificate#from_address SesEmailIdentityCertificate#from_address}
	FromAddress *string `field:"required" json:"fromAddress" yaml:"fromAddress"`
}

