// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ec2verifiedaccessinstance/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference interface {
	cdktn.ComplexObject
	CloudwatchLogs() Ec2VerifiedAccessInstanceLoggingConfigurationsCloudwatchLogsOutputReference
	CloudwatchLogsInput() interface{}
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
	IncludeTrustContext() interface{}
	SetIncludeTrustContext(val interface{})
	IncludeTrustContextInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KinesisDataFirehose() Ec2VerifiedAccessInstanceLoggingConfigurationsKinesisDataFirehoseOutputReference
	KinesisDataFirehoseInput() interface{}
	LogVersion() *string
	SetLogVersion(val *string)
	LogVersionInput() *string
	S3() Ec2VerifiedAccessInstanceLoggingConfigurationsS3OutputReference
	S3Input() interface{}
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
	PutCloudwatchLogs(value *Ec2VerifiedAccessInstanceLoggingConfigurationsCloudwatchLogs)
	PutKinesisDataFirehose(value *Ec2VerifiedAccessInstanceLoggingConfigurationsKinesisDataFirehose)
	PutS3(value *Ec2VerifiedAccessInstanceLoggingConfigurationsS3)
	ResetCloudwatchLogs()
	ResetIncludeTrustContext()
	ResetKinesisDataFirehose()
	ResetLogVersion()
	ResetS3()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference
type jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) CloudwatchLogs() Ec2VerifiedAccessInstanceLoggingConfigurationsCloudwatchLogsOutputReference {
	var returns Ec2VerifiedAccessInstanceLoggingConfigurationsCloudwatchLogsOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) CloudwatchLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) IncludeTrustContext() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) IncludeTrustContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTrustContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) KinesisDataFirehose() Ec2VerifiedAccessInstanceLoggingConfigurationsKinesisDataFirehoseOutputReference {
	var returns Ec2VerifiedAccessInstanceLoggingConfigurationsKinesisDataFirehoseOutputReference
	_jsii_.Get(
		j,
		"kinesisDataFirehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) KinesisDataFirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisDataFirehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) LogVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) LogVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) S3() Ec2VerifiedAccessInstanceLoggingConfigurationsS3OutputReference {
	var returns Ec2VerifiedAccessInstanceLoggingConfigurationsS3OutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2VerifiedAccessInstanceLoggingConfigurationsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2VerifiedAccessInstanceLoggingConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2VerifiedAccessInstance.Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2VerifiedAccessInstanceLoggingConfigurationsOutputReference_Override(e Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ec2VerifiedAccessInstance.Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference)SetIncludeTrustContext(val interface{}) {
	if err := j.validateSetIncludeTrustContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTrustContext",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference)SetLogVersion(val *string) {
	if err := j.validateSetLogVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logVersion",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) PutCloudwatchLogs(value *Ec2VerifiedAccessInstanceLoggingConfigurationsCloudwatchLogs) {
	if err := e.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) PutKinesisDataFirehose(value *Ec2VerifiedAccessInstanceLoggingConfigurationsKinesisDataFirehose) {
	if err := e.validatePutKinesisDataFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putKinesisDataFirehose",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) PutS3(value *Ec2VerifiedAccessInstanceLoggingConfigurationsS3) {
	if err := e.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putS3",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		e,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) ResetIncludeTrustContext() {
	_jsii_.InvokeVoid(
		e,
		"resetIncludeTrustContext",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) ResetKinesisDataFirehose() {
	_jsii_.InvokeVoid(
		e,
		"resetKinesisDataFirehose",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) ResetLogVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetLogVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		e,
		"resetS3",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceLoggingConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

