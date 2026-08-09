// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/athenaworkgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference interface {
	cdktn.ComplexObject
	AdditionalConfigs() *map[string]*string
	SetAdditionalConfigs(val *map[string]*string)
	AdditionalConfigsInput() *map[string]*string
	Classifications() AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationClassificationsList
	ClassificationsInput() interface{}
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
	CoordinatorDpuSize() *float64
	SetCoordinatorDpuSize(val *float64)
	CoordinatorDpuSizeInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultExecutorDpuSize() *float64
	SetDefaultExecutorDpuSize(val *float64)
	DefaultExecutorDpuSizeInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxConcurrentDpus() *float64
	SetMaxConcurrentDpus(val *float64)
	MaxConcurrentDpusInput() *float64
	SparkProperties() *map[string]*string
	SetSparkProperties(val *map[string]*string)
	SparkPropertiesInput() *map[string]*string
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
	PutClassifications(value interface{})
	ResetAdditionalConfigs()
	ResetClassifications()
	ResetCoordinatorDpuSize()
	ResetDefaultExecutorDpuSize()
	ResetMaxConcurrentDpus()
	ResetSparkProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference
type jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) AdditionalConfigs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) AdditionalConfigsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) Classifications() AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationClassificationsList {
	var returns AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationClassificationsList
	_jsii_.Get(
		j,
		"classifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ClassificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"classificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) CoordinatorDpuSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coordinatorDpuSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) CoordinatorDpuSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"coordinatorDpuSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) DefaultExecutorDpuSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultExecutorDpuSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) DefaultExecutorDpuSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultExecutorDpuSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) MaxConcurrentDpus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentDpus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) MaxConcurrentDpusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentDpusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) SparkProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) SparkPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"sparkPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.athenaWorkGroup.AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference_Override(a AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.athenaWorkGroup.AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetAdditionalConfigs(val *map[string]*string) {
	if err := j.validateSetAdditionalConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalConfigs",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetCoordinatorDpuSize(val *float64) {
	if err := j.validateSetCoordinatorDpuSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coordinatorDpuSize",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetDefaultExecutorDpuSize(val *float64) {
	if err := j.validateSetDefaultExecutorDpuSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultExecutorDpuSize",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetMaxConcurrentDpus(val *float64) {
	if err := j.validateSetMaxConcurrentDpusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentDpus",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetSparkProperties(val *map[string]*string) {
	if err := j.validateSetSparkPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkProperties",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) PutClassifications(value interface{}) {
	if err := a.validatePutClassificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClassifications",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ResetAdditionalConfigs() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalConfigs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ResetClassifications() {
	_jsii_.InvokeVoid(
		a,
		"resetClassifications",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ResetCoordinatorDpuSize() {
	_jsii_.InvokeVoid(
		a,
		"resetCoordinatorDpuSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ResetDefaultExecutorDpuSize() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultExecutorDpuSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ResetMaxConcurrentDpus() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrentDpus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ResetSparkProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetSparkProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationUpdatesEngineConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

