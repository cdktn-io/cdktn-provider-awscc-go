// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationseteventdestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sesconfigurationseteventdestination/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesConfigurationSetEventDestinationEventDestinationOutputReference interface {
	cdktn.ComplexObject
	CloudwatchDestination() SesConfigurationSetEventDestinationEventDestinationCloudwatchDestinationOutputReference
	CloudwatchDestinationInput() interface{}
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EventBridgeDestination() SesConfigurationSetEventDestinationEventDestinationEventBridgeDestinationOutputReference
	EventBridgeDestinationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KinesisFirehoseDestination() SesConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationOutputReference
	KinesisFirehoseDestinationInput() interface{}
	MatchingEventTypes() *[]*string
	SetMatchingEventTypes(val *[]*string)
	MatchingEventTypesInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	SnsDestination() SesConfigurationSetEventDestinationEventDestinationSnsDestinationOutputReference
	SnsDestinationInput() interface{}
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
	PutCloudwatchDestination(value *SesConfigurationSetEventDestinationEventDestinationCloudwatchDestination)
	PutEventBridgeDestination(value *SesConfigurationSetEventDestinationEventDestinationEventBridgeDestination)
	PutKinesisFirehoseDestination(value *SesConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination)
	PutSnsDestination(value *SesConfigurationSetEventDestinationEventDestinationSnsDestination)
	ResetCloudwatchDestination()
	ResetEnabled()
	ResetEventBridgeDestination()
	ResetKinesisFirehoseDestination()
	ResetName()
	ResetSnsDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SesConfigurationSetEventDestinationEventDestinationOutputReference
type jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) CloudwatchDestination() SesConfigurationSetEventDestinationEventDestinationCloudwatchDestinationOutputReference {
	var returns SesConfigurationSetEventDestinationEventDestinationCloudwatchDestinationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) CloudwatchDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) EventBridgeDestination() SesConfigurationSetEventDestinationEventDestinationEventBridgeDestinationOutputReference {
	var returns SesConfigurationSetEventDestinationEventDestinationEventBridgeDestinationOutputReference
	_jsii_.Get(
		j,
		"eventBridgeDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) EventBridgeDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventBridgeDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) KinesisFirehoseDestination() SesConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationOutputReference {
	var returns SesConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestinationOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) KinesisFirehoseDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisFirehoseDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) MatchingEventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) MatchingEventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingEventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) SnsDestination() SesConfigurationSetEventDestinationEventDestinationSnsDestinationOutputReference {
	var returns SesConfigurationSetEventDestinationEventDestinationSnsDestinationOutputReference
	_jsii_.Get(
		j,
		"snsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) SnsDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSesConfigurationSetEventDestinationEventDestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SesConfigurationSetEventDestinationEventDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewSesConfigurationSetEventDestinationEventDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sesConfigurationSetEventDestination.SesConfigurationSetEventDestinationEventDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSesConfigurationSetEventDestinationEventDestinationOutputReference_Override(s SesConfigurationSetEventDestinationEventDestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sesConfigurationSetEventDestination.SesConfigurationSetEventDestinationEventDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference)SetMatchingEventTypes(val *[]*string) {
	if err := j.validateSetMatchingEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingEventTypes",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) PutCloudwatchDestination(value *SesConfigurationSetEventDestinationEventDestinationCloudwatchDestination) {
	if err := s.validatePutCloudwatchDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCloudwatchDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) PutEventBridgeDestination(value *SesConfigurationSetEventDestinationEventDestinationEventBridgeDestination) {
	if err := s.validatePutEventBridgeDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putEventBridgeDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) PutKinesisFirehoseDestination(value *SesConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination) {
	if err := s.validatePutKinesisFirehoseDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putKinesisFirehoseDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) PutSnsDestination(value *SesConfigurationSetEventDestinationEventDestinationSnsDestination) {
	if err := s.validatePutSnsDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSnsDestination",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ResetCloudwatchDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatchDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ResetEventBridgeDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetEventBridgeDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ResetKinesisFirehoseDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetKinesisFirehoseDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ResetSnsDestination() {
	_jsii_.InvokeVoid(
		s,
		"resetSnsDestination",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SesConfigurationSetEventDestinationEventDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

