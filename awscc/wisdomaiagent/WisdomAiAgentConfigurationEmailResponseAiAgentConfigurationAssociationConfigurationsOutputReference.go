// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wisdomaiagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/wisdomaiagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference interface {
	cdktn.ComplexObject
	AssociationConfigurationData() WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference
	AssociationConfigurationDataInput() interface{}
	AssociationId() *string
	SetAssociationId(val *string)
	AssociationIdInput() *string
	AssociationType() *string
	SetAssociationType(val *string)
	AssociationTypeInput() *string
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
	PutAssociationConfigurationData(value *WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsAssociationConfigurationData)
	ResetAssociationConfigurationData()
	ResetAssociationId()
	ResetAssociationType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference
type jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) AssociationConfigurationData() WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference {
	var returns WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataOutputReference
	_jsii_.Get(
		j,
		"associationConfigurationData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) AssociationConfigurationDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associationConfigurationDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) AssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) AssociationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) AssociationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) AssociationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewWisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference_Override(w WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.wisdomAiAgent.WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference)SetAssociationId(val *string) {
	if err := j.validateSetAssociationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associationId",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference)SetAssociationType(val *string) {
	if err := j.validateSetAssociationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associationType",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) PutAssociationConfigurationData(value *WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsAssociationConfigurationData) {
	if err := w.validatePutAssociationConfigurationDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAssociationConfigurationData",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) ResetAssociationConfigurationData() {
	_jsii_.InvokeVoid(
		w,
		"resetAssociationConfigurationData",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) ResetAssociationId() {
	_jsii_.InvokeVoid(
		w,
		"resetAssociationId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) ResetAssociationType() {
	_jsii_.InvokeVoid(
		w,
		"resetAssociationType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := w.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WisdomAiAgentConfigurationEmailResponseAiAgentConfigurationAssociationConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

