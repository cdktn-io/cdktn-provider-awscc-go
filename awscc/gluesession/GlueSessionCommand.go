// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluesession


type GlueSessionCommand struct {
	// Specifies the name of the SessionCommand. Can be 'glueetl' or 'gluestreaming'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_session#name GlueSession#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies the Python version. The Python version indicates the version supported for jobs of type Spark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/glue_session#python_version GlueSession#python_version}
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
}

