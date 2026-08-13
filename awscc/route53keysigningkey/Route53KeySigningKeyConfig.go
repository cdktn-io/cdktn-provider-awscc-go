// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53keysigningkey

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53KeySigningKeyConfig struct {
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
	// The unique string (ID) used to identify a hosted zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/route53_key_signing_key#hosted_zone_id Route53KeySigningKey#hosted_zone_id}
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// The Amazon resource name (ARN) for a customer managed key (CMK) in AWS Key Management Service (KMS).
	//
	// The KeyManagementServiceArn must be unique for each key signing key (KSK) in a single hosted zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/route53_key_signing_key#key_management_service_arn Route53KeySigningKey#key_management_service_arn}
	KeyManagementServiceArn *string `field:"required" json:"keyManagementServiceArn" yaml:"keyManagementServiceArn"`
	// An alphanumeric string used to identify a key signing key (KSK).
	//
	// Name must be unique for each key signing key in the same hosted zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/route53_key_signing_key#name Route53KeySigningKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A string specifying the initial status of the key signing key (KSK).
	//
	// You can set the value to ACTIVE or INACTIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/route53_key_signing_key#status Route53KeySigningKey#status}
	Status *string `field:"required" json:"status" yaml:"status"`
}

