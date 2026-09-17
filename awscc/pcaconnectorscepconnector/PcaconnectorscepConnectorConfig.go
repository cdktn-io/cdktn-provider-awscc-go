// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectorscepconnector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcaconnectorscepConnectorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pcaconnectorscep_connector#certificate_authority_arn PcaconnectorscepConnector#certificate_authority_arn}.
	CertificateAuthorityArn *string `field:"required" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pcaconnectorscep_connector#mobile_device_management PcaconnectorscepConnector#mobile_device_management}.
	MobileDeviceManagement *PcaconnectorscepConnectorMobileDeviceManagement `field:"optional" json:"mobileDeviceManagement" yaml:"mobileDeviceManagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pcaconnectorscep_connector#tags PcaconnectorscepConnector#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pcaconnectorscep_connector#vpc_endpoint_id PcaconnectorscepConnector#vpc_endpoint_id}.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

