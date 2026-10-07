// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsinstrumentationconfig


type ApplicationsignalsInstrumentationConfigCaptureConfiguration struct {
	// Defines what data to capture for code-level instrumentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#code_capture ApplicationsignalsInstrumentationConfig#code_capture}
	CodeCapture *ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCapture `field:"required" json:"codeCapture" yaml:"codeCapture"`
}

