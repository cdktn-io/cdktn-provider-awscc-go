// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iamservercertificate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IamServerCertificateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_server_certificate#certificate_body IamServerCertificate#certificate_body}.
	CertificateBody *string `field:"optional" json:"certificateBody" yaml:"certificateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_server_certificate#certificate_chain IamServerCertificate#certificate_chain}.
	CertificateChain *string `field:"optional" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_server_certificate#path IamServerCertificate#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_server_certificate#private_key IamServerCertificate#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_server_certificate#server_certificate_name IamServerCertificate#server_certificate_name}.
	ServerCertificateName *string `field:"optional" json:"serverCertificateName" yaml:"serverCertificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iam_server_certificate#tags IamServerCertificate#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

