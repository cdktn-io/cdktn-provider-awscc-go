// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfiguration struct {
	// The code location and type parameters for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#application_code_configuration Kinesisanalyticsv2Application#application_code_configuration}
	ApplicationCodeConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationCodeConfiguration `field:"optional" json:"applicationCodeConfiguration" yaml:"applicationCodeConfiguration"`
	// Describes whether customer managed key is enabled and key details for customer data encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#application_encryption_configuration Kinesisanalyticsv2Application#application_encryption_configuration}
	ApplicationEncryptionConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationEncryptionConfiguration `field:"optional" json:"applicationEncryptionConfiguration" yaml:"applicationEncryptionConfiguration"`
	// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#application_snapshot_configuration Kinesisanalyticsv2Application#application_snapshot_configuration}
	ApplicationSnapshotConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSnapshotConfiguration `field:"optional" json:"applicationSnapshotConfiguration" yaml:"applicationSnapshotConfiguration"`
	// Describes whether system initiated rollbacks are enabled for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#application_system_rollback_configuration Kinesisanalyticsv2Application#application_system_rollback_configuration}
	ApplicationSystemRollbackConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationApplicationSystemRollbackConfiguration `field:"optional" json:"applicationSystemRollbackConfiguration" yaml:"applicationSystemRollbackConfiguration"`
	// Describes execution properties for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#environment_properties Kinesisanalyticsv2Application#environment_properties}
	EnvironmentProperties *Kinesisanalyticsv2ApplicationApplicationConfigurationEnvironmentProperties `field:"optional" json:"environmentProperties" yaml:"environmentProperties"`
	// The creation and update parameters for a Flink-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#flink_application_configuration Kinesisanalyticsv2Application#flink_application_configuration}
	FlinkApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationFlinkApplicationConfiguration `field:"optional" json:"flinkApplicationConfiguration" yaml:"flinkApplicationConfiguration"`
	// The creation and update parameters for a SQL-based Kinesis Data Analytics application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#sql_application_configuration Kinesisanalyticsv2Application#sql_application_configuration}
	SqlApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfiguration `field:"optional" json:"sqlApplicationConfiguration" yaml:"sqlApplicationConfiguration"`
	// The array of descriptions of VPC configurations available to the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#vpc_configurations Kinesisanalyticsv2Application#vpc_configurations}
	VpcConfigurations interface{} `field:"optional" json:"vpcConfigurations" yaml:"vpcConfigurations"`
	// The configuration parameters for a Kinesis Data Analytics Studio notebook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisanalyticsv2_application#zeppelin_application_configuration Kinesisanalyticsv2Application#zeppelin_application_configuration}
	ZeppelinApplicationConfiguration *Kinesisanalyticsv2ApplicationApplicationConfigurationZeppelinApplicationConfiguration `field:"optional" json:"zeppelinApplicationConfiguration" yaml:"zeppelinApplicationConfiguration"`
}

