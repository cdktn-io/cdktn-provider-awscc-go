// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionmatchingworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/entityresolutionmatchingworkflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList interface {
	cdktn.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WrapsSet() *bool
	// Experimental.
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList
type jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList {
	_init_.Initialize()

	if err := validateNewEntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList_Override(e EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.entityresolutionMatchingWorkflow.EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
	if err := e.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktn.DynamicListTerraformIterator

	_jsii_.Invoke(
		e,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) Get(index *float64) EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EntityresolutionMatchingWorkflowResolutionTechniquesRuleConditionPropertiesRulesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

