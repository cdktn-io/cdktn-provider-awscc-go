// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticeservicenetworkresourceassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VpclatticeServiceNetworkResourceAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/vpclattice_service_network_resource_association#private_dns_enabled VpclatticeServiceNetworkResourceAssociation#private_dns_enabled}.
	PrivateDnsEnabled interface{} `field:"optional" json:"privateDnsEnabled" yaml:"privateDnsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/vpclattice_service_network_resource_association#resource_configuration_id VpclatticeServiceNetworkResourceAssociation#resource_configuration_id}.
	ResourceConfigurationId *string `field:"optional" json:"resourceConfigurationId" yaml:"resourceConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/vpclattice_service_network_resource_association#service_network_id VpclatticeServiceNetworkResourceAssociation#service_network_id}.
	ServiceNetworkId *string `field:"optional" json:"serviceNetworkId" yaml:"serviceNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/vpclattice_service_network_resource_association#tags VpclatticeServiceNetworkResourceAssociation#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

