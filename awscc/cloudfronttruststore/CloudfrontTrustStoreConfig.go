// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfronttruststore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudfrontTrustStoreConfig struct {
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
	// The trust store's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_trust_store#name CloudfrontTrustStore#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A CA certificates bundle source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_trust_store#ca_certificates_bundle_source CloudfrontTrustStore#ca_certificates_bundle_source}
	CaCertificatesBundleSource *CloudfrontTrustStoreCaCertificatesBundleSource `field:"optional" json:"caCertificatesBundleSource" yaml:"caCertificatesBundleSource"`
	// A complex type that contains zero or more ``Tag`` elements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_trust_store#tags CloudfrontTrustStore#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// A boolean. When true, performs real-time certificate revocation checks by querying the OCSP endpoint specified within the client certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudfront_trust_store#use_client_certificate_ocsp_endpoint CloudfrontTrustStore#use_client_certificate_ocsp_endpoint}
	UseClientCertificateOcspEndpoint interface{} `field:"optional" json:"useClientCertificateOcspEndpoint" yaml:"useClientCertificateOcspEndpoint"`
}

