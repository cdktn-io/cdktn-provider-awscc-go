// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sagemakerendpointconfig/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerEndpointConfigShadowProductionVariantsOutputReference interface {
	cdktn.ComplexObject
	CapacityReservationConfig() SagemakerEndpointConfigShadowProductionVariantsCapacityReservationConfigOutputReference
	CapacityReservationConfigInput() interface{}
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
	ContainerStartupHealthCheckTimeoutInSeconds() *float64
	SetContainerStartupHealthCheckTimeoutInSeconds(val *float64)
	ContainerStartupHealthCheckTimeoutInSecondsInput() *float64
	CoreDumpConfig() SagemakerEndpointConfigShadowProductionVariantsCoreDumpConfigOutputReference
	CoreDumpConfigInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableSsmAccess() interface{}
	SetEnableSsmAccess(val interface{})
	EnableSsmAccessInput() interface{}
	// Experimental.
	Fqn() *string
	InferenceAmiVersion() *string
	SetInferenceAmiVersion(val *string)
	InferenceAmiVersionInput() *string
	InitialInstanceCount() *float64
	SetInitialInstanceCount(val *float64)
	InitialInstanceCountInput() *float64
	InitialVariantWeight() *float64
	SetInitialVariantWeight(val *float64)
	InitialVariantWeightInput() *float64
	InstancePools() SagemakerEndpointConfigShadowProductionVariantsInstancePoolsList
	InstancePoolsInput() interface{}
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManagedInstanceScaling() SagemakerEndpointConfigShadowProductionVariantsManagedInstanceScalingOutputReference
	ManagedInstanceScalingInput() interface{}
	ModelDataDownloadTimeoutInSeconds() *float64
	SetModelDataDownloadTimeoutInSeconds(val *float64)
	ModelDataDownloadTimeoutInSecondsInput() *float64
	ModelName() *string
	SetModelName(val *string)
	ModelNameInput() *string
	RoutingConfig() SagemakerEndpointConfigShadowProductionVariantsRoutingConfigOutputReference
	RoutingConfigInput() interface{}
	ServerlessConfig() SagemakerEndpointConfigShadowProductionVariantsServerlessConfigOutputReference
	ServerlessConfigInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VariantInstanceProvisionTimeoutInSeconds() *float64
	SetVariantInstanceProvisionTimeoutInSeconds(val *float64)
	VariantInstanceProvisionTimeoutInSecondsInput() *float64
	VariantName() *string
	SetVariantName(val *string)
	VariantNameInput() *string
	VolumeSizeInGb() *float64
	SetVolumeSizeInGb(val *float64)
	VolumeSizeInGbInput() *float64
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
	PutCapacityReservationConfig(value *SagemakerEndpointConfigShadowProductionVariantsCapacityReservationConfig)
	PutCoreDumpConfig(value *SagemakerEndpointConfigShadowProductionVariantsCoreDumpConfig)
	PutInstancePools(value interface{})
	PutManagedInstanceScaling(value *SagemakerEndpointConfigShadowProductionVariantsManagedInstanceScaling)
	PutRoutingConfig(value *SagemakerEndpointConfigShadowProductionVariantsRoutingConfig)
	PutServerlessConfig(value *SagemakerEndpointConfigShadowProductionVariantsServerlessConfig)
	ResetCapacityReservationConfig()
	ResetContainerStartupHealthCheckTimeoutInSeconds()
	ResetCoreDumpConfig()
	ResetEnableSsmAccess()
	ResetInferenceAmiVersion()
	ResetInitialInstanceCount()
	ResetInitialVariantWeight()
	ResetInstancePools()
	ResetInstanceType()
	ResetManagedInstanceScaling()
	ResetModelDataDownloadTimeoutInSeconds()
	ResetModelName()
	ResetRoutingConfig()
	ResetServerlessConfig()
	ResetVariantInstanceProvisionTimeoutInSeconds()
	ResetVariantName()
	ResetVolumeSizeInGb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerEndpointConfigShadowProductionVariantsOutputReference
type jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) CapacityReservationConfig() SagemakerEndpointConfigShadowProductionVariantsCapacityReservationConfigOutputReference {
	var returns SagemakerEndpointConfigShadowProductionVariantsCapacityReservationConfigOutputReference
	_jsii_.Get(
		j,
		"capacityReservationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) CapacityReservationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityReservationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ContainerStartupHealthCheckTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ContainerStartupHealthCheckTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerStartupHealthCheckTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) CoreDumpConfig() SagemakerEndpointConfigShadowProductionVariantsCoreDumpConfigOutputReference {
	var returns SagemakerEndpointConfigShadowProductionVariantsCoreDumpConfigOutputReference
	_jsii_.Get(
		j,
		"coreDumpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) CoreDumpConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coreDumpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) EnableSsmAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) EnableSsmAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSsmAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InferenceAmiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InferenceAmiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferenceAmiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InitialInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InitialInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InitialVariantWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InitialVariantWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialVariantWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InstancePools() SagemakerEndpointConfigShadowProductionVariantsInstancePoolsList {
	var returns SagemakerEndpointConfigShadowProductionVariantsInstancePoolsList
	_jsii_.Get(
		j,
		"instancePools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InstancePoolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instancePoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ManagedInstanceScaling() SagemakerEndpointConfigShadowProductionVariantsManagedInstanceScalingOutputReference {
	var returns SagemakerEndpointConfigShadowProductionVariantsManagedInstanceScalingOutputReference
	_jsii_.Get(
		j,
		"managedInstanceScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ManagedInstanceScalingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedInstanceScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ModelDataDownloadTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ModelDataDownloadTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelDataDownloadTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) RoutingConfig() SagemakerEndpointConfigShadowProductionVariantsRoutingConfigOutputReference {
	var returns SagemakerEndpointConfigShadowProductionVariantsRoutingConfigOutputReference
	_jsii_.Get(
		j,
		"routingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) RoutingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ServerlessConfig() SagemakerEndpointConfigShadowProductionVariantsServerlessConfigOutputReference {
	var returns SagemakerEndpointConfigShadowProductionVariantsServerlessConfigOutputReference
	_jsii_.Get(
		j,
		"serverlessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ServerlessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) VariantInstanceProvisionTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"variantInstanceProvisionTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) VariantInstanceProvisionTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"variantInstanceProvisionTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) VariantName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) VariantNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variantNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) VolumeSizeInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) VolumeSizeInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInGbInput",
		&returns,
	)
	return returns
}


func NewSagemakerEndpointConfigShadowProductionVariantsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SagemakerEndpointConfigShadowProductionVariantsOutputReference {
	_init_.Initialize()

	if err := validateNewSagemakerEndpointConfigShadowProductionVariantsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigShadowProductionVariantsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSagemakerEndpointConfigShadowProductionVariantsOutputReference_Override(s SagemakerEndpointConfigShadowProductionVariantsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sagemakerEndpointConfig.SagemakerEndpointConfigShadowProductionVariantsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetContainerStartupHealthCheckTimeoutInSeconds(val *float64) {
	if err := j.validateSetContainerStartupHealthCheckTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerStartupHealthCheckTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetEnableSsmAccess(val interface{}) {
	if err := j.validateSetEnableSsmAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSsmAccess",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetInferenceAmiVersion(val *string) {
	if err := j.validateSetInferenceAmiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inferenceAmiVersion",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetInitialInstanceCount(val *float64) {
	if err := j.validateSetInitialInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialInstanceCount",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetInitialVariantWeight(val *float64) {
	if err := j.validateSetInitialVariantWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialVariantWeight",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetModelDataDownloadTimeoutInSeconds(val *float64) {
	if err := j.validateSetModelDataDownloadTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelDataDownloadTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetModelName(val *string) {
	if err := j.validateSetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelName",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetVariantInstanceProvisionTimeoutInSeconds(val *float64) {
	if err := j.validateSetVariantInstanceProvisionTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variantInstanceProvisionTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetVariantName(val *string) {
	if err := j.validateSetVariantNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variantName",
		val,
	)
}

func (j *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference)SetVolumeSizeInGb(val *float64) {
	if err := j.validateSetVolumeSizeInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSizeInGb",
		val,
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) PutCapacityReservationConfig(value *SagemakerEndpointConfigShadowProductionVariantsCapacityReservationConfig) {
	if err := s.validatePutCapacityReservationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCapacityReservationConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) PutCoreDumpConfig(value *SagemakerEndpointConfigShadowProductionVariantsCoreDumpConfig) {
	if err := s.validatePutCoreDumpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCoreDumpConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) PutInstancePools(value interface{}) {
	if err := s.validatePutInstancePoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInstancePools",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) PutManagedInstanceScaling(value *SagemakerEndpointConfigShadowProductionVariantsManagedInstanceScaling) {
	if err := s.validatePutManagedInstanceScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putManagedInstanceScaling",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) PutRoutingConfig(value *SagemakerEndpointConfigShadowProductionVariantsRoutingConfig) {
	if err := s.validatePutRoutingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRoutingConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) PutServerlessConfig(value *SagemakerEndpointConfigShadowProductionVariantsServerlessConfig) {
	if err := s.validatePutServerlessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putServerlessConfig",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetCapacityReservationConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacityReservationConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetContainerStartupHealthCheckTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetContainerStartupHealthCheckTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetCoreDumpConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetCoreDumpConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetEnableSsmAccess() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableSsmAccess",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetInferenceAmiVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetInferenceAmiVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetInitialInstanceCount() {
	_jsii_.InvokeVoid(
		s,
		"resetInitialInstanceCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetInitialVariantWeight() {
	_jsii_.InvokeVoid(
		s,
		"resetInitialVariantWeight",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetInstancePools() {
	_jsii_.InvokeVoid(
		s,
		"resetInstancePools",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetInstanceType() {
	_jsii_.InvokeVoid(
		s,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetManagedInstanceScaling() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedInstanceScaling",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetModelDataDownloadTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetModelDataDownloadTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetModelName() {
	_jsii_.InvokeVoid(
		s,
		"resetModelName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetRoutingConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetRoutingConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetServerlessConfig() {
	_jsii_.InvokeVoid(
		s,
		"resetServerlessConfig",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetVariantInstanceProvisionTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetVariantInstanceProvisionTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetVariantName() {
	_jsii_.InvokeVoid(
		s,
		"resetVariantName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ResetVolumeSizeInGb() {
	_jsii_.InvokeVoid(
		s,
		"resetVolumeSizeInGb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (s *jsiiProxy_SagemakerEndpointConfigShadowProductionVariantsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

