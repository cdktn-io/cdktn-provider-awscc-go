// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotanalyticsdataset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotanalyticsDatasetActionsContainerActionVariablesOutputReference interface {
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
	DatasetContentVersionValue() IotanalyticsDatasetActionsContainerActionVariablesDatasetContentVersionValueOutputReference
	DatasetContentVersionValueInput() interface{}
	DoubleValue() *float64
	SetDoubleValue(val *float64)
	DoubleValueInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutputFileUriValue() IotanalyticsDatasetActionsContainerActionVariablesOutputFileUriValueOutputReference
	OutputFileUriValueInput() interface{}
	StringValue() *string
	SetStringValue(val *string)
	StringValueInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VariableName() *string
	SetVariableName(val *string)
	VariableNameInput() *string
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
	PutDatasetContentVersionValue(value *IotanalyticsDatasetActionsContainerActionVariablesDatasetContentVersionValue)
	PutOutputFileUriValue(value *IotanalyticsDatasetActionsContainerActionVariablesOutputFileUriValue)
	ResetDatasetContentVersionValue()
	ResetDoubleValue()
	ResetOutputFileUriValue()
	ResetStringValue()
	ResetVariableName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotanalyticsDatasetActionsContainerActionVariablesOutputReference
type jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) DatasetContentVersionValue() IotanalyticsDatasetActionsContainerActionVariablesDatasetContentVersionValueOutputReference {
	var returns IotanalyticsDatasetActionsContainerActionVariablesDatasetContentVersionValueOutputReference
	_jsii_.Get(
		j,
		"datasetContentVersionValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) DatasetContentVersionValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"datasetContentVersionValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) DoubleValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"doubleValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) DoubleValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"doubleValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) OutputFileUriValue() IotanalyticsDatasetActionsContainerActionVariablesOutputFileUriValueOutputReference {
	var returns IotanalyticsDatasetActionsContainerActionVariablesOutputFileUriValueOutputReference
	_jsii_.Get(
		j,
		"outputFileUriValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) OutputFileUriValueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputFileUriValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) VariableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) VariableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variableNameInput",
		&returns,
	)
	return returns
}


func NewIotanalyticsDatasetActionsContainerActionVariablesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotanalyticsDatasetActionsContainerActionVariablesOutputReference {
	_init_.Initialize()

	if err := validateNewIotanalyticsDatasetActionsContainerActionVariablesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotanalyticsDataset.IotanalyticsDatasetActionsContainerActionVariablesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotanalyticsDatasetActionsContainerActionVariablesOutputReference_Override(i IotanalyticsDatasetActionsContainerActionVariablesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotanalyticsDataset.IotanalyticsDatasetActionsContainerActionVariablesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference)SetDoubleValue(val *float64) {
	if err := j.validateSetDoubleValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"doubleValue",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference)SetVariableName(val *string) {
	if err := j.validateSetVariableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variableName",
		val,
	)
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) PutDatasetContentVersionValue(value *IotanalyticsDatasetActionsContainerActionVariablesDatasetContentVersionValue) {
	if err := i.validatePutDatasetContentVersionValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDatasetContentVersionValue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) PutOutputFileUriValue(value *IotanalyticsDatasetActionsContainerActionVariablesOutputFileUriValue) {
	if err := i.validatePutOutputFileUriValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOutputFileUriValue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) ResetDatasetContentVersionValue() {
	_jsii_.InvokeVoid(
		i,
		"resetDatasetContentVersionValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) ResetDoubleValue() {
	_jsii_.InvokeVoid(
		i,
		"resetDoubleValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) ResetOutputFileUriValue() {
	_jsii_.InvokeVoid(
		i,
		"resetOutputFileUriValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		i,
		"resetStringValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) ResetVariableName() {
	_jsii_.InvokeVoid(
		i,
		"resetVariableName",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotanalyticsDatasetActionsContainerActionVariablesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

