// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package refactorspacesapplication


type RefactorspacesApplicationApiGatewayProxy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/refactorspaces_application#endpoint_type RefactorspacesApplication#endpoint_type}.
	EndpointType *string `field:"optional" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/refactorspaces_application#stage_name RefactorspacesApplication#stage_name}.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

