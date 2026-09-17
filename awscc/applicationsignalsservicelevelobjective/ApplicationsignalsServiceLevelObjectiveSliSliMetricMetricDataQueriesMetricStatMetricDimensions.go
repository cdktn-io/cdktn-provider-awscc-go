// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveSliSliMetricMetricDataQueriesMetricStatMetricDimensions struct {
	// The name of the dimension.
	//
	// Dimension names must contain only ASCII characters, must include at least one non-whitespace character, and cannot start with a colon (:). ASCII control characters are not supported as part of dimension names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#name ApplicationsignalsServiceLevelObjective#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the dimension.
	//
	// Dimension values must contain only ASCII characters and must include at least one non-whitespace character. ASCII control characters are not supported as part of dimension values
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#value ApplicationsignalsServiceLevelObjective#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

