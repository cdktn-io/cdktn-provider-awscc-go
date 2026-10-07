// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsinstrumentationconfig


type ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCapture struct {
	// Safety limits that bound what is captured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#capture_limits ApplicationsignalsInstrumentationConfig#capture_limits}
	CaptureLimits *ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimits `field:"required" json:"captureLimits" yaml:"captureLimits"`
	// The function arguments to capture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#capture_arguments ApplicationsignalsInstrumentationConfig#capture_arguments}
	CaptureArguments *[]*string `field:"optional" json:"captureArguments" yaml:"captureArguments"`
	// The local variables to capture by name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#capture_locals ApplicationsignalsInstrumentationConfig#capture_locals}
	CaptureLocals *[]*string `field:"optional" json:"captureLocals" yaml:"captureLocals"`
	// Whether to capture the return value. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#capture_return ApplicationsignalsInstrumentationConfig#capture_return}
	CaptureReturn interface{} `field:"optional" json:"captureReturn" yaml:"captureReturn"`
	// Whether to capture a stack trace. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#capture_stack_trace ApplicationsignalsInstrumentationConfig#capture_stack_trace}
	CaptureStackTrace interface{} `field:"optional" json:"captureStackTrace" yaml:"captureStackTrace"`
}

