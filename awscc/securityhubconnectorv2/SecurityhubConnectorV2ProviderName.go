// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubconnectorv2


type SecurityhubConnectorV2ProviderName struct {
	// The configuration settings required to establish an integration between AWS Security Hub and Azure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_connector_v2#azure SecurityhubConnectorV2#azure}
	Azure *SecurityhubConnectorV2ProviderNameAzure `field:"optional" json:"azure" yaml:"azure"`
	// The initial configuration settings required to establish an integration between Security Hub and Jira Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_connector_v2#jira_cloud SecurityhubConnectorV2#jira_cloud}
	JiraCloud *SecurityhubConnectorV2ProviderNameJiraCloud `field:"optional" json:"jiraCloud" yaml:"jiraCloud"`
	// The initial configuration settings required to establish an integration between Security Hub and ServiceNow ITSM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/securityhub_connector_v2#service_now SecurityhubConnectorV2#service_now}
	ServiceNow *SecurityhubConnectorV2ProviderNameServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
}

