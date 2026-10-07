// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsinstrumentationconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/applicationsignalsinstrumentationconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference interface {
	cdktn.ComplexObject
	ClassName() *string
	SetClassName(val *string)
	ClassNameInput() *string
	CodeUnit() *string
	SetCodeUnit(val *string)
	CodeUnitInput() *string
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
	FilePath() *string
	SetFilePath(val *string)
	FilePathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Language() *string
	SetLanguage(val *string)
	LanguageInput() *string
	LineNumber() *float64
	SetLineNumber(val *float64)
	LineNumberInput() *float64
	MethodName() *string
	SetMethodName(val *string)
	MethodNameInput() *string
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
	ResetClassName()
	ResetCodeUnit()
	ResetLineNumber()
	ResetMethodName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference
type jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ClassName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"className",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ClassNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) CodeUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) CodeUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) FilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) FilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) Language() *string {
	var returns *string
	_jsii_.Get(
		j,
		"language",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) LanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) LineNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lineNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) LineNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lineNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) MethodName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) MethodNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.applicationsignalsInstrumentationConfig.ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference_Override(a ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.applicationsignalsInstrumentationConfig.ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetClassName(val *string) {
	if err := j.validateSetClassNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"className",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetCodeUnit(val *string) {
	if err := j.validateSetCodeUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeUnit",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetFilePath(val *string) {
	if err := j.validateSetFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filePath",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetLanguage(val *string) {
	if err := j.validateSetLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"language",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetLineNumber(val *float64) {
	if err := j.validateSetLineNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lineNumber",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetMethodName(val *string) {
	if err := j.validateSetMethodNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"methodName",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ResetClassName() {
	_jsii_.InvokeVoid(
		a,
		"resetClassName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ResetCodeUnit() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeUnit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ResetLineNumber() {
	_jsii_.InvokeVoid(
		a,
		"resetLineNumber",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ResetMethodName() {
	_jsii_.InvokeVoid(
		a,
		"resetMethodName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationsignalsInstrumentationConfigLocationCodeLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

