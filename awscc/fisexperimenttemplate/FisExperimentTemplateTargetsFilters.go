// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateTargetsFilters struct {
	// The attribute path for the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#path FisExperimentTemplate#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// The attribute values for the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#values FisExperimentTemplate#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

