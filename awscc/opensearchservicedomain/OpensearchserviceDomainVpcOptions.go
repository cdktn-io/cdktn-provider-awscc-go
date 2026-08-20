// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain


type OpensearchserviceDomainVpcOptions struct {
	// Controls whether egress traffic from the domain is routed through the customer VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#egress_enabled OpensearchserviceDomain#egress_enabled}
	EgressEnabled interface{} `field:"optional" json:"egressEnabled" yaml:"egressEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#security_group_ids OpensearchserviceDomain#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/opensearchservice_domain#subnet_ids OpensearchserviceDomain#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

