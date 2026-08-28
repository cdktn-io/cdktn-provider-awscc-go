// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mgnnetworkmigrationdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/mgnnetworkmigrationdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MgnNetworkMigrationDefinitionTargetNetworkOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InboundCidr() *string
	SetInboundCidr(val *string)
	InboundCidrInput() *string
	InspectionCidr() *string
	SetInspectionCidr(val *string)
	InspectionCidrInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OutboundCidr() *string
	SetOutboundCidr(val *string)
	OutboundCidrInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Topology() *string
	SetTopology(val *string)
	TopologyInput() *string
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
	ResetInboundCidr()
	ResetInspectionCidr()
	ResetOutboundCidr()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MgnNetworkMigrationDefinitionTargetNetworkOutputReference
type jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) InboundCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) InboundCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) InspectionCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectionCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) InspectionCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inspectionCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) OutboundCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) OutboundCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outboundCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) Topology() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topology",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) TopologyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyInput",
		&returns,
	)
	return returns
}


func NewMgnNetworkMigrationDefinitionTargetNetworkOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MgnNetworkMigrationDefinitionTargetNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewMgnNetworkMigrationDefinitionTargetNetworkOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mgnNetworkMigrationDefinition.MgnNetworkMigrationDefinitionTargetNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMgnNetworkMigrationDefinitionTargetNetworkOutputReference_Override(m MgnNetworkMigrationDefinitionTargetNetworkOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mgnNetworkMigrationDefinition.MgnNetworkMigrationDefinitionTargetNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference)SetInboundCidr(val *string) {
	if err := j.validateSetInboundCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundCidr",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference)SetInspectionCidr(val *string) {
	if err := j.validateSetInspectionCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectionCidr",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference)SetOutboundCidr(val *string) {
	if err := j.validateSetOutboundCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outboundCidr",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference)SetTopology(val *string) {
	if err := j.validateSetTopologyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topology",
		val,
	)
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) ResetInboundCidr() {
	_jsii_.InvokeVoid(
		m,
		"resetInboundCidr",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) ResetInspectionCidr() {
	_jsii_.InvokeVoid(
		m,
		"resetInspectionCidr",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) ResetOutboundCidr() {
	_jsii_.InvokeVoid(
		m,
		"resetOutboundCidr",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MgnNetworkMigrationDefinitionTargetNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

