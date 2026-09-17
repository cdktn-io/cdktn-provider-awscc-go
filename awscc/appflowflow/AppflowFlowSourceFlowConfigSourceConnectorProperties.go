// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowSourceFlowConfigSourceConnectorProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#amplitude AppflowFlow#amplitude}.
	Amplitude *AppflowFlowSourceFlowConfigSourceConnectorPropertiesAmplitude `field:"optional" json:"amplitude" yaml:"amplitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#custom_connector AppflowFlow#custom_connector}.
	CustomConnector *AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector `field:"optional" json:"customConnector" yaml:"customConnector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#datadog AppflowFlow#datadog}.
	Datadog *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDatadog `field:"optional" json:"datadog" yaml:"datadog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#dynatrace AppflowFlow#dynatrace}.
	Dynatrace *AppflowFlowSourceFlowConfigSourceConnectorPropertiesDynatrace `field:"optional" json:"dynatrace" yaml:"dynatrace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#google_analytics AppflowFlow#google_analytics}.
	GoogleAnalytics *AppflowFlowSourceFlowConfigSourceConnectorPropertiesGoogleAnalytics `field:"optional" json:"googleAnalytics" yaml:"googleAnalytics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#infor_nexus AppflowFlow#infor_nexus}.
	InforNexus *AppflowFlowSourceFlowConfigSourceConnectorPropertiesInforNexus `field:"optional" json:"inforNexus" yaml:"inforNexus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#marketo AppflowFlow#marketo}.
	Marketo *AppflowFlowSourceFlowConfigSourceConnectorPropertiesMarketo `field:"optional" json:"marketo" yaml:"marketo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#pardot AppflowFlow#pardot}.
	Pardot *AppflowFlowSourceFlowConfigSourceConnectorPropertiesPardot `field:"optional" json:"pardot" yaml:"pardot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#s3 AppflowFlow#s3}.
	S3 *AppflowFlowSourceFlowConfigSourceConnectorPropertiesS3 `field:"optional" json:"s3" yaml:"s3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#salesforce AppflowFlow#salesforce}.
	Salesforce *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSalesforce `field:"optional" json:"salesforce" yaml:"salesforce"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#sapo_data AppflowFlow#sapo_data}.
	SapoData *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSapoData `field:"optional" json:"sapoData" yaml:"sapoData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#service_now AppflowFlow#service_now}.
	ServiceNow *AppflowFlowSourceFlowConfigSourceConnectorPropertiesServiceNow `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#singular AppflowFlow#singular}.
	Singular *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSingular `field:"optional" json:"singular" yaml:"singular"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#slack AppflowFlow#slack}.
	Slack *AppflowFlowSourceFlowConfigSourceConnectorPropertiesSlack `field:"optional" json:"slack" yaml:"slack"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#trendmicro AppflowFlow#trendmicro}.
	Trendmicro *AppflowFlowSourceFlowConfigSourceConnectorPropertiesTrendmicro `field:"optional" json:"trendmicro" yaml:"trendmicro"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#veeva AppflowFlow#veeva}.
	Veeva *AppflowFlowSourceFlowConfigSourceConnectorPropertiesVeeva `field:"optional" json:"veeva" yaml:"veeva"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/appflow_flow#zendesk AppflowFlow#zendesk}.
	Zendesk *AppflowFlowSourceFlowConfigSourceConnectorPropertiesZendesk `field:"optional" json:"zendesk" yaml:"zendesk"`
}

