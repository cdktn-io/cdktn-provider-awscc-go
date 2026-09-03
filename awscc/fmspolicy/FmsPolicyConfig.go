// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fmspolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FmsPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#exclude_resource_tags FmsPolicy#exclude_resource_tags}.
	ExcludeResourceTags interface{} `field:"required" json:"excludeResourceTags" yaml:"excludeResourceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#policy_name FmsPolicy#policy_name}.
	PolicyName *string `field:"required" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#remediation_enabled FmsPolicy#remediation_enabled}.
	RemediationEnabled interface{} `field:"required" json:"remediationEnabled" yaml:"remediationEnabled"`
	// Firewall security service policy data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#security_service_policy_data FmsPolicy#security_service_policy_data}
	SecurityServicePolicyData *FmsPolicySecurityServicePolicyData `field:"required" json:"securityServicePolicyData" yaml:"securityServicePolicyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#delete_all_policy_resources FmsPolicy#delete_all_policy_resources}.
	DeleteAllPolicyResources interface{} `field:"optional" json:"deleteAllPolicyResources" yaml:"deleteAllPolicyResources"`
	// An FMS includeMap or excludeMap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#exclude_map FmsPolicy#exclude_map}
	ExcludeMap *FmsPolicyExcludeMap `field:"optional" json:"excludeMap" yaml:"excludeMap"`
	// An FMS includeMap or excludeMap.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#include_map FmsPolicy#include_map}
	IncludeMap *FmsPolicyIncludeMap `field:"optional" json:"includeMap" yaml:"includeMap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#policy_description FmsPolicy#policy_description}.
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#resources_clean_up FmsPolicy#resources_clean_up}.
	ResourcesCleanUp interface{} `field:"optional" json:"resourcesCleanUp" yaml:"resourcesCleanUp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#resource_set_ids FmsPolicy#resource_set_ids}.
	ResourceSetIds *[]*string `field:"optional" json:"resourceSetIds" yaml:"resourceSetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#resource_tag_logical_operator FmsPolicy#resource_tag_logical_operator}.
	ResourceTagLogicalOperator *string `field:"optional" json:"resourceTagLogicalOperator" yaml:"resourceTagLogicalOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#resource_tags FmsPolicy#resource_tags}.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// An AWS resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#resource_type FmsPolicy#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#resource_type_list FmsPolicy#resource_type_list}.
	ResourceTypeList *[]*string `field:"optional" json:"resourceTypeList" yaml:"resourceTypeList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fms_policy#tags FmsPolicy#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

