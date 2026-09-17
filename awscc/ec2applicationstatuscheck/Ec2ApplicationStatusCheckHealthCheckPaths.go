// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2applicationstatuscheck


type Ec2ApplicationStatusCheckHealthCheckPaths struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_application_status_check#destinations Ec2ApplicationStatusCheck#destinations}.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ec2_application_status_check#source Ec2ApplicationStatusCheck#source}.
	Source *Ec2ApplicationStatusCheckHealthCheckPathsSource `field:"optional" json:"source" yaml:"source"`
}

