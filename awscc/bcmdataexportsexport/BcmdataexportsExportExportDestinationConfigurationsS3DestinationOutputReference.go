// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmdataexportsexport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/bcmdataexportsexport/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3BucketOwner() *string
	SetS3BucketOwner(val *string)
	S3BucketOwnerInput() *string
	S3OutputConfigurations() BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsOutputReference
	S3OutputConfigurationsInput() interface{}
	S3Prefix() *string
	SetS3Prefix(val *string)
	S3PrefixInput() *string
	S3Region() *string
	SetS3Region(val *string)
	S3RegionInput() *string
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
	PutS3OutputConfigurations(value *BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurations)
	ResetS3BucketOwner()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference
type jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3BucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3BucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3OutputConfigurations() BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsOutputReference {
	var returns BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurationsOutputReference
	_jsii_.Get(
		j,
		"s3OutputConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3OutputConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3OutputConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) S3RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3RegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference {
	_init_.Initialize()

	if err := validateNewBcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.bcmdataexportsExport.BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference_Override(b BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.bcmdataexportsExport.BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference)SetS3BucketOwner(val *string) {
	if err := j.validateSetS3BucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketOwner",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference)SetS3Prefix(val *string) {
	if err := j.validateSetS3PrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Prefix",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference)SetS3Region(val *string) {
	if err := j.validateSetS3RegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Region",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) PutS3OutputConfigurations(value *BcmdataexportsExportExportDestinationConfigurationsS3DestinationS3OutputConfigurations) {
	if err := b.validatePutS3OutputConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putS3OutputConfigurations",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) ResetS3BucketOwner() {
	_jsii_.InvokeVoid(
		b,
		"resetS3BucketOwner",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (b *jsiiProxy_BcmdataexportsExportExportDestinationConfigurationsS3DestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

