// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationMaximumCapacity struct {
	// Per worker CPU resource. vCPU is the only supported unit and specifying vCPU is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/emrserverless_application#cpu EmrserverlessApplication#cpu}
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Per worker Disk resource. GB is the only supported unit and specifying GB is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/emrserverless_application#disk EmrserverlessApplication#disk}
	Disk *string `field:"optional" json:"disk" yaml:"disk"`
	// Per worker memory resource. GB is the only supported unit and specifying GB is optional.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/emrserverless_application#memory EmrserverlessApplication#memory}
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}

