// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmsapplicationslist


type FmsApplicationsListAppsListStruct struct {
	// The application's name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_applications_list#app_name FmsApplicationsList#app_name}
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// The application's port number, for example 80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_applications_list#port FmsApplicationsList#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The IP protocol name or number.
	//
	// The name can be one of tcp, udp, or icmp. For information on possible numbers, see Protocol Numbers (https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fms_applications_list#protocol FmsApplicationsList#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

