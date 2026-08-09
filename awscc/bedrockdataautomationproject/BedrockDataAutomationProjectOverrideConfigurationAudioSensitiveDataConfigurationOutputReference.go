// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdataautomationproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bedrockdataautomationproject/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference interface {
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
	DetectionMode() *string
	SetDetectionMode(val *string)
	DetectionModeInput() *string
	DetectionScope() *[]*string
	SetDetectionScope(val *[]*string)
	DetectionScopeInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PiiEntitiesConfiguration() BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference
	PiiEntitiesConfigurationInput() interface{}
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
	PutPiiEntitiesConfiguration(value *BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationPiiEntitiesConfiguration)
	ResetDetectionMode()
	ResetDetectionScope()
	ResetPiiEntitiesConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference
type jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) DetectionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) DetectionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"detectionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) DetectionScope() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"detectionScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) DetectionScopeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"detectionScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) PiiEntitiesConfiguration() BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference {
	var returns BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationPiiEntitiesConfigurationOutputReference
	_jsii_.Get(
		j,
		"piiEntitiesConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) PiiEntitiesConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"piiEntitiesConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewBedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference_Override(b BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bedrockDataAutomationProject.BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference)SetDetectionMode(val *string) {
	if err := j.validateSetDetectionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectionMode",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference)SetDetectionScope(val *[]*string) {
	if err := j.validateSetDetectionScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectionScope",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) PutPiiEntitiesConfiguration(value *BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationPiiEntitiesConfiguration) {
	if err := b.validatePutPiiEntitiesConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPiiEntitiesConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) ResetDetectionMode() {
	_jsii_.InvokeVoid(
		b,
		"resetDetectionMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) ResetDetectionScope() {
	_jsii_.InvokeVoid(
		b,
		"resetDetectionScope",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) ResetPiiEntitiesConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetPiiEntitiesConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockDataAutomationProjectOverrideConfigurationAudioSensitiveDataConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

