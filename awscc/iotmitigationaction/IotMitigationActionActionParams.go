// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotmitigationaction


type IotMitigationActionActionParams struct {
	// Parameters to define a mitigation action that moves devices associated with a certificate to one or more specified thing groups, typically for quarantine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_mitigation_action#add_things_to_thing_group_params IotMitigationAction#add_things_to_thing_group_params}
	AddThingsToThingGroupParams *IotMitigationActionActionParamsAddThingsToThingGroupParams `field:"optional" json:"addThingsToThingGroupParams" yaml:"addThingsToThingGroupParams"`
	// Parameters to define a mitigation action that enables AWS IoT logging at a specified level of detail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_mitigation_action#enable_io_t_logging_params IotMitigationAction#enable_io_t_logging_params}
	EnableIoTLoggingParams *IotMitigationActionActionParamsEnableIoTLoggingParams `field:"optional" json:"enableIoTLoggingParams" yaml:"enableIoTLoggingParams"`
	// Parameters, to define a mitigation action that publishes findings to Amazon SNS.
	//
	// You can implement your own custom actions in response to the Amazon SNS messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_mitigation_action#publish_finding_to_sns_params IotMitigationAction#publish_finding_to_sns_params}
	PublishFindingToSnsParams *IotMitigationActionActionParamsPublishFindingToSnsParams `field:"optional" json:"publishFindingToSnsParams" yaml:"publishFindingToSnsParams"`
	// Parameters to define a mitigation action that adds a blank policy to restrict permissions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_mitigation_action#replace_default_policy_version_params IotMitigationAction#replace_default_policy_version_params}
	ReplaceDefaultPolicyVersionParams *IotMitigationActionActionParamsReplaceDefaultPolicyVersionParams `field:"optional" json:"replaceDefaultPolicyVersionParams" yaml:"replaceDefaultPolicyVersionParams"`
	// Parameters to define a mitigation action that changes the state of the CA certificate to inactive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_mitigation_action#update_ca_certificate_params IotMitigationAction#update_ca_certificate_params}
	UpdateCaCertificateParams *IotMitigationActionActionParamsUpdateCaCertificateParams `field:"optional" json:"updateCaCertificateParams" yaml:"updateCaCertificateParams"`
	// Parameters to define a mitigation action that changes the state of the device certificate to inactive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iot_mitigation_action#update_device_certificate_params IotMitigationAction#update_device_certificate_params}
	UpdateDeviceCertificateParams *IotMitigationActionActionParamsUpdateDeviceCertificateParams `field:"optional" json:"updateDeviceCertificateParams" yaml:"updateDeviceCertificateParams"`
}

