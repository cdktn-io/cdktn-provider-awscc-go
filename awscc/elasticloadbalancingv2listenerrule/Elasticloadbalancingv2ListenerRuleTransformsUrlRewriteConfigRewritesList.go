// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingv2listenerrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticloadbalancingv2listenerrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList interface {
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
	Get(index *float64) Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList
type jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList struct {
	internal.Type__cdktnComplexList
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewElasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2ListenerRule.Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewElasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList_Override(e Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingv2ListenerRule.Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) AllWithMapKey(mapKeyAttributeName *string) cdktn.DynamicListTerraformIterator {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) Get(index *float64) Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2ListenerRuleTransformsUrlRewriteConfigRewritesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

