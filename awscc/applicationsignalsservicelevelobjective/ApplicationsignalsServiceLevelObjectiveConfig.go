// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApplicationsignalsServiceLevelObjectiveConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of this SLO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#name ApplicationsignalsServiceLevelObjective#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Each object in this array defines the length of the look-back window used to calculate one burn rate metric for this SLO.
	//
	// The burn rate measures how fast the service is consuming the error budget, relative to the attainment goal of the SLO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#burn_rate_configurations ApplicationsignalsServiceLevelObjective#burn_rate_configurations}
	BurnRateConfigurations interface{} `field:"optional" json:"burnRateConfigurations" yaml:"burnRateConfigurations"`
	// An optional description for this SLO. Default is 'No description'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#description ApplicationsignalsServiceLevelObjective#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Each object in this array defines a time exclusion window for this SLO.
	//
	// The time exclusion window is used to exclude breaching data points from affecting attainment rate, error budget, and burn rate metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#exclusion_windows ApplicationsignalsServiceLevelObjective#exclusion_windows}
	ExclusionWindows interface{} `field:"optional" json:"exclusionWindows" yaml:"exclusionWindows"`
	// A structure that contains the attributes that determine the goal of the SLO.
	//
	// This includes the time period for evaluation and the attainment threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#goal ApplicationsignalsServiceLevelObjective#goal}
	Goal *ApplicationsignalsServiceLevelObjectiveGoal `field:"optional" json:"goal" yaml:"goal"`
	// This structure contains information about the performance metric that a request-based SLO monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#request_based_sli ApplicationsignalsServiceLevelObjective#request_based_sli}
	RequestBasedSli *ApplicationsignalsServiceLevelObjectiveRequestBasedSli `field:"optional" json:"requestBasedSli" yaml:"requestBasedSli"`
	// This structure contains information about the performance metric that an SLO monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#sli ApplicationsignalsServiceLevelObjective#sli}
	Sli *ApplicationsignalsServiceLevelObjectiveSli `field:"optional" json:"sli" yaml:"sli"`
	// The list of tag keys and values associated with the resource you specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/applicationsignals_service_level_objective#tags ApplicationsignalsServiceLevelObjective#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

