// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkforce


type SagemakerWorkforceSourceIpConfig struct {
	// A list of one to ten Classless Inter-Domain Routing (CIDR) values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_workforce#cidrs SagemakerWorkforce#cidrs}
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
}

