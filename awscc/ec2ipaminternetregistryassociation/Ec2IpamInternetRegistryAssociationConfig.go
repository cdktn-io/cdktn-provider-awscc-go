// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipaminternetregistryassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2IpamInternetRegistryAssociationConfig struct {
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
	// The ID of the IPAM to associate with the internet registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ec2_ipam_internet_registry_association#ipam_id Ec2IpamInternetRegistryAssociation#ipam_id}
	IpamId *string `field:"required" json:"ipamId" yaml:"ipamId"`
	// The organization handle at the internet registry (for example, a RIPE NCC organization ID or ARIN Org ID).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ec2_ipam_internet_registry_association#organization_handle Ec2IpamInternetRegistryAssociation#organization_handle}
	OrganizationHandle *string `field:"required" json:"organizationHandle" yaml:"organizationHandle"`
	// The Regional Internet Registry to associate with.
	//
	// The service returns these values in upper case, so the enum uses the round-tripping form.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ec2_ipam_internet_registry_association#rir Ec2IpamInternetRegistryAssociation#rir}
	Rir *string `field:"required" json:"rir" yaml:"rir"`
	// A description for the internet registry association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ec2_ipam_internet_registry_association#description Ec2IpamInternetRegistryAssociation#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags assigned to the internet registry association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/ec2_ipam_internet_registry_association#tags Ec2IpamInternetRegistryAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

