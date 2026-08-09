// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package frauddetectordetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/frauddetectordetector/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FrauddetectorDetectorEventTypeOutputReference interface {
	cdktn.ComplexObject
	Arn() *string
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
	CreatedTime() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EntityTypes() FrauddetectorDetectorEventTypeEntityTypesList
	EntityTypesInput() interface{}
	EventVariables() FrauddetectorDetectorEventTypeEventVariablesList
	EventVariablesInput() interface{}
	// Experimental.
	Fqn() *string
	Inline() interface{}
	SetInline(val interface{})
	InlineInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Labels() FrauddetectorDetectorEventTypeLabelsList
	LabelsInput() interface{}
	LastUpdatedTime() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Tags() FrauddetectorDetectorEventTypeTagsList
	TagsInput() interface{}
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
	PutEntityTypes(value interface{})
	PutEventVariables(value interface{})
	PutLabels(value interface{})
	PutTags(value interface{})
	ResetDescription()
	ResetEntityTypes()
	ResetEventVariables()
	ResetInline()
	ResetLabels()
	ResetName()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrauddetectorDetectorEventTypeOutputReference
type jsiiProxy_FrauddetectorDetectorEventTypeOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) EntityTypes() FrauddetectorDetectorEventTypeEntityTypesList {
	var returns FrauddetectorDetectorEventTypeEntityTypesList
	_jsii_.Get(
		j,
		"entityTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) EntityTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"entityTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) EventVariables() FrauddetectorDetectorEventTypeEventVariablesList {
	var returns FrauddetectorDetectorEventTypeEventVariablesList
	_jsii_.Get(
		j,
		"eventVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) EventVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) Inline() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) InlineInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inlineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) Labels() FrauddetectorDetectorEventTypeLabelsList {
	var returns FrauddetectorDetectorEventTypeLabelsList
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) LabelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) Tags() FrauddetectorDetectorEventTypeTagsList {
	var returns FrauddetectorDetectorEventTypeTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrauddetectorDetectorEventTypeOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FrauddetectorDetectorEventTypeOutputReference {
	_init_.Initialize()

	if err := validateNewFrauddetectorDetectorEventTypeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FrauddetectorDetectorEventTypeOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.frauddetectorDetector.FrauddetectorDetectorEventTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFrauddetectorDetectorEventTypeOutputReference_Override(f FrauddetectorDetectorEventTypeOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.frauddetectorDetector.FrauddetectorDetectorEventTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference)SetInline(val interface{}) {
	if err := j.validateSetInlineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inline",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) PutEntityTypes(value interface{}) {
	if err := f.validatePutEntityTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putEntityTypes",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) PutEventVariables(value interface{}) {
	if err := f.validatePutEventVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putEventVariables",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) PutLabels(value interface{}) {
	if err := f.validatePutLabelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putLabels",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) PutTags(value interface{}) {
	if err := f.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTags",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		f,
		"resetDescription",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ResetEntityTypes() {
	_jsii_.InvokeVoid(
		f,
		"resetEntityTypes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ResetEventVariables() {
	_jsii_.InvokeVoid(
		f,
		"resetEventVariables",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ResetInline() {
	_jsii_.InvokeVoid(
		f,
		"resetInline",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		f,
		"resetLabels",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		f,
		"resetName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrauddetectorDetectorEventTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

