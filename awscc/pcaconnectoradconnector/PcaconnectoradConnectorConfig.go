// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcaconnectoradConnectorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_connector#certificate_authority_arn PcaconnectoradConnector#certificate_authority_arn}.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_connector#directory_id PcaconnectoradConnector#directory_id}.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_connector#vpc_information PcaconnectoradConnector#vpc_information}.
	VpcInformation *PcaconnectoradConnectorVpcInformation `field:"required" json:"vpcInformation" yaml:"vpcInformation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcaconnectorad_connector#tags PcaconnectoradConnector#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

