// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrstep


type EmrStepHadoopJarStep struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/emr_step#jar EmrStep#jar}.
	Jar *string `field:"required" json:"jar" yaml:"jar"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/emr_step#args EmrStep#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/emr_step#main_class EmrStep#main_class}.
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/emr_step#step_properties EmrStep#step_properties}.
	StepProperties interface{} `field:"optional" json:"stepProperties" yaml:"stepProperties"`
}

