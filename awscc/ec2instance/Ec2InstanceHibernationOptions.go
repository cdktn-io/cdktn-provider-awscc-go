// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2instance


type Ec2InstanceHibernationOptions struct {
	// If you set this parameter to true, your instance is enabled for hibernation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_instance#configured Ec2Instance#configured}
	Configured interface{} `field:"optional" json:"configured" yaml:"configured"`
}

