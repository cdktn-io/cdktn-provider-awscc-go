// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkfirewallcontainerassociation


type NetworkfirewallContainerAssociationContainerMonitoringConfigurationsAttributeFilters struct {
	// The attribute key to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/networkfirewall_container_association#key NetworkfirewallContainerAssociation#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The attribute value to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/networkfirewall_container_association#value NetworkfirewallContainerAssociation#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

