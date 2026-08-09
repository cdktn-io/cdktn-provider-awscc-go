// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rumappmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/rumappmonitor/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference interface {
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
	DimensionKeys() *map[string]*string
	SetDimensionKeys(val *map[string]*string)
	DimensionKeysInput() *map[string]*string
	EventPattern() *string
	SetEventPattern(val *string)
	EventPatternInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UnitLabel() *string
	SetUnitLabel(val *string)
	UnitLabelInput() *string
	ValueKey() *string
	SetValueKey(val *string)
	ValueKeyInput() *string
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
	ResetDimensionKeys()
	ResetEventPattern()
	ResetName()
	ResetNamespace()
	ResetUnitLabel()
	ResetValueKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference
type jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) DimensionKeys() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) DimensionKeysInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) EventPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) EventPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) UnitLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) UnitLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unitLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ValueKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ValueKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueKeyInput",
		&returns,
	)
	return returns
}


func NewRumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewRumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.rumAppMonitor.RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference_Override(r RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.rumAppMonitor.RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetDimensionKeys(val *map[string]*string) {
	if err := j.validateSetDimensionKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensionKeys",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetEventPattern(val *string) {
	if err := j.validateSetEventPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventPattern",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetUnitLabel(val *string) {
	if err := j.validateSetUnitLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unitLabel",
		val,
	)
}

func (j *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference)SetValueKey(val *string) {
	if err := j.validateSetValueKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueKey",
		val,
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ResetDimensionKeys() {
	_jsii_.InvokeVoid(
		r,
		"resetDimensionKeys",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ResetEventPattern() {
	_jsii_.InvokeVoid(
		r,
		"resetEventPattern",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		r,
		"resetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		r,
		"resetNamespace",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ResetUnitLabel() {
	_jsii_.InvokeVoid(
		r,
		"resetUnitLabel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ResetValueKey() {
	_jsii_.InvokeVoid(
		r,
		"resetValueKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

