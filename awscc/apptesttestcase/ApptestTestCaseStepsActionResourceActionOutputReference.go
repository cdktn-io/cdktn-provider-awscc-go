// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apptesttestcase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/apptesttestcase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApptestTestCaseStepsActionResourceActionOutputReference interface {
	cdktn.ComplexObject
	CloudformationAction() ApptestTestCaseStepsActionResourceActionCloudformationActionOutputReference
	CloudformationActionInput() interface{}
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
	M2ManagedApplicationAction() ApptestTestCaseStepsActionResourceActionM2ManagedApplicationActionOutputReference
	M2ManagedApplicationActionInput() interface{}
	M2NonManagedApplicationAction() ApptestTestCaseStepsActionResourceActionM2NonManagedApplicationActionOutputReference
	M2NonManagedApplicationActionInput() interface{}
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
	PutCloudformationAction(value *ApptestTestCaseStepsActionResourceActionCloudformationAction)
	PutM2ManagedApplicationAction(value *ApptestTestCaseStepsActionResourceActionM2ManagedApplicationAction)
	PutM2NonManagedApplicationAction(value *ApptestTestCaseStepsActionResourceActionM2NonManagedApplicationAction)
	ResetCloudformationAction()
	ResetM2ManagedApplicationAction()
	ResetM2NonManagedApplicationAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApptestTestCaseStepsActionResourceActionOutputReference
type jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) CloudformationAction() ApptestTestCaseStepsActionResourceActionCloudformationActionOutputReference {
	var returns ApptestTestCaseStepsActionResourceActionCloudformationActionOutputReference
	_jsii_.Get(
		j,
		"cloudformationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) CloudformationActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudformationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) M2ManagedApplicationAction() ApptestTestCaseStepsActionResourceActionM2ManagedApplicationActionOutputReference {
	var returns ApptestTestCaseStepsActionResourceActionM2ManagedApplicationActionOutputReference
	_jsii_.Get(
		j,
		"m2ManagedApplicationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) M2ManagedApplicationActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"m2ManagedApplicationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) M2NonManagedApplicationAction() ApptestTestCaseStepsActionResourceActionM2NonManagedApplicationActionOutputReference {
	var returns ApptestTestCaseStepsActionResourceActionM2NonManagedApplicationActionOutputReference
	_jsii_.Get(
		j,
		"m2NonManagedApplicationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) M2NonManagedApplicationActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"m2NonManagedApplicationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApptestTestCaseStepsActionResourceActionOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ApptestTestCaseStepsActionResourceActionOutputReference {
	_init_.Initialize()

	if err := validateNewApptestTestCaseStepsActionResourceActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.apptestTestCase.ApptestTestCaseStepsActionResourceActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApptestTestCaseStepsActionResourceActionOutputReference_Override(a ApptestTestCaseStepsActionResourceActionOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.apptestTestCase.ApptestTestCaseStepsActionResourceActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) PutCloudformationAction(value *ApptestTestCaseStepsActionResourceActionCloudformationAction) {
	if err := a.validatePutCloudformationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudformationAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) PutM2ManagedApplicationAction(value *ApptestTestCaseStepsActionResourceActionM2ManagedApplicationAction) {
	if err := a.validatePutM2ManagedApplicationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putM2ManagedApplicationAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) PutM2NonManagedApplicationAction(value *ApptestTestCaseStepsActionResourceActionM2NonManagedApplicationAction) {
	if err := a.validatePutM2NonManagedApplicationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putM2NonManagedApplicationAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) ResetCloudformationAction() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudformationAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) ResetM2ManagedApplicationAction() {
	_jsii_.InvokeVoid(
		a,
		"resetM2ManagedApplicationAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) ResetM2NonManagedApplicationAction() {
	_jsii_.InvokeVoid(
		a,
		"resetM2NonManagedApplicationAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApptestTestCaseStepsActionResourceActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

