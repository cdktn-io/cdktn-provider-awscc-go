// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailcontainer


type LightsailContainerPublicDomainNames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_container#certificate_name LightsailContainer#certificate_name}.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// An object that describes the configuration for the containers of the deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lightsail_container#domain_names LightsailContainer#domain_names}
	DomainNames *[]*string `field:"optional" json:"domainNames" yaml:"domainNames"`
}

