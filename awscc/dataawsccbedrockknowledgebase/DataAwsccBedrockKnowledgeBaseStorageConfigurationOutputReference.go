// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccbedrockknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccbedrockknowledgebase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference interface {
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
	InternalValue() *DataAwsccBedrockKnowledgeBaseStorageConfiguration
	SetInternalValue(val *DataAwsccBedrockKnowledgeBaseStorageConfiguration)
	MongoDbAtlasConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference
	NeptuneAnalyticsConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationOutputReference
	OpensearchManagedClusterConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationOutputReference
	OpensearchServerlessConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationOutputReference
	PineconeConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationPineconeConfigurationOutputReference
	RdsConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationRdsConfigurationOutputReference
	S3VectorsConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Type() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference
type jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) InternalValue() *DataAwsccBedrockKnowledgeBaseStorageConfiguration {
	var returns *DataAwsccBedrockKnowledgeBaseStorageConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) MongoDbAtlasConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference {
	var returns DataAwsccBedrockKnowledgeBaseStorageConfigurationMongoDbAtlasConfigurationOutputReference
	_jsii_.Get(
		j,
		"mongoDbAtlasConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) NeptuneAnalyticsConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationOutputReference {
	var returns DataAwsccBedrockKnowledgeBaseStorageConfigurationNeptuneAnalyticsConfigurationOutputReference
	_jsii_.Get(
		j,
		"neptuneAnalyticsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) OpensearchManagedClusterConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationOutputReference {
	var returns DataAwsccBedrockKnowledgeBaseStorageConfigurationOpensearchManagedClusterConfigurationOutputReference
	_jsii_.Get(
		j,
		"opensearchManagedClusterConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) OpensearchServerlessConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationOutputReference {
	var returns DataAwsccBedrockKnowledgeBaseStorageConfigurationOpensearchServerlessConfigurationOutputReference
	_jsii_.Get(
		j,
		"opensearchServerlessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) PineconeConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationPineconeConfigurationOutputReference {
	var returns DataAwsccBedrockKnowledgeBaseStorageConfigurationPineconeConfigurationOutputReference
	_jsii_.Get(
		j,
		"pineconeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) RdsConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationRdsConfigurationOutputReference {
	var returns DataAwsccBedrockKnowledgeBaseStorageConfigurationRdsConfigurationOutputReference
	_jsii_.Get(
		j,
		"rdsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) S3VectorsConfiguration() DataAwsccBedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference {
	var returns DataAwsccBedrockKnowledgeBaseStorageConfigurationS3VectorsConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3VectorsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


func NewDataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockKnowledgeBase.DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference_Override(d DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccBedrockKnowledgeBase.DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference)SetInternalValue(val *DataAwsccBedrockKnowledgeBaseStorageConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccBedrockKnowledgeBaseStorageConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

