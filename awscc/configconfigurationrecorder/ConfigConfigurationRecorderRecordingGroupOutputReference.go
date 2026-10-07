// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigurationrecorder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/configconfigurationrecorder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConfigConfigurationRecorderRecordingGroupOutputReference interface {
	cdktn.ComplexObject
	AllSupported() interface{}
	SetAllSupported(val interface{})
	AllSupportedInput() interface{}
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
	ExclusionByResourceTypes() ConfigConfigurationRecorderRecordingGroupExclusionByResourceTypesOutputReference
	ExclusionByResourceTypesInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeGlobalResourceTypes() interface{}
	SetIncludeGlobalResourceTypes(val interface{})
	IncludeGlobalResourceTypesInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RecordingStrategy() ConfigConfigurationRecorderRecordingGroupRecordingStrategyOutputReference
	RecordingStrategyInput() interface{}
	ResourceTypes() *[]*string
	SetResourceTypes(val *[]*string)
	ResourceTypesInput() *[]*string
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
	PutExclusionByResourceTypes(value *ConfigConfigurationRecorderRecordingGroupExclusionByResourceTypes)
	PutRecordingStrategy(value *ConfigConfigurationRecorderRecordingGroupRecordingStrategy)
	ResetAllSupported()
	ResetExclusionByResourceTypes()
	ResetIncludeGlobalResourceTypes()
	ResetRecordingStrategy()
	ResetResourceTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ConfigConfigurationRecorderRecordingGroupOutputReference
type jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) AllSupported() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) AllSupportedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allSupportedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ExclusionByResourceTypes() ConfigConfigurationRecorderRecordingGroupExclusionByResourceTypesOutputReference {
	var returns ConfigConfigurationRecorderRecordingGroupExclusionByResourceTypesOutputReference
	_jsii_.Get(
		j,
		"exclusionByResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ExclusionByResourceTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"exclusionByResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) IncludeGlobalResourceTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) IncludeGlobalResourceTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) RecordingStrategy() ConfigConfigurationRecorderRecordingGroupRecordingStrategyOutputReference {
	var returns ConfigConfigurationRecorderRecordingGroupRecordingStrategyOutputReference
	_jsii_.Get(
		j,
		"recordingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) RecordingStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewConfigConfigurationRecorderRecordingGroupOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) ConfigConfigurationRecorderRecordingGroupOutputReference {
	_init_.Initialize()

	if err := validateNewConfigConfigurationRecorderRecordingGroupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.configConfigurationRecorder.ConfigConfigurationRecorderRecordingGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewConfigConfigurationRecorderRecordingGroupOutputReference_Override(c ConfigConfigurationRecorderRecordingGroupOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.configConfigurationRecorder.ConfigConfigurationRecorderRecordingGroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetAllSupported(val interface{}) {
	if err := j.validateSetAllSupportedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allSupported",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetIncludeGlobalResourceTypes(val interface{}) {
	if err := j.validateSetIncludeGlobalResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeGlobalResourceTypes",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) PutExclusionByResourceTypes(value *ConfigConfigurationRecorderRecordingGroupExclusionByResourceTypes) {
	if err := c.validatePutExclusionByResourceTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExclusionByResourceTypes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) PutRecordingStrategy(value *ConfigConfigurationRecorderRecordingGroupRecordingStrategy) {
	if err := c.validatePutRecordingStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRecordingStrategy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetAllSupported() {
	_jsii_.InvokeVoid(
		c,
		"resetAllSupported",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetExclusionByResourceTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetExclusionByResourceTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetIncludeGlobalResourceTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeGlobalResourceTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetRecordingStrategy() {
	_jsii_.InvokeVoid(
		c,
		"resetRecordingStrategy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConfigConfigurationRecorderRecordingGroupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

