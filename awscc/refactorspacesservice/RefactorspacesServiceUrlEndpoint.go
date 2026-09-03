// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package refactorspacesservice


type RefactorspacesServiceUrlEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/refactorspaces_service#health_url RefactorspacesService#health_url}.
	HealthUrl *string `field:"optional" json:"healthUrl" yaml:"healthUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/refactorspaces_service#url RefactorspacesService#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}

