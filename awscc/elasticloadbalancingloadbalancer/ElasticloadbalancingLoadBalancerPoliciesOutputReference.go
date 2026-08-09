// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticloadbalancingloadbalancer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/elasticloadbalancingloadbalancer/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticloadbalancingLoadBalancerPoliciesOutputReference interface {
	cdktn.ComplexObject
	Attributes() ElasticloadbalancingLoadBalancerPoliciesAttributesList
	AttributesInput() interface{}
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
	InstancePorts() *[]*string
	SetInstancePorts(val *[]*string)
	InstancePortsInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancerPorts() *[]*string
	SetLoadBalancerPorts(val *[]*string)
	LoadBalancerPortsInput() *[]*string
	PolicyName() *string
	SetPolicyName(val *string)
	PolicyNameInput() *string
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
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
	PutAttributes(value interface{})
	ResetAttributes()
	ResetInstancePorts()
	ResetLoadBalancerPorts()
	ResetPolicyName()
	ResetPolicyType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ElasticloadbalancingLoadBalancerPoliciesOutputReference
type jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) Attributes() ElasticloadbalancingLoadBalancerPoliciesAttributesList {
	var returns ElasticloadbalancingLoadBalancerPoliciesAttributesList
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) AttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) InstancePorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instancePorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) InstancePortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instancePortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) LoadBalancerPorts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerPorts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) LoadBalancerPortsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerPortsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) PolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticloadbalancingLoadBalancerPoliciesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ElasticloadbalancingLoadBalancerPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingLoadBalancerPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingLoadBalancer.ElasticloadbalancingLoadBalancerPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElasticloadbalancingLoadBalancerPoliciesOutputReference_Override(e ElasticloadbalancingLoadBalancerPoliciesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.elasticloadbalancingLoadBalancer.ElasticloadbalancingLoadBalancerPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference)SetInstancePorts(val *[]*string) {
	if err := j.validateSetInstancePortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instancePorts",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference)SetLoadBalancerPorts(val *[]*string) {
	if err := j.validateSetLoadBalancerPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerPorts",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference)SetPolicyName(val *string) {
	if err := j.validateSetPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) PutAttributes(value interface{}) {
	if err := e.validatePutAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAttributes",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) ResetAttributes() {
	_jsii_.InvokeVoid(
		e,
		"resetAttributes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) ResetInstancePorts() {
	_jsii_.InvokeVoid(
		e,
		"resetInstancePorts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) ResetLoadBalancerPorts() {
	_jsii_.InvokeVoid(
		e,
		"resetLoadBalancerPorts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) ResetPolicyName() {
	_jsii_.InvokeVoid(
		e,
		"resetPolicyName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) ResetPolicyType() {
	_jsii_.InvokeVoid(
		e,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_ElasticloadbalancingLoadBalancerPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

