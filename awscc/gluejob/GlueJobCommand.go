// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluejob


type GlueJobCommand struct {
	// The name of the job command.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#name GlueJob#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The Python version being used to execute a Python shell job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#python_version GlueJob#python_version}
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Runtime is used to specify the versions of Ray, Python and additional libraries available in your environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#runtime GlueJob#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes a job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/glue_job#script_location GlueJob#script_location}
	ScriptLocation *string `field:"optional" json:"scriptLocation" yaml:"scriptLocation"`
}

