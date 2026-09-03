// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionImageConfig struct {
	// Specifies parameters that you want to pass in with ENTRYPOINT.
	//
	// You can specify a maximum of 1,500 parameters in the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_function#command LambdaFunction#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Specifies the entry point to their application, which is typically the location of the runtime executable.
	//
	// You can specify a maximum of 1,500 string entries in the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_function#entry_point LambdaFunction#entry_point}
	EntryPoint *[]*string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// Specifies the working directory. The length of the directory string cannot exceed 1,000 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/lambda_function#working_directory LambdaFunction#working_directory}
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

