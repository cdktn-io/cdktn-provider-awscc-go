// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsinstrumentationconfig


type ApplicationsignalsInstrumentationConfigCaptureConfigurationCodeCaptureCaptureLimits struct {
	// Maximum nesting depth to traverse inside collections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#max_collection_depth ApplicationsignalsInstrumentationConfig#max_collection_depth}
	MaxCollectionDepth *float64 `field:"optional" json:"maxCollectionDepth" yaml:"maxCollectionDepth"`
	// Maximum number of items to capture from any collection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#max_collection_width ApplicationsignalsInstrumentationConfig#max_collection_width}
	MaxCollectionWidth *float64 `field:"optional" json:"maxCollectionWidth" yaml:"maxCollectionWidth"`
	// Maximum number of fields to capture for any object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#max_fields_per_object ApplicationsignalsInstrumentationConfig#max_fields_per_object}
	MaxFieldsPerObject *float64 `field:"optional" json:"maxFieldsPerObject" yaml:"maxFieldsPerObject"`
	// Maximum number of times the instrumentation point can be hit before disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#max_hits ApplicationsignalsInstrumentationConfig#max_hits}
	MaxHits *float64 `field:"optional" json:"maxHits" yaml:"maxHits"`
	// Maximum depth for nested object traversal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#max_object_depth ApplicationsignalsInstrumentationConfig#max_object_depth}
	MaxObjectDepth *float64 `field:"optional" json:"maxObjectDepth" yaml:"maxObjectDepth"`
	// Maximum number of stack frames to capture.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#max_stack_frames ApplicationsignalsInstrumentationConfig#max_stack_frames}
	MaxStackFrames *float64 `field:"optional" json:"maxStackFrames" yaml:"maxStackFrames"`
	// Maximum total size in bytes of a captured stack trace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#max_stack_trace_size ApplicationsignalsInstrumentationConfig#max_stack_trace_size}
	MaxStackTraceSize *float64 `field:"optional" json:"maxStackTraceSize" yaml:"maxStackTraceSize"`
	// Maximum length of captured string values in characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#max_string_length ApplicationsignalsInstrumentationConfig#max_string_length}
	MaxStringLength *float64 `field:"optional" json:"maxStringLength" yaml:"maxStringLength"`
}

