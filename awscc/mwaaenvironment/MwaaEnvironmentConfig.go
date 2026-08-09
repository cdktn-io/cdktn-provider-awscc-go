// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwaaenvironment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MwaaEnvironmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Customer-defined identifier for the environment, unique per customer region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#name MwaaEnvironment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Key/value pairs representing Airflow configuration variables.     Keys are prefixed by their section:.
	//
	// [core]
	//     dags_folder={AIRFLOW_HOME}/dags
	//
	//     Would be represented as
	//
	//     "core.dags_folder": "{AIRFLOW_HOME}/dags"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#airflow_configuration_options MwaaEnvironment#airflow_configuration_options}
	AirflowConfigurationOptions *string `field:"optional" json:"airflowConfigurationOptions" yaml:"airflowConfigurationOptions"`
	// Version of airflow to deploy to the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#airflow_version MwaaEnvironment#airflow_version}
	AirflowVersion *string `field:"optional" json:"airflowVersion" yaml:"airflowVersion"`
	// Represents an S3 prefix relative to the root of an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#dag_s3_path MwaaEnvironment#dag_s3_path}
	DagS3Path *string `field:"optional" json:"dagS3Path" yaml:"dagS3Path"`
	// Defines whether the VPC endpoints configured for the environment are created, and managed, by the customer or by Amazon MWAA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#endpoint_management MwaaEnvironment#endpoint_management}
	EndpointManagement *string `field:"optional" json:"endpointManagement" yaml:"endpointManagement"`
	// Templated configuration for airflow processes and backing infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#environment_class MwaaEnvironment#environment_class}
	EnvironmentClass *string `field:"optional" json:"environmentClass" yaml:"environmentClass"`
	// IAM role to be used by tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#execution_role_arn MwaaEnvironment#execution_role_arn}
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for MWAA data encryption.
	//
	// You can specify the CMK using any of the following:
	//
	//     Key ID. For example, key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	//     Key alias. For example, alias/ExampleAlias.
	//
	//     Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	//     Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	//     AWS authenticates the CMK asynchronously. Therefore, if you specify an ID, alias, or ARN that is not valid, the action can appear to complete, but eventually fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#kms_key MwaaEnvironment#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Logging configuration for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#logging_configuration MwaaEnvironment#logging_configuration}
	LoggingConfiguration *MwaaEnvironmentLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Maximum webserver compute units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#max_webservers MwaaEnvironment#max_webservers}
	MaxWebservers *float64 `field:"optional" json:"maxWebservers" yaml:"maxWebservers"`
	// Maximum worker compute units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#max_workers MwaaEnvironment#max_workers}
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// Minimum webserver compute units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#min_webservers MwaaEnvironment#min_webservers}
	MinWebservers *float64 `field:"optional" json:"minWebservers" yaml:"minWebservers"`
	// Minimum worker compute units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#min_workers MwaaEnvironment#min_workers}
	MinWorkers *float64 `field:"optional" json:"minWorkers" yaml:"minWorkers"`
	// Configures the network resources of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#network_configuration MwaaEnvironment#network_configuration}
	NetworkConfiguration *MwaaEnvironmentNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Represents an version ID for an S3 object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#plugins_s3_object_version MwaaEnvironment#plugins_s3_object_version}
	PluginsS3ObjectVersion *string `field:"optional" json:"pluginsS3ObjectVersion" yaml:"pluginsS3ObjectVersion"`
	// Represents an S3 prefix relative to the root of an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#plugins_s3_path MwaaEnvironment#plugins_s3_path}
	PluginsS3Path *string `field:"optional" json:"pluginsS3Path" yaml:"pluginsS3Path"`
	// Represents an version ID for an S3 object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#requirements_s3_object_version MwaaEnvironment#requirements_s3_object_version}
	RequirementsS3ObjectVersion *string `field:"optional" json:"requirementsS3ObjectVersion" yaml:"requirementsS3ObjectVersion"`
	// Represents an S3 prefix relative to the root of an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#requirements_s3_path MwaaEnvironment#requirements_s3_path}
	RequirementsS3Path *string `field:"optional" json:"requirementsS3Path" yaml:"requirementsS3Path"`
	// Scheduler compute units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#schedulers MwaaEnvironment#schedulers}
	Schedulers *float64 `field:"optional" json:"schedulers" yaml:"schedulers"`
	// ARN for the AWS S3 bucket to use as the source of DAGs and plugins for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#source_bucket_arn MwaaEnvironment#source_bucket_arn}
	SourceBucketArn *string `field:"optional" json:"sourceBucketArn" yaml:"sourceBucketArn"`
	// Represents an version ID for an S3 object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#startup_script_s3_object_version MwaaEnvironment#startup_script_s3_object_version}
	StartupScriptS3ObjectVersion *string `field:"optional" json:"startupScriptS3ObjectVersion" yaml:"startupScriptS3ObjectVersion"`
	// Represents an S3 prefix relative to the root of an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#startup_script_s3_path MwaaEnvironment#startup_script_s3_path}
	StartupScriptS3Path *string `field:"optional" json:"startupScriptS3Path" yaml:"startupScriptS3Path"`
	// A map of tags for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#tags MwaaEnvironment#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
	// Choice for mode of webserver access including over public internet or via private VPC endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#webserver_access_mode MwaaEnvironment#webserver_access_mode}
	WebserverAccessMode *string `field:"optional" json:"webserverAccessMode" yaml:"webserverAccessMode"`
	// Start time for the weekly maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#weekly_maintenance_window_start MwaaEnvironment#weekly_maintenance_window_start}
	WeeklyMaintenanceWindowStart *string `field:"optional" json:"weeklyMaintenanceWindowStart" yaml:"weeklyMaintenanceWindowStart"`
	// The worker replacement strategy to use when updating the environment.
	//
	// Valid values: `FORCED`, `GRACEFUL`. FORCED means Apache Airflow workers will be stopped and replaced without waiting for tasks to complete before an update. GRACEFUL means Apache Airflow workers will be able to complete running tasks for up to 12 hours during an update before being stopped and replaced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/mwaa_environment#worker_replacement_strategy MwaaEnvironment#worker_replacement_strategy}
	WorkerReplacementStrategy *string `field:"optional" json:"workerReplacementStrategy" yaml:"workerReplacementStrategy"`
}

