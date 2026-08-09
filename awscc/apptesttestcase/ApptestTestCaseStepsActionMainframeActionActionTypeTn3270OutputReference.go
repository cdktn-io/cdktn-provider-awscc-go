// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apptesttestcase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference interface {
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
	ExportDataSetNames() *[]*string
	SetExportDataSetNames(val *[]*string)
	ExportDataSetNamesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Script() ApptestTestCaseStepsActionMainframeActionActionTypeTn3270ScriptOutputReference
	ScriptInput() interface{}
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
	PutScript(value *ApptestTestCaseStepsActionMainframeActionActionTypeTn3270Script)
	ResetExportDataSetNames()
	ResetScript()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference
type jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) ExportDataSetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportDataSetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) ExportDataSetNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportDataSetNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) Script() ApptestTestCaseStepsActionMainframeActionActionTypeTn3270ScriptOutputReference {
	var returns ApptestTestCaseStepsActionMainframeActionActionTypeTn3270ScriptOutputReference
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) ScriptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference {
	_init_.Initialize()

	if err := validateNewApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apptestTestCase.ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference_Override(a ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apptestTestCase.ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference)SetExportDataSetNames(val *[]*string) {
	if err := j.validateSetExportDataSetNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportDataSetNames",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) PutScript(value *ApptestTestCaseStepsActionMainframeActionActionTypeTn3270Script) {
	if err := a.validatePutScriptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScript",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) ResetExportDataSetNames() {
	_jsii_.InvokeVoid(
		a,
		"resetExportDataSetNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) ResetScript() {
	_jsii_.InvokeVoid(
		a,
		"resetScript",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionMainframeActionActionTypeTn3270OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

