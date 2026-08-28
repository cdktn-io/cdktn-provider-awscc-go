// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlineworker


type DeadlineWorkerHostProperties struct {
	// The host name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_worker#host_name DeadlineWorker#host_name}
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// The IP addresses for a host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/deadline_worker#ip_addresses DeadlineWorker#ip_addresses}
	IpAddresses *DeadlineWorkerHostPropertiesIpAddresses `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
}

