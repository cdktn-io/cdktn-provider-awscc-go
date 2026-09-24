// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FisExperimentTemplateConfig struct {
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
	// A description for the experiment template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#description FisExperimentTemplate#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of an IAM role that grants the AWS FIS service permission to perform service actions on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#role_arn FisExperimentTemplate#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// One or more stop conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#stop_conditions FisExperimentTemplate#stop_conditions}
	StopConditions interface{} `field:"required" json:"stopConditions" yaml:"stopConditions"`
	// The targets for the experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#targets FisExperimentTemplate#targets}
	Targets interface{} `field:"required" json:"targets" yaml:"targets"`
	// The actions for the experiment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#actions FisExperimentTemplate#actions}
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#experiment_options FisExperimentTemplate#experiment_options}.
	ExperimentOptions *FisExperimentTemplateExperimentOptions `field:"optional" json:"experimentOptions" yaml:"experimentOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#experiment_report_configuration FisExperimentTemplate#experiment_report_configuration}.
	ExperimentReportConfiguration *FisExperimentTemplateExperimentReportConfiguration `field:"optional" json:"experimentReportConfiguration" yaml:"experimentReportConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#log_configuration FisExperimentTemplate#log_configuration}.
	LogConfiguration *FisExperimentTemplateLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#tags FisExperimentTemplate#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

