// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisesignalcatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotfleetwisesignalcatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotfleetwiseSignalCatalogNodesOutputReference interface {
	cdktn.ComplexObject
	Actuator() IotfleetwiseSignalCatalogNodesActuatorOutputReference
	ActuatorInput() interface{}
	Attribute() IotfleetwiseSignalCatalogNodesAttributeOutputReference
	AttributeInput() interface{}
	Branch() IotfleetwiseSignalCatalogNodesBranchOutputReference
	BranchInput() interface{}
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
	Sensor() IotfleetwiseSignalCatalogNodesSensorOutputReference
	SensorInput() interface{}
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
	PutActuator(value *IotfleetwiseSignalCatalogNodesActuator)
	PutAttribute(value *IotfleetwiseSignalCatalogNodesAttribute)
	PutBranch(value *IotfleetwiseSignalCatalogNodesBranch)
	PutSensor(value *IotfleetwiseSignalCatalogNodesSensor)
	ResetActuator()
	ResetAttribute()
	ResetBranch()
	ResetSensor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotfleetwiseSignalCatalogNodesOutputReference
type jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) Actuator() IotfleetwiseSignalCatalogNodesActuatorOutputReference {
	var returns IotfleetwiseSignalCatalogNodesActuatorOutputReference
	_jsii_.Get(
		j,
		"actuator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) ActuatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actuatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) Attribute() IotfleetwiseSignalCatalogNodesAttributeOutputReference {
	var returns IotfleetwiseSignalCatalogNodesAttributeOutputReference
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) AttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) Branch() IotfleetwiseSignalCatalogNodesBranchOutputReference {
	var returns IotfleetwiseSignalCatalogNodesBranchOutputReference
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) BranchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) Sensor() IotfleetwiseSignalCatalogNodesSensorOutputReference {
	var returns IotfleetwiseSignalCatalogNodesSensorOutputReference
	_jsii_.Get(
		j,
		"sensor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) SensorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sensorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotfleetwiseSignalCatalogNodesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IotfleetwiseSignalCatalogNodesOutputReference {
	_init_.Initialize()

	if err := validateNewIotfleetwiseSignalCatalogNodesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseSignalCatalog.IotfleetwiseSignalCatalogNodesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIotfleetwiseSignalCatalogNodesOutputReference_Override(i IotfleetwiseSignalCatalogNodesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotfleetwiseSignalCatalog.IotfleetwiseSignalCatalogNodesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) PutActuator(value *IotfleetwiseSignalCatalogNodesActuator) {
	if err := i.validatePutActuatorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putActuator",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) PutAttribute(value *IotfleetwiseSignalCatalogNodesAttribute) {
	if err := i.validatePutAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAttribute",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) PutBranch(value *IotfleetwiseSignalCatalogNodesBranch) {
	if err := i.validatePutBranchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putBranch",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) PutSensor(value *IotfleetwiseSignalCatalogNodesSensor) {
	if err := i.validatePutSensorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSensor",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) ResetActuator() {
	_jsii_.InvokeVoid(
		i,
		"resetActuator",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) ResetAttribute() {
	_jsii_.InvokeVoid(
		i,
		"resetAttribute",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) ResetBranch() {
	_jsii_.InvokeVoid(
		i,
		"resetBranch",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) ResetSensor() {
	_jsii_.InvokeVoid(
		i,
		"resetSensor",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotfleetwiseSignalCatalogNodesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

