// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesmailmanagertrafficpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sesmailmanagertrafficpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference interface {
	cdktn.ComplexObject
	BooleanExpression() SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionOutputReference
	BooleanExpressionInput() interface{}
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
	IpExpression() SesMailManagerTrafficPolicyPolicyStatementsConditionsIpExpressionOutputReference
	IpExpressionInput() interface{}
	Ipv6Expression() SesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6ExpressionOutputReference
	Ipv6ExpressionInput() interface{}
	StringExpression() SesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpressionOutputReference
	StringExpressionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TlsExpression() SesMailManagerTrafficPolicyPolicyStatementsConditionsTlsExpressionOutputReference
	TlsExpressionInput() interface{}
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
	PutBooleanExpression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpression)
	PutIpExpression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsIpExpression)
	PutIpv6Expression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6Expression)
	PutStringExpression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpression)
	PutTlsExpression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsTlsExpression)
	ResetBooleanExpression()
	ResetIpExpression()
	ResetIpv6Expression()
	ResetStringExpression()
	ResetTlsExpression()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference
type jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) BooleanExpression() SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionOutputReference {
	var returns SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionOutputReference
	_jsii_.Get(
		j,
		"booleanExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) BooleanExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) IpExpression() SesMailManagerTrafficPolicyPolicyStatementsConditionsIpExpressionOutputReference {
	var returns SesMailManagerTrafficPolicyPolicyStatementsConditionsIpExpressionOutputReference
	_jsii_.Get(
		j,
		"ipExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) IpExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) Ipv6Expression() SesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6ExpressionOutputReference {
	var returns SesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6ExpressionOutputReference
	_jsii_.Get(
		j,
		"ipv6Expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) Ipv6ExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6ExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) StringExpression() SesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpressionOutputReference {
	var returns SesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpressionOutputReference
	_jsii_.Get(
		j,
		"stringExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) StringExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) TlsExpression() SesMailManagerTrafficPolicyPolicyStatementsConditionsTlsExpressionOutputReference {
	var returns SesMailManagerTrafficPolicyPolicyStatementsConditionsTlsExpressionOutputReference
	_jsii_.Get(
		j,
		"tlsExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) TlsExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsExpressionInput",
		&returns,
	)
	return returns
}


func NewSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerTrafficPolicy.SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference_Override(s SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sesMailManagerTrafficPolicy.SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) PutBooleanExpression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpression) {
	if err := s.validatePutBooleanExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putBooleanExpression",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) PutIpExpression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsIpExpression) {
	if err := s.validatePutIpExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIpExpression",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) PutIpv6Expression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6Expression) {
	if err := s.validatePutIpv6ExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIpv6Expression",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) PutStringExpression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpression) {
	if err := s.validatePutStringExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStringExpression",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) PutTlsExpression(value *SesMailManagerTrafficPolicyPolicyStatementsConditionsTlsExpression) {
	if err := s.validatePutTlsExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTlsExpression",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ResetBooleanExpression() {
	_jsii_.InvokeVoid(
		s,
		"resetBooleanExpression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ResetIpExpression() {
	_jsii_.InvokeVoid(
		s,
		"resetIpExpression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ResetIpv6Expression() {
	_jsii_.InvokeVoid(
		s,
		"resetIpv6Expression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ResetStringExpression() {
	_jsii_.InvokeVoid(
		s,
		"resetStringExpression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ResetTlsExpression() {
	_jsii_.InvokeVoid(
		s,
		"resetTlsExpression",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

