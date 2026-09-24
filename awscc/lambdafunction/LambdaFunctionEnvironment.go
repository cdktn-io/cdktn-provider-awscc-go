// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction


type LambdaFunctionEnvironment struct {
	// Environment variable key-value pairs.
	//
	// For more information, see [Using Lambda environment variables](https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html).
	//  If the value of the environment variable is a time or a duration, enclose the value in quotes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lambda_function#variables LambdaFunction#variables}
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}

