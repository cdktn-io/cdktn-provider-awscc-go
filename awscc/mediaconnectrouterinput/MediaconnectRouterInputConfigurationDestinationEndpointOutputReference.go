// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/mediaconnectrouterinput/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MediaconnectRouterInputConfigurationDestinationEndpointOutputReference interface {
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
	EndpointLifecycleManagement() *string
	SetEndpointLifecycleManagement(val *string)
	EndpointLifecycleManagementInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SourceTransitDecryption() MediaconnectRouterInputConfigurationDestinationEndpointSourceTransitDecryptionOutputReference
	SourceTransitDecryptionInput() interface{}
	SourceTransitDecryptionType() *string
	SetSourceTransitDecryptionType(val *string)
	SourceTransitDecryptionTypeInput() *string
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
	PutSourceTransitDecryption(value *MediaconnectRouterInputConfigurationDestinationEndpointSourceTransitDecryption)
	ResetEndpointLifecycleManagement()
	ResetSourceTransitDecryption()
	ResetSourceTransitDecryptionType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MediaconnectRouterInputConfigurationDestinationEndpointOutputReference
type jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) EndpointLifecycleManagement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointLifecycleManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) EndpointLifecycleManagementInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointLifecycleManagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) SourceTransitDecryption() MediaconnectRouterInputConfigurationDestinationEndpointSourceTransitDecryptionOutputReference {
	var returns MediaconnectRouterInputConfigurationDestinationEndpointSourceTransitDecryptionOutputReference
	_jsii_.Get(
		j,
		"sourceTransitDecryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) SourceTransitDecryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceTransitDecryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) SourceTransitDecryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitDecryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) SourceTransitDecryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceTransitDecryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMediaconnectRouterInputConfigurationDestinationEndpointOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) MediaconnectRouterInputConfigurationDestinationEndpointOutputReference {
	_init_.Initialize()

	if err := validateNewMediaconnectRouterInputConfigurationDestinationEndpointOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterInput.MediaconnectRouterInputConfigurationDestinationEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMediaconnectRouterInputConfigurationDestinationEndpointOutputReference_Override(m MediaconnectRouterInputConfigurationDestinationEndpointOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.mediaconnectRouterInput.MediaconnectRouterInputConfigurationDestinationEndpointOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference)SetEndpointLifecycleManagement(val *string) {
	if err := j.validateSetEndpointLifecycleManagementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointLifecycleManagement",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference)SetSourceTransitDecryptionType(val *string) {
	if err := j.validateSetSourceTransitDecryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceTransitDecryptionType",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) PutSourceTransitDecryption(value *MediaconnectRouterInputConfigurationDestinationEndpointSourceTransitDecryption) {
	if err := m.validatePutSourceTransitDecryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSourceTransitDecryption",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) ResetEndpointLifecycleManagement() {
	_jsii_.InvokeVoid(
		m,
		"resetEndpointLifecycleManagement",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) ResetSourceTransitDecryption() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceTransitDecryption",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) ResetSourceTransitDecryptionType() {
	_jsii_.InvokeVoid(
		m,
		"resetSourceTransitDecryptionType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (m *jsiiProxy_MediaconnectRouterInputConfigurationDestinationEndpointOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

