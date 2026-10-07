// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsinstrumentationconfig


type ApplicationsignalsInstrumentationConfigLocationCodeLocation struct {
	// The source file path relative to the project or source root.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#file_path ApplicationsignalsInstrumentationConfig#file_path}
	FilePath *string `field:"required" json:"filePath" yaml:"filePath"`
	// The programming language for this instrumentation point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#language ApplicationsignalsInstrumentationConfig#language}
	Language *string `field:"required" json:"language" yaml:"language"`
	// The class or type name that contains the method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#class_name ApplicationsignalsInstrumentationConfig#class_name}
	ClassName *string `field:"optional" json:"className" yaml:"className"`
	// The package, module, or namespace that contains the target code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#code_unit ApplicationsignalsInstrumentationConfig#code_unit}
	CodeUnit *string `field:"optional" json:"codeUnit" yaml:"codeUnit"`
	// The line number to instrument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#line_number ApplicationsignalsInstrumentationConfig#line_number}
	LineNumber *float64 `field:"optional" json:"lineNumber" yaml:"lineNumber"`
	// The method or function name to instrument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#method_name ApplicationsignalsInstrumentationConfig#method_name}
	MethodName *string `field:"optional" json:"methodName" yaml:"methodName"`
}

