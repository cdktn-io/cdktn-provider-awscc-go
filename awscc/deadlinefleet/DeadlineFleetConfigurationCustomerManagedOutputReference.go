// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinefleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/deadlinefleet/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineFleetConfigurationCustomerManagedOutputReference interface {
	cdktn.ComplexObject
	AutoScalingConfiguration() DeadlineFleetConfigurationCustomerManagedAutoScalingConfigurationOutputReference
	AutoScalingConfigurationInput() interface{}
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
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	StorageProfileId() *string
	SetStorageProfileId(val *string)
	StorageProfileIdInput() *string
	TagPropagationMode() *string
	SetTagPropagationMode(val *string)
	TagPropagationModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkerCapabilities() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference
	WorkerCapabilitiesInput() interface{}
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
	PutAutoScalingConfiguration(value *DeadlineFleetConfigurationCustomerManagedAutoScalingConfiguration)
	PutWorkerCapabilities(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilities)
	ResetAutoScalingConfiguration()
	ResetMode()
	ResetStorageProfileId()
	ResetTagPropagationMode()
	ResetWorkerCapabilities()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DeadlineFleetConfigurationCustomerManagedOutputReference
type jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) AutoScalingConfiguration() DeadlineFleetConfigurationCustomerManagedAutoScalingConfigurationOutputReference {
	var returns DeadlineFleetConfigurationCustomerManagedAutoScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"autoScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) AutoScalingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) StorageProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) StorageProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) TagPropagationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPropagationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) TagPropagationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPropagationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) WorkerCapabilities() DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference {
	var returns DeadlineFleetConfigurationCustomerManagedWorkerCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"workerCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) WorkerCapabilitiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workerCapabilitiesInput",
		&returns,
	)
	return returns
}


func NewDeadlineFleetConfigurationCustomerManagedOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DeadlineFleetConfigurationCustomerManagedOutputReference {
	_init_.Initialize()

	if err := validateNewDeadlineFleetConfigurationCustomerManagedOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationCustomerManagedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDeadlineFleetConfigurationCustomerManagedOutputReference_Override(d DeadlineFleetConfigurationCustomerManagedOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.deadlineFleet.DeadlineFleetConfigurationCustomerManagedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference)SetStorageProfileId(val *string) {
	if err := j.validateSetStorageProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageProfileId",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference)SetTagPropagationMode(val *string) {
	if err := j.validateSetTagPropagationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagPropagationMode",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) PutAutoScalingConfiguration(value *DeadlineFleetConfigurationCustomerManagedAutoScalingConfiguration) {
	if err := d.validatePutAutoScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutoScalingConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) PutWorkerCapabilities(value *DeadlineFleetConfigurationCustomerManagedWorkerCapabilities) {
	if err := d.validatePutWorkerCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkerCapabilities",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ResetAutoScalingConfiguration() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoScalingConfiguration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		d,
		"resetMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ResetStorageProfileId() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageProfileId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ResetTagPropagationMode() {
	_jsii_.InvokeVoid(
		d,
		"resetTagPropagationMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ResetWorkerCapabilities() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkerCapabilities",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeadlineFleetConfigurationCustomerManagedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

