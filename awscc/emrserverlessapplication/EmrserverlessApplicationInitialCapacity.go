// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationInitialCapacity struct {
	// Worker type for an analytics framework.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/emrserverless_application#key EmrserverlessApplication#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/emrserverless_application#value EmrserverlessApplication#value}.
	Value *EmrserverlessApplicationInitialCapacityValue `field:"optional" json:"value" yaml:"value"`
}

