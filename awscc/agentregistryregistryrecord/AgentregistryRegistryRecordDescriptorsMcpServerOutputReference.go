// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryregistryrecord

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/agentregistryregistryrecord/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AgentregistryRegistryRecordDescriptorsMcpServerOutputReference interface {
	cdktn.ComplexObject
	AdditionalData() AgentregistryRegistryRecordDescriptorsMcpServerAdditionalDataOutputReference
	AdditionalDataInput() interface{}
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
	Data() *string
	SetData(val *string)
	DataInput() *string
	DataSchemaVersion() *string
	SetDataSchemaVersion(val *string)
	DataSchemaVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Source() AgentregistryRegistryRecordDescriptorsMcpServerSourceOutputReference
	SourceInput() interface{}
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
	PutAdditionalData(value *AgentregistryRegistryRecordDescriptorsMcpServerAdditionalData)
	PutSource(value *AgentregistryRegistryRecordDescriptorsMcpServerSource)
	ResetAdditionalData()
	ResetData()
	ResetDataSchemaVersion()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AgentregistryRegistryRecordDescriptorsMcpServerOutputReference
type jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) AdditionalData() AgentregistryRegistryRecordDescriptorsMcpServerAdditionalDataOutputReference {
	var returns AgentregistryRegistryRecordDescriptorsMcpServerAdditionalDataOutputReference
	_jsii_.Get(
		j,
		"additionalData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) AdditionalDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) Data() *string {
	var returns *string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) DataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) DataSchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSchemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) DataSchemaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSchemaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) Source() AgentregistryRegistryRecordDescriptorsMcpServerSourceOutputReference {
	var returns AgentregistryRegistryRecordDescriptorsMcpServerSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) SourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAgentregistryRegistryRecordDescriptorsMcpServerOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AgentregistryRegistryRecordDescriptorsMcpServerOutputReference {
	_init_.Initialize()

	if err := validateNewAgentregistryRegistryRecordDescriptorsMcpServerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.agentregistryRegistryRecord.AgentregistryRegistryRecordDescriptorsMcpServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAgentregistryRegistryRecordDescriptorsMcpServerOutputReference_Override(a AgentregistryRegistryRecordDescriptorsMcpServerOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.agentregistryRegistryRecord.AgentregistryRegistryRecordDescriptorsMcpServerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference)SetData(val *string) {
	if err := j.validateSetDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference)SetDataSchemaVersion(val *string) {
	if err := j.validateSetDataSchemaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSchemaVersion",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) PutAdditionalData(value *AgentregistryRegistryRecordDescriptorsMcpServerAdditionalData) {
	if err := a.validatePutAdditionalDataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdditionalData",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) PutSource(value *AgentregistryRegistryRecordDescriptorsMcpServerSource) {
	if err := a.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) ResetAdditionalData() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) ResetData() {
	_jsii_.InvokeVoid(
		a,
		"resetData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) ResetDataSchemaVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDataSchemaVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		a,
		"resetSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AgentregistryRegistryRecordDescriptorsMcpServerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

