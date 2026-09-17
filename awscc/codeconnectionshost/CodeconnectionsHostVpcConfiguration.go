// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codeconnectionshost


type CodeconnectionsHostVpcConfiguration struct {
	// The ID of the security group or security groups associated with the Amazon VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codeconnections_host#security_group_ids CodeconnectionsHost#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The ID of the subnet or subnets associated with the Amazon VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codeconnections_host#subnet_ids CodeconnectionsHost#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The value of the Transport Layer Security (TLS) certificate associated with the infrastructure where your provider type is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codeconnections_host#tls_certificate CodeconnectionsHost#tls_certificate}
	TlsCertificate *string `field:"optional" json:"tlsCertificate" yaml:"tlsCertificate"`
	// The ID of the Amazon VPC connected to the infrastructure where your provider type is installed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/codeconnections_host#vpc_id CodeconnectionsHost#vpc_id}
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

