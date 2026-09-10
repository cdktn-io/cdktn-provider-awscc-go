// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamresourcediscoveryassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2IpamResourceDiscoveryAssociationConfig struct {
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
	// The Id of the IPAM this Resource Discovery is associated to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_resource_discovery_association#ipam_id Ec2IpamResourceDiscoveryAssociation#ipam_id}
	IpamId *string `field:"required" json:"ipamId" yaml:"ipamId"`
	// The Amazon Resource Name (ARN) of the IPAM Resource Discovery Association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_resource_discovery_association#ipam_resource_discovery_id Ec2IpamResourceDiscoveryAssociation#ipam_resource_discovery_id}
	IpamResourceDiscoveryId *string `field:"required" json:"ipamResourceDiscoveryId" yaml:"ipamResourceDiscoveryId"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_ipam_resource_discovery_association#tags Ec2IpamResourceDiscoveryAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

