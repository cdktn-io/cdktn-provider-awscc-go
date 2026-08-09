// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsHanaPrometheusExporter struct {
	// A flag which indicates agreeing to install SAP HANA DB client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#agree_to_install_hanadb_client ApplicationinsightsApplication#agree_to_install_hanadb_client}
	AgreeToInstallHanadbClient interface{} `field:"optional" json:"agreeToInstallHanadbClient" yaml:"agreeToInstallHanadbClient"`
	// The HANA DB port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#hana_port ApplicationinsightsApplication#hana_port}
	HanaPort *string `field:"optional" json:"hanaPort" yaml:"hanaPort"`
	// The secret name which manages the HANA DB credentials e.g. {   "username": "<>",   "password": "<>" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#hana_secret_name ApplicationinsightsApplication#hana_secret_name}
	HanaSecretName *string `field:"optional" json:"hanaSecretName" yaml:"hanaSecretName"`
	// HANA DB SID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#hanasid ApplicationinsightsApplication#hanasid}
	Hanasid *string `field:"optional" json:"hanasid" yaml:"hanasid"`
	// Prometheus exporter port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/applicationinsights_application#prometheus_port ApplicationinsightsApplication#prometheus_port}
	PrometheusPort *string `field:"optional" json:"prometheusPort" yaml:"prometheusPort"`
}

