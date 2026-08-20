// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfiguration struct {
	// The Amazon Glue Data Catalog that you use in queries in a Kinesis Data Analytics Studio notebook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#catalog_configuration Kinesisanalyticsv2Application#catalog_configuration}
	CatalogConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationCatalogConfiguration `field:"optional" json:"catalogConfiguration" yaml:"catalogConfiguration"`
	// A list of CustomArtifactConfiguration objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#custom_artifacts_configuration Kinesisanalyticsv2Application#custom_artifacts_configuration}
	CustomArtifactsConfiguration interface{} `field:"optional" json:"customArtifactsConfiguration" yaml:"customArtifactsConfiguration"`
	// The information required to deploy a Kinesis Data Analytics Studio notebook as an application with durable state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#deploy_as_application_configuration Kinesisanalyticsv2Application#deploy_as_application_configuration}
	DeployAsApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationDeployAsApplicationConfiguration `field:"optional" json:"deployAsApplicationConfiguration" yaml:"deployAsApplicationConfiguration"`
	// The monitoring configuration of a Kinesis Data Analytics Studio notebook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#monitoring_configuration Kinesisanalyticsv2Application#monitoring_configuration}
	MonitoringConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfigurationMonitoringConfiguration `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
}

