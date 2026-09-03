// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscapability


type EksCapabilityConfigurationAck struct {
	// A list of ACK service names to disable. Controllers for services in this list are not installed or managed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_capability#disabled_services EksCapability#disabled_services}
	DisabledServices *[]*string `field:"optional" json:"disabledServices" yaml:"disabledServices"`
	// Whether cross-namespace references are enabled for ACK controllers. When not specified, the service default applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_capability#enable_cross_namespace EksCapability#enable_cross_namespace}
	EnableCrossNamespace interface{} `field:"optional" json:"enableCrossNamespace" yaml:"enableCrossNamespace"`
}

