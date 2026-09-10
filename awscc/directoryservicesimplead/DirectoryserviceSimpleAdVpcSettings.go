// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package directoryservicesimplead


type DirectoryserviceSimpleAdVpcSettings struct {
	// The identifiers of the subnets for the directory servers.
	//
	// The two subnets must be in different Availability Zones. AWS Directory Service specifies a directory server and a DNS server in each of these subnets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/directoryservice_simple_ad#subnet_ids DirectoryserviceSimpleAd#subnet_ids}
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The identifier of the VPC in which to create the directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/directoryservice_simple_ad#vpc_id DirectoryserviceSimpleAd#vpc_id}
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

