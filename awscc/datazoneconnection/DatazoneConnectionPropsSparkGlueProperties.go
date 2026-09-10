// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsSparkGlueProperties struct {
	// Spark Glue Args.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#additional_args DatazoneConnection#additional_args}
	AdditionalArgs *DatazoneConnectionPropsSparkGluePropertiesAdditionalArgs `field:"optional" json:"additionalArgs" yaml:"additionalArgs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#glue_connection_name DatazoneConnection#glue_connection_name}.
	GlueConnectionName *string `field:"optional" json:"glueConnectionName" yaml:"glueConnectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#glue_version DatazoneConnection#glue_version}.
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#idle_timeout DatazoneConnection#idle_timeout}.
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#java_virtual_env DatazoneConnection#java_virtual_env}.
	JavaVirtualEnv *string `field:"optional" json:"javaVirtualEnv" yaml:"javaVirtualEnv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#number_of_workers DatazoneConnection#number_of_workers}.
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#python_virtual_env DatazoneConnection#python_virtual_env}.
	PythonVirtualEnv *string `field:"optional" json:"pythonVirtualEnv" yaml:"pythonVirtualEnv"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datazone_connection#worker_type DatazoneConnection#worker_type}.
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}

