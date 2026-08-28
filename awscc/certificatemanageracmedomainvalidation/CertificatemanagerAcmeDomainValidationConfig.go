// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package certificatemanageracmedomainvalidation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CertificatemanagerAcmeDomainValidationConfig struct {
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
	// The ARN of the ACME endpoint this domain validation is associated with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/certificatemanager_acme_domain_validation#acme_endpoint_arn CertificatemanagerAcmeDomainValidation#acme_endpoint_arn}
	AcmeEndpointArn *string `field:"required" json:"acmeEndpointArn" yaml:"acmeEndpointArn"`
	// The domain name to validate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/certificatemanager_acme_domain_validation#domain_name CertificatemanagerAcmeDomainValidation#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Prevalidation method configuration. Currently only DNS-based prevalidation is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/certificatemanager_acme_domain_validation#prevalidation_options CertificatemanagerAcmeDomainValidation#prevalidation_options}
	PrevalidationOptions *CertificatemanagerAcmeDomainValidationPrevalidationOptions `field:"required" json:"prevalidationOptions" yaml:"prevalidationOptions"`
	// Tags associated with the domain validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/certificatemanager_acme_domain_validation#tags CertificatemanagerAcmeDomainValidation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

