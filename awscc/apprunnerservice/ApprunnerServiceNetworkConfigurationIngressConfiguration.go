// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceNetworkConfigurationIngressConfiguration struct {
	// It's set to true if the Apprunner service is publicly accessible. It's set to false otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apprunner_service#is_publicly_accessible ApprunnerService#is_publicly_accessible}
	IsPubliclyAccessible interface{} `field:"optional" json:"isPubliclyAccessible" yaml:"isPubliclyAccessible"`
}

