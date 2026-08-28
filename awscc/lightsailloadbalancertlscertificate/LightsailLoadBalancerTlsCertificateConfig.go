// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailloadbalancertlscertificate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailLoadBalancerTlsCertificateConfig struct {
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
	// The domain name (e.g., example.com ) for your SSL/TLS certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_load_balancer_tls_certificate#certificate_domain_name LightsailLoadBalancerTlsCertificate#certificate_domain_name}
	CertificateDomainName *string `field:"required" json:"certificateDomainName" yaml:"certificateDomainName"`
	// The SSL/TLS certificate name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_load_balancer_tls_certificate#certificate_name LightsailLoadBalancerTlsCertificate#certificate_name}
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// The name of your load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_load_balancer_tls_certificate#load_balancer_name LightsailLoadBalancerTlsCertificate#load_balancer_name}
	LoadBalancerName *string `field:"required" json:"loadBalancerName" yaml:"loadBalancerName"`
	// An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_load_balancer_tls_certificate#certificate_alternative_names LightsailLoadBalancerTlsCertificate#certificate_alternative_names}
	CertificateAlternativeNames *[]*string `field:"optional" json:"certificateAlternativeNames" yaml:"certificateAlternativeNames"`
	// A Boolean value that indicates whether HTTPS redirection is enabled for the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_load_balancer_tls_certificate#https_redirection_enabled LightsailLoadBalancerTlsCertificate#https_redirection_enabled}
	HttpsRedirectionEnabled interface{} `field:"optional" json:"httpsRedirectionEnabled" yaml:"httpsRedirectionEnabled"`
	// When true, the SSL/TLS certificate is attached to the Lightsail load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_load_balancer_tls_certificate#is_attached LightsailLoadBalancerTlsCertificate#is_attached}
	IsAttached interface{} `field:"optional" json:"isAttached" yaml:"isAttached"`
}

