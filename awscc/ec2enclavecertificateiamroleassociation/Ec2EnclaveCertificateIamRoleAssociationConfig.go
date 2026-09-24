// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2enclavecertificateiamroleassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2EnclaveCertificateIamRoleAssociationConfig struct {
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
	// The Amazon Resource Name (ARN) of the ACM certificate with which to associate the IAM role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_enclave_certificate_iam_role_association#certificate_arn Ec2EnclaveCertificateIamRoleAssociation#certificate_arn}
	CertificateArn *string `field:"required" json:"certificateArn" yaml:"certificateArn"`
	// The Amazon Resource Name (ARN) of the IAM role to associate with the ACM certificate.
	//
	// You can associate up to 16 IAM roles with an ACM certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_enclave_certificate_iam_role_association#role_arn Ec2EnclaveCertificateIamRoleAssociation#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}

