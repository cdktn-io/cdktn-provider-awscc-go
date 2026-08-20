// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanageringresspoint

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesMailManagerIngressPointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_ingress_point#rule_set_id SesMailManagerIngressPoint#rule_set_id}.
	RuleSetId *string `field:"required" json:"ruleSetId" yaml:"ruleSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_ingress_point#traffic_policy_id SesMailManagerIngressPoint#traffic_policy_id}.
	TrafficPolicyId *string `field:"required" json:"trafficPolicyId" yaml:"trafficPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_ingress_point#type SesMailManagerIngressPoint#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_ingress_point#ingress_point_configuration SesMailManagerIngressPoint#ingress_point_configuration}.
	IngressPointConfiguration *SesMailManagerIngressPointIngressPointConfiguration `field:"optional" json:"ingressPointConfiguration" yaml:"ingressPointConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_ingress_point#ingress_point_name SesMailManagerIngressPoint#ingress_point_name}.
	IngressPointName *string `field:"optional" json:"ingressPointName" yaml:"ingressPointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_ingress_point#network_configuration SesMailManagerIngressPoint#network_configuration}.
	NetworkConfiguration *SesMailManagerIngressPointNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_ingress_point#status_to_update SesMailManagerIngressPoint#status_to_update}.
	StatusToUpdate *string `field:"optional" json:"statusToUpdate" yaml:"statusToUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_ingress_point#tags SesMailManagerIngressPoint#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_mail_manager_ingress_point#tls_policy SesMailManagerIngressPoint#tls_policy}.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}

