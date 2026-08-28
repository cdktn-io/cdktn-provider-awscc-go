// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveSliSliMetricDependencyConfig struct {
	// If this SLO is related to a metric collected by Application Signals, you must use this field to specify which dependency the SLO metric is related to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationsignals_service_level_objective#dependency_key_attributes ApplicationsignalsServiceLevelObjective#dependency_key_attributes}
	DependencyKeyAttributes *map[string]*string `field:"optional" json:"dependencyKeyAttributes" yaml:"dependencyKeyAttributes"`
	// When the SLO monitors a specific operation of the dependency, this field specifies the name of that operation in the dependency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationsignals_service_level_objective#dependency_operation_name ApplicationsignalsServiceLevelObjective#dependency_operation_name}
	DependencyOperationName *string `field:"optional" json:"dependencyOperationName" yaml:"dependencyOperationName"`
}

