// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanagercertificate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CertificatemanagerCertificateConfig struct {
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
	// The fully qualified domain name (FQDN), such as www.example.com, with which you want to secure an ACM certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#domain_name CertificatemanagerCertificate#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be used to issue the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#certificate_authority_arn CertificatemanagerCertificate#certificate_authority_arn}
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Specifies whether the certificate can be exported. ENABLED allows the certificate to be exported, DISABLED prevents export.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#certificate_export CertificatemanagerCertificate#certificate_export}
	CertificateExport *string `field:"optional" json:"certificateExport" yaml:"certificateExport"`
	// You can opt out of certificate transparency logging by specifying the DISABLED option. Opt in by specifying ENABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#certificate_transparency_logging_preference CertificatemanagerCertificate#certificate_transparency_logging_preference}
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// Domain information that domain name registrars use to verify your identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#domain_validation_options CertificatemanagerCertificate#domain_validation_options}
	DomainValidationOptions interface{} `field:"optional" json:"domainValidationOptions" yaml:"domainValidationOptions"`
	// Specifies the algorithm of the public and private key pair that your certificate uses to encrypt data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#key_algorithm CertificatemanagerCertificate#key_algorithm}
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Additional FQDNs to be included in the Subject Alternative Name extension of the ACM certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#subject_alternative_names CertificatemanagerCertificate#subject_alternative_names}
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// Key-value pairs that can identify the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#tags CertificatemanagerCertificate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The method you want to use to validate that you own or control the domain associated with a public certificate.
	//
	// Valid values are DNS, EMAIL or HTTP
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/certificatemanager_certificate#validation_method CertificatemanagerCertificate#validation_method}
	ValidationMethod *string `field:"optional" json:"validationMethod" yaml:"validationMethod"`
}

