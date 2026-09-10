// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedomain

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneDomainConfig struct {
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
	// The name of the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_domain#name DatazoneDomain#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_domain#description DatazoneDomain#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The domain execution role that is created when an Amazon DataZone domain is created.
	//
	// The domain execution role is created in the AWS account that houses the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_domain#domain_execution_role DatazoneDomain#domain_execution_role}
	DomainExecutionRole *string `field:"optional" json:"domainExecutionRole" yaml:"domainExecutionRole"`
	// The version of the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_domain#domain_version DatazoneDomain#domain_version}
	DomainVersion *string `field:"optional" json:"domainVersion" yaml:"domainVersion"`
	// The identifier of the AWS Key Management Service (KMS) key that is used to encrypt the Amazon DataZone domain, metadata, and reporting data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_domain#kms_key_identifier DatazoneDomain#kms_key_identifier}
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// The service role of the domain that is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_domain#service_role DatazoneDomain#service_role}
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// The single-sign on configuration of the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_domain#single_sign_on DatazoneDomain#single_sign_on}
	SingleSignOn *DatazoneDomainSingleSignOn `field:"optional" json:"singleSignOn" yaml:"singleSignOn"`
	// The tags specified for the Amazon DataZone domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_domain#tags DatazoneDomain#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

