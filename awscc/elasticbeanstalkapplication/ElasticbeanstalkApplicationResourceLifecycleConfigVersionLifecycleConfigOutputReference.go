// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticbeanstalkapplication/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference interface {
	cdktn.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxAgeRule() ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxAgeRuleOutputReference
	MaxAgeRuleInput() interface{}
	MaxCountRule() ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxCountRuleOutputReference
	MaxCountRuleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutMaxAgeRule(value *ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxAgeRule)
	PutMaxCountRule(value *ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxCountRule)
	ResetMaxAgeRule()
	ResetMaxCountRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference
type jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) MaxAgeRule() ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxAgeRuleOutputReference {
	var returns ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxAgeRuleOutputReference
	_jsii_.Get(
		j,
		"maxAgeRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) MaxAgeRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxAgeRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) MaxCountRule() ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxCountRuleOutputReference {
	var returns ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxCountRuleOutputReference
	_jsii_.Get(
		j,
		"maxCountRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) MaxCountRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxCountRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference {
	_init_.Initialize()

	if err := validateNewElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticbeanstalkApplication.ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference_Override(e ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticbeanstalkApplication.ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) PutMaxAgeRule(value *ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxAgeRule) {
	if err := e.validatePutMaxAgeRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMaxAgeRule",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) PutMaxCountRule(value *ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigMaxCountRule) {
	if err := e.validatePutMaxCountRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMaxCountRule",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) ResetMaxAgeRule() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxAgeRule",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) ResetMaxCountRule() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxCountRule",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElasticbeanstalkApplicationResourceLifecycleConfigVersionLifecycleConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

