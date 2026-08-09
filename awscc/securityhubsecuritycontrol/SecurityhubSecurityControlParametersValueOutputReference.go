// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubsecuritycontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/securityhubsecuritycontrol/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubSecurityControlParametersValueOutputReference interface {
	cdktn.ComplexObject
	Boolean() interface{}
	SetBoolean(val interface{})
	BooleanInput() interface{}
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
	Double() *float64
	SetDouble(val *float64)
	DoubleInput() *float64
	Enum() *string
	SetEnum(val *string)
	EnumInput() *string
	EnumList() *[]*string
	SetEnumList(val *[]*string)
	EnumListInput() *[]*string
	// Experimental.
	Fqn() *string
	Integer() *float64
	SetInteger(val *float64)
	IntegerInput() *float64
	IntegerList() *[]*float64
	SetIntegerList(val *[]*float64)
	IntegerListInput() *[]*float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	String() *string
	SetString(val *string)
	StringInput() *string
	StringList() *[]*string
	SetStringList(val *[]*string)
	StringListInput() *[]*string
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
	ResetBoolean()
	ResetDouble()
	ResetEnum()
	ResetEnumList()
	ResetInteger()
	ResetIntegerList()
	ResetString()
	ResetStringList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityhubSecurityControlParametersValueOutputReference
type jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) Boolean() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolean",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) BooleanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"booleanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) Double() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"double",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) DoubleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"doubleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) Enum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) EnumInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) EnumList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) EnumListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enumListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) Integer() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"integer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) IntegerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"integerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) IntegerList() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"integerList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) IntegerListInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"integerListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) String() *string {
	var returns *string
	_jsii_.Get(
		j,
		"string",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) StringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) StringList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) StringListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"stringListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityhubSecurityControlParametersValueOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SecurityhubSecurityControlParametersValueOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityhubSecurityControlParametersValueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.securityhubSecurityControl.SecurityhubSecurityControlParametersValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityhubSecurityControlParametersValueOutputReference_Override(s SecurityhubSecurityControlParametersValueOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.securityhubSecurityControl.SecurityhubSecurityControlParametersValueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetBoolean(val interface{}) {
	if err := j.validateSetBooleanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boolean",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetDouble(val *float64) {
	if err := j.validateSetDoubleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"double",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetEnum(val *string) {
	if err := j.validateSetEnumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enum",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetEnumList(val *[]*string) {
	if err := j.validateSetEnumListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enumList",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetInteger(val *float64) {
	if err := j.validateSetIntegerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integer",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetIntegerList(val *[]*float64) {
	if err := j.validateSetIntegerListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integerList",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetString(val *string) {
	if err := j.validateSetStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"string",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetStringList(val *[]*string) {
	if err := j.validateSetStringListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringList",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ResetBoolean() {
	_jsii_.InvokeVoid(
		s,
		"resetBoolean",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ResetDouble() {
	_jsii_.InvokeVoid(
		s,
		"resetDouble",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ResetEnum() {
	_jsii_.InvokeVoid(
		s,
		"resetEnum",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ResetEnumList() {
	_jsii_.InvokeVoid(
		s,
		"resetEnumList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ResetInteger() {
	_jsii_.InvokeVoid(
		s,
		"resetInteger",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ResetIntegerList() {
	_jsii_.InvokeVoid(
		s,
		"resetIntegerList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ResetString() {
	_jsii_.InvokeVoid(
		s,
		"resetString",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ResetStringList() {
	_jsii_.InvokeVoid(
		s,
		"resetStringList",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityhubSecurityControlParametersValueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

