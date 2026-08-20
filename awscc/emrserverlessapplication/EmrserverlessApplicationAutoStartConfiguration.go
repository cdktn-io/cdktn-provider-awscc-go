// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationAutoStartConfiguration struct {
	// If set to true, the Application will automatically start. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/emrserverless_application#enabled EmrserverlessApplication#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

