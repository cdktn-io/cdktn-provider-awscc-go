// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/athenaworkgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference interface {
	cdktn.ComplexObject
	CloudwatchLoggingConfiguration() AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference
	CloudwatchLoggingConfigurationInput() interface{}
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
	ManagedLoggingConfiguration() AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationManagedLoggingConfigurationOutputReference
	ManagedLoggingConfigurationInput() interface{}
	S3LoggingConfiguration() AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationS3LoggingConfigurationOutputReference
	S3LoggingConfigurationInput() interface{}
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
	PutCloudwatchLoggingConfiguration(value *AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationCloudwatchLoggingConfiguration)
	PutManagedLoggingConfiguration(value *AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationManagedLoggingConfiguration)
	PutS3LoggingConfiguration(value *AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationS3LoggingConfiguration)
	ResetCloudwatchLoggingConfiguration()
	ResetManagedLoggingConfiguration()
	ResetS3LoggingConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference
type jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) CloudwatchLoggingConfiguration() AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationCloudwatchLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) CloudwatchLoggingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) ManagedLoggingConfiguration() AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationManagedLoggingConfigurationOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationManagedLoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"managedLoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) ManagedLoggingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedLoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) S3LoggingConfiguration() AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationS3LoggingConfigurationOutputReference {
	var returns AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationS3LoggingConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3LoggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) S3LoggingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3LoggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewAthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.athenaWorkGroup.AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference_Override(a AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.athenaWorkGroup.AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) PutCloudwatchLoggingConfiguration(value *AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationCloudwatchLoggingConfiguration) {
	if err := a.validatePutCloudwatchLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) PutManagedLoggingConfiguration(value *AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationManagedLoggingConfiguration) {
	if err := a.validatePutManagedLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManagedLoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) PutS3LoggingConfiguration(value *AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationS3LoggingConfiguration) {
	if err := a.validatePutS3LoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3LoggingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) ResetCloudwatchLoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) ResetManagedLoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedLoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) ResetS3LoggingConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetS3LoggingConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

