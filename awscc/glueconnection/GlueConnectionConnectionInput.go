// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueconnection


type GlueConnectionConnectionInput struct {
	// The type of the connection that needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#connection_type GlueConnection#connection_type}
	ConnectionType *string `field:"required" json:"connectionType" yaml:"connectionType"`
	// Connection properties specific to the Athena compute environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#athena_properties GlueConnection#athena_properties}
	AthenaProperties *string `field:"optional" json:"athenaProperties" yaml:"athenaProperties"`
	// The authentication configuration used to connect to the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#authentication_configuration GlueConnection#authentication_configuration}
	AuthenticationConfiguration *GlueConnectionConnectionInputAuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// A map of key-value pairs used as parameters for this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#connection_properties GlueConnection#connection_properties}
	ConnectionProperties *string `field:"optional" json:"connectionProperties" yaml:"connectionProperties"`
	// A description of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#description GlueConnection#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of criteria that can be used in selecting this connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#match_criteria GlueConnection#match_criteria}
	MatchCriteria *[]*string `field:"optional" json:"matchCriteria" yaml:"matchCriteria"`
	// The name of the connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#name GlueConnection#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The physical connection requirements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#physical_connection_requirements GlueConnection#physical_connection_requirements}
	PhysicalConnectionRequirements *GlueConnectionConnectionInputPhysicalConnectionRequirements `field:"optional" json:"physicalConnectionRequirements" yaml:"physicalConnectionRequirements"`
	// Connection properties specific to the Python compute environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#python_properties GlueConnection#python_properties}
	PythonProperties *string `field:"optional" json:"pythonProperties" yaml:"pythonProperties"`
	// Connection properties specific to the Spark compute environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#spark_properties GlueConnection#spark_properties}
	SparkProperties *string `field:"optional" json:"sparkProperties" yaml:"sparkProperties"`
	// A flag to validate the credentials during create connection. Default is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#validate_credentials GlueConnection#validate_credentials}
	ValidateCredentials interface{} `field:"optional" json:"validateCredentials" yaml:"validateCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/glue_connection#validate_for_compute_environments GlueConnection#validate_for_compute_environments}.
	ValidateForComputeEnvironments *[]*string `field:"optional" json:"validateForComputeEnvironments" yaml:"validateForComputeEnvironments"`
}

