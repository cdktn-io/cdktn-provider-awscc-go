// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccsesmailmanagertrafficpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccsesmailmanagertrafficpolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference interface {
	cdktn.ComplexObject
	BooleanExpression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionOutputReference
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
	InternalValue() *DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditions
	SetInternalValue(val *DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditions)
	IpExpression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsIpExpressionOutputReference
	Ipv6Expression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6ExpressionOutputReference
	StringExpression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpressionOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TlsExpression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsTlsExpressionOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference
type jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) BooleanExpression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionOutputReference {
	var returns DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsBooleanExpressionOutputReference
	_jsii_.Get(
		j,
		"booleanExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) InternalValue() *DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditions {
	var returns *DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) IpExpression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsIpExpressionOutputReference {
	var returns DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsIpExpressionOutputReference
	_jsii_.Get(
		j,
		"ipExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) Ipv6Expression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6ExpressionOutputReference {
	var returns DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsIpv6ExpressionOutputReference
	_jsii_.Get(
		j,
		"ipv6Expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) StringExpression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpressionOutputReference {
	var returns DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpressionOutputReference
	_jsii_.Get(
		j,
		"stringExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) TlsExpression() DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsTlsExpressionOutputReference {
	var returns DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsTlsExpressionOutputReference
	_jsii_.Get(
		j,
		"tlsExpression",
		&returns,
	)
	return returns
}


func NewDataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSesMailManagerTrafficPolicy.DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference_Override(d DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSesMailManagerTrafficPolicy.DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetInternalValue(val *DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSesMailManagerTrafficPolicyPolicyStatementsConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

