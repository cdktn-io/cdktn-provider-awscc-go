// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/appfabricingestiondestination/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference interface {
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
	FirehoseStream() AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationFirehoseStreamOutputReference
	FirehoseStreamInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3Bucket() AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketOutputReference
	S3BucketInput() interface{}
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
	PutFirehoseStream(value *AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationFirehoseStream)
	PutS3Bucket(value *AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3Bucket)
	ResetFirehoseStream()
	ResetS3Bucket()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference
type jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) FirehoseStream() AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationFirehoseStreamOutputReference {
	var returns AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationFirehoseStreamOutputReference
	_jsii_.Get(
		j,
		"firehoseStream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) FirehoseStreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseStreamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) S3Bucket() AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketOutputReference {
	var returns AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3BucketOutputReference
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) S3BucketInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewAppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.appfabricIngestionDestination.AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference_Override(a AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.appfabricIngestionDestination.AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) PutFirehoseStream(value *AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationFirehoseStream) {
	if err := a.validatePutFirehoseStreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirehoseStream",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) PutS3Bucket(value *AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationS3Bucket) {
	if err := a.validatePutS3BucketParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Bucket",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) ResetFirehoseStream() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehoseStream",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppfabricIngestionDestinationDestinationConfigurationAuditLogDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

