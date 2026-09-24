// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnectorv2


type SecurityhubConnectorV2ProviderNameJiraCloud struct {
	// The project key for a Jira Cloud instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_connector_v2#project_key SecurityhubConnectorV2#project_key}
	ProjectKey *string `field:"optional" json:"projectKey" yaml:"projectKey"`
}

