// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamprefixlistresolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2ipamprefixlistresolver/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2IpamPrefixListResolverRulesConditionsOutputReference interface {
	cdktn.ComplexObject
	Cidr() *string
	SetCidr(val *string)
	CidrInput() *string
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
	IpamPoolId() *string
	SetIpamPoolId(val *string)
	IpamPoolIdInput() *string
	Operation() *string
	SetOperation(val *string)
	OperationInput() *string
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResourceOwner() *string
	SetResourceOwner(val *string)
	ResourceOwnerInput() *string
	ResourceRegion() *string
	SetResourceRegion(val *string)
	ResourceRegionInput() *string
	ResourceTag() Ec2IpamPrefixListResolverRulesConditionsResourceTagOutputReference
	ResourceTagInput() interface{}
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
	PutResourceTag(value *Ec2IpamPrefixListResolverRulesConditionsResourceTag)
	ResetCidr()
	ResetIpamPoolId()
	ResetOperation()
	ResetResourceId()
	ResetResourceOwner()
	ResetResourceRegion()
	ResetResourceTag()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2IpamPrefixListResolverRulesConditionsOutputReference
type jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) IpamPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) OperationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResourceOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResourceOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResourceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResourceTag() Ec2IpamPrefixListResolverRulesConditionsResourceTagOutputReference {
	var returns Ec2IpamPrefixListResolverRulesConditionsResourceTagOutputReference
	_jsii_.Get(
		j,
		"resourceTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResourceTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2IpamPrefixListResolverRulesConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2IpamPrefixListResolverRulesConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2IpamPrefixListResolverRulesConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2IpamPrefixListResolver.Ec2IpamPrefixListResolverRulesConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2IpamPrefixListResolverRulesConditionsOutputReference_Override(e Ec2IpamPrefixListResolverRulesConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2IpamPrefixListResolver.Ec2IpamPrefixListResolverRulesConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetCidr(val *string) {
	if err := j.validateSetCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidr",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetIpamPoolId(val *string) {
	if err := j.validateSetIpamPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipamPoolId",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetOperation(val *string) {
	if err := j.validateSetOperationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operation",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetResourceOwner(val *string) {
	if err := j.validateSetResourceOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceOwner",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetResourceRegion(val *string) {
	if err := j.validateSetResourceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceRegion",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) PutResourceTag(value *Ec2IpamPrefixListResolverRulesConditionsResourceTag) {
	if err := e.validatePutResourceTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putResourceTag",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResetCidr() {
	_jsii_.InvokeVoid(
		e,
		"resetCidr",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResetIpamPoolId() {
	_jsii_.InvokeVoid(
		e,
		"resetIpamPoolId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResetOperation() {
	_jsii_.InvokeVoid(
		e,
		"resetOperation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResetResourceOwner() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceOwner",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResetResourceRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ResetResourceTag() {
	_jsii_.InvokeVoid(
		e,
		"resetResourceTag",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2IpamPrefixListResolverRulesConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

