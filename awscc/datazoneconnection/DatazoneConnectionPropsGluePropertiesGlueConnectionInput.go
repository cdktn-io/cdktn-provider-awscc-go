// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection


type DatazoneConnectionPropsGluePropertiesGlueConnectionInput struct {
	// Property Map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#athena_properties DatazoneConnection#athena_properties}
	AthenaProperties *map[string]*string `field:"optional" json:"athenaProperties" yaml:"athenaProperties"`
	// Authentication Configuration Input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#authentication_configuration DatazoneConnection#authentication_configuration}
	AuthenticationConfiguration *DatazoneConnectionPropsGluePropertiesGlueConnectionInputAuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// Connection Properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#connection_properties DatazoneConnection#connection_properties}
	ConnectionProperties *map[string]*string `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// Glue Connection Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#connection_type DatazoneConnection#connection_type}
	ConnectionType *string `field:"optional" json:"connectionType" yaml:"connectionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#description DatazoneConnection#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#match_criteria DatazoneConnection#match_criteria}.
	MatchCriteria *string `field:"optional" json:"matchCriteria" yaml:"matchCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#name DatazoneConnection#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Physical Connection Requirements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#physical_connection_requirements DatazoneConnection#physical_connection_requirements}
	PhysicalConnectionRequirements *DatazoneConnectionPropsGluePropertiesGlueConnectionInputPhysicalConnectionRequirements `field:"optional" json:"physicalConnectionRequirements" yaml:"physicalConnectionRequirements"`
	// Property Map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#python_properties DatazoneConnection#python_properties}
	PythonProperties *map[string]*string `field:"optional" json:"pythonProperties" yaml:"pythonProperties"`
	// Property Map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#spark_properties DatazoneConnection#spark_properties}
	SparkProperties *map[string]*string `field:"optional" json:"sparkProperties" yaml:"sparkProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#validate_credentials DatazoneConnection#validate_credentials}.
	ValidateCredentials interface{} `field:"optional" json:"validateCredentials" yaml:"validateCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datazone_connection#validate_for_compute_environments DatazoneConnection#validate_for_compute_environments}.
	ValidateForComputeEnvironments *[]*string `field:"optional" json:"validateForComputeEnvironments" yaml:"validateForComputeEnvironments"`
}

