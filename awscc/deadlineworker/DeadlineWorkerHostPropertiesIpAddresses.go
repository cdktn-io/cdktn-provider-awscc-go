// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlineworker


type DeadlineWorkerHostPropertiesIpAddresses struct {
	// The IpV4 address of the network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/deadline_worker#ip_v4_addresses DeadlineWorker#ip_v4_addresses}
	IpV4Addresses *[]*string `field:"optional" json:"ipV4Addresses" yaml:"ipV4Addresses"`
	// The IpV6 address for the network and node component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/deadline_worker#ip_v6_addresses DeadlineWorker#ip_v6_addresses}
	IpV6Addresses *[]*string `field:"optional" json:"ipV6Addresses" yaml:"ipV6Addresses"`
}

