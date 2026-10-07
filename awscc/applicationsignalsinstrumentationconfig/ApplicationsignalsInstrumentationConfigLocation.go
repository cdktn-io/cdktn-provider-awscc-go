// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsinstrumentationconfig


type ApplicationsignalsInstrumentationConfigLocation struct {
	// Identifies a code location to instrument.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/applicationsignals_instrumentation_config#code_location ApplicationsignalsInstrumentationConfig#code_location}
	CodeLocation *ApplicationsignalsInstrumentationConfigLocationCodeLocation `field:"required" json:"codeLocation" yaml:"codeLocation"`
}

