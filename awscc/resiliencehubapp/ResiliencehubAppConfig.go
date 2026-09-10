// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubapp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ResiliencehubAppConfig struct {
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
	// A string containing full ResilienceHub app template body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehub_app#app_template_body ResiliencehubApp#app_template_body}
	AppTemplateBody *string `field:"required" json:"appTemplateBody" yaml:"appTemplateBody"`
	// Name of the app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehub_app#name ResiliencehubApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of ResourceMapping objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehub_app#resource_mappings ResiliencehubApp#resource_mappings}
	ResourceMappings interface{} `field:"required" json:"resourceMappings" yaml:"resourceMappings"`
	// Assessment execution schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehub_app#app_assessment_schedule ResiliencehubApp#app_assessment_schedule}
	AppAssessmentSchedule *string `field:"optional" json:"appAssessmentSchedule" yaml:"appAssessmentSchedule"`
	// App description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehub_app#description ResiliencehubApp#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The list of events you would like to subscribe and get notification for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehub_app#event_subscriptions ResiliencehubApp#event_subscriptions}
	EventSubscriptions interface{} `field:"optional" json:"eventSubscriptions" yaml:"eventSubscriptions"`
	// Defines the roles and credentials that AWS Resilience Hub would use while creating the application, importing its resources, and running an assessment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehub_app#permission_model ResiliencehubApp#permission_model}
	PermissionModel *ResiliencehubAppPermissionModel `field:"optional" json:"permissionModel" yaml:"permissionModel"`
	// Amazon Resource Name (ARN) of the Resiliency Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehub_app#resiliency_policy_arn ResiliencehubApp#resiliency_policy_arn}
	ResiliencyPolicyArn *string `field:"optional" json:"resiliencyPolicyArn" yaml:"resiliencyPolicyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/resiliencehub_app#tags ResiliencehubApp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

