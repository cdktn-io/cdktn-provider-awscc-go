// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsMcpServerNewRelicAuthorizationConfigApiKey struct {
	// New Relic Account ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#account_id DevopsagentService#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// List of alert policy IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#alert_policy_ids DevopsagentService#alert_policy_ids}
	AlertPolicyIds *[]*string `field:"optional" json:"alertPolicyIds" yaml:"alertPolicyIds"`
	// New Relic User API Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#api_key DevopsagentService#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// List of monitored APM application IDs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#application_ids DevopsagentService#application_ids}
	ApplicationIds *[]*string `field:"optional" json:"applicationIds" yaml:"applicationIds"`
	// List of globally unique IDs for New Relic resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#entity_guids DevopsagentService#entity_guids}
	EntityGuids *[]*string `field:"optional" json:"entityGuids" yaml:"entityGuids"`
	// New Relic region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/devopsagent_service#region DevopsagentService#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

