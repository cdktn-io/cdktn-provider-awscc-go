// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfiguration struct {
	// Additional Configuration that are passed to Athena Spark Calculations running in this workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#additional_configuration AthenaWorkGroup#additional_configuration}
	AdditionalConfiguration *string `field:"optional" json:"additionalConfiguration" yaml:"additionalConfiguration"`
	// The upper data usage limit (cutoff) for the amount of bytes a single query in a workgroup is allowed to scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#bytes_scanned_cutoff_per_query AthenaWorkGroup#bytes_scanned_cutoff_per_query}
	BytesScannedCutoffPerQuery *float64 `field:"optional" json:"bytesScannedCutoffPerQuery" yaml:"bytesScannedCutoffPerQuery"`
	// Indicates the KMS key for encrypting notebook content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#customer_content_encryption_configuration AthenaWorkGroup#customer_content_encryption_configuration}
	CustomerContentEncryptionConfiguration *AthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfiguration `field:"optional" json:"customerContentEncryptionConfiguration" yaml:"customerContentEncryptionConfiguration"`
	// If set to "true", the settings for the workgroup override client-side settings.
	//
	// If set to "false", client-side settings are used
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#enforce_work_group_configuration AthenaWorkGroup#enforce_work_group_configuration}
	EnforceWorkGroupConfiguration interface{} `field:"optional" json:"enforceWorkGroupConfiguration" yaml:"enforceWorkGroupConfiguration"`
	// The engine configuration for running queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#engine_configuration AthenaWorkGroup#engine_configuration}
	EngineConfiguration *AthenaWorkGroupWorkGroupConfigurationEngineConfiguration `field:"optional" json:"engineConfiguration" yaml:"engineConfiguration"`
	// The Athena engine version for running queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#engine_version AthenaWorkGroup#engine_version}
	EngineVersion *AthenaWorkGroupWorkGroupConfigurationEngineVersion `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The ARN of the execution role used to access user resources for Spark sessions and Identity Center enabled workgroups.
	//
	// This property applies only to Spark enabled workgroups and Identity Center enabled workgroups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#execution_role AthenaWorkGroup#execution_role}
	ExecutionRole *string `field:"optional" json:"executionRole" yaml:"executionRole"`
	// The configuration for the managed query results and encryption option.
	//
	// ResultConfiguration and ManagedQueryResultsConfiguration cannot be set at the same time
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#managed_query_results_configuration AthenaWorkGroup#managed_query_results_configuration}
	ManagedQueryResultsConfiguration *AthenaWorkGroupWorkGroupConfigurationManagedQueryResultsConfiguration `field:"optional" json:"managedQueryResultsConfiguration" yaml:"managedQueryResultsConfiguration"`
	// Contains the configuration settings for managed log persistence, delivering logs to Amazon S3 buckets, Amazon CloudWatch log groups etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#monitoring_configuration AthenaWorkGroup#monitoring_configuration}
	MonitoringConfiguration *AthenaWorkGroupWorkGroupConfigurationMonitoringConfiguration `field:"optional" json:"monitoringConfiguration" yaml:"monitoringConfiguration"`
	// Indicates that the Amazon CloudWatch metrics are enabled for the workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#publish_cloudwatch_metrics_enabled AthenaWorkGroup#publish_cloudwatch_metrics_enabled}
	PublishCloudwatchMetricsEnabled interface{} `field:"optional" json:"publishCloudwatchMetricsEnabled" yaml:"publishCloudwatchMetricsEnabled"`
	// If set to true, allows members assigned to a workgroup to reference Amazon S3 Requester Pays buckets in queries.
	//
	// If set to false, workgroup members cannot query data from Requester Pays buckets, and queries that retrieve data from Requester Pays buckets cause an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#requester_pays_enabled AthenaWorkGroup#requester_pays_enabled}
	RequesterPaysEnabled interface{} `field:"optional" json:"requesterPaysEnabled" yaml:"requesterPaysEnabled"`
	// The location in Amazon S3 where query results are stored and the encryption option, if any, used for query results.
	//
	// These are known as "client-side settings". If workgroup settings override client-side settings, then the query uses the workgroup settings.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/athena_work_group#result_configuration AthenaWorkGroup#result_configuration}
	ResultConfiguration *AthenaWorkGroupWorkGroupConfigurationResultConfiguration `field:"optional" json:"resultConfiguration" yaml:"resultConfiguration"`
}

