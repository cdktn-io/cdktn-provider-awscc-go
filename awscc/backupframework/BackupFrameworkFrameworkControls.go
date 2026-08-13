// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupframework


type BackupFrameworkFrameworkControls struct {
	// The name of a control. This name is between 1 and 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/backup_framework#control_name BackupFramework#control_name}
	ControlName *string `field:"required" json:"controlName" yaml:"controlName"`
	// A list of ParameterName and ParameterValue pairs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/backup_framework#control_input_parameters BackupFramework#control_input_parameters}
	ControlInputParameters interface{} `field:"optional" json:"controlInputParameters" yaml:"controlInputParameters"`
	// The scope of a control.
	//
	// The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/backup_framework#control_scope BackupFramework#control_scope}
	ControlScope *BackupFrameworkFrameworkControlsControlScope `field:"optional" json:"controlScope" yaml:"controlScope"`
}

