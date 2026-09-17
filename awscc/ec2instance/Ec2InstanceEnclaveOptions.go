// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceEnclaveOptions struct {
	// If this parameter is set to true, the instance is enabled for AWS Nitro Enclaves;
	//
	// otherwise, it is not enabled for AWS Nitro Enclaves.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance#enabled Ec2Instance#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

