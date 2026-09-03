// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveExclusionWindowsRecurrenceRule struct {
	// A cron or rate expression denoting how often to repeat this exclusion window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/applicationsignals_service_level_objective#expression ApplicationsignalsServiceLevelObjective#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

