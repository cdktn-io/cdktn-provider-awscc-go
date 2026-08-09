// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccdatazoneconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccdatazoneconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccDatazoneConnectionPropsOutputReference interface {
	cdktn.ComplexObject
	AmazonQProperties() DataAwsccDatazoneConnectionPropsAmazonQPropertiesOutputReference
	AthenaProperties() DataAwsccDatazoneConnectionPropsAthenaPropertiesOutputReference
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
	GlueProperties() DataAwsccDatazoneConnectionPropsGluePropertiesOutputReference
	HyperPodProperties() DataAwsccDatazoneConnectionPropsHyperPodPropertiesOutputReference
	IamProperties() DataAwsccDatazoneConnectionPropsIamPropertiesOutputReference
	InternalValue() *DataAwsccDatazoneConnectionProps
	SetInternalValue(val *DataAwsccDatazoneConnectionProps)
	LakehouseProperties() DataAwsccDatazoneConnectionPropsLakehousePropertiesOutputReference
	MlflowProperties() DataAwsccDatazoneConnectionPropsMlflowPropertiesOutputReference
	RedshiftProperties() DataAwsccDatazoneConnectionPropsRedshiftPropertiesOutputReference
	S3Properties() DataAwsccDatazoneConnectionPropsS3PropertiesOutputReference
	SparkEmrProperties() DataAwsccDatazoneConnectionPropsSparkEmrPropertiesOutputReference
	SparkGlueProperties() DataAwsccDatazoneConnectionPropsSparkGluePropertiesOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkflowsMwaaProperties() DataAwsccDatazoneConnectionPropsWorkflowsMwaaPropertiesOutputReference
	WorkflowsServerlessProperties() *string
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

// The jsii proxy struct for DataAwsccDatazoneConnectionPropsOutputReference
type jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) AmazonQProperties() DataAwsccDatazoneConnectionPropsAmazonQPropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsAmazonQPropertiesOutputReference
	_jsii_.Get(
		j,
		"amazonQProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) AthenaProperties() DataAwsccDatazoneConnectionPropsAthenaPropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsAthenaPropertiesOutputReference
	_jsii_.Get(
		j,
		"athenaProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GlueProperties() DataAwsccDatazoneConnectionPropsGluePropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsGluePropertiesOutputReference
	_jsii_.Get(
		j,
		"glueProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) HyperPodProperties() DataAwsccDatazoneConnectionPropsHyperPodPropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsHyperPodPropertiesOutputReference
	_jsii_.Get(
		j,
		"hyperPodProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) IamProperties() DataAwsccDatazoneConnectionPropsIamPropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsIamPropertiesOutputReference
	_jsii_.Get(
		j,
		"iamProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) InternalValue() *DataAwsccDatazoneConnectionProps {
	var returns *DataAwsccDatazoneConnectionProps
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) LakehouseProperties() DataAwsccDatazoneConnectionPropsLakehousePropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsLakehousePropertiesOutputReference
	_jsii_.Get(
		j,
		"lakehouseProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) MlflowProperties() DataAwsccDatazoneConnectionPropsMlflowPropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsMlflowPropertiesOutputReference
	_jsii_.Get(
		j,
		"mlflowProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) RedshiftProperties() DataAwsccDatazoneConnectionPropsRedshiftPropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsRedshiftPropertiesOutputReference
	_jsii_.Get(
		j,
		"redshiftProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) S3Properties() DataAwsccDatazoneConnectionPropsS3PropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsS3PropertiesOutputReference
	_jsii_.Get(
		j,
		"s3Properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) SparkEmrProperties() DataAwsccDatazoneConnectionPropsSparkEmrPropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsSparkEmrPropertiesOutputReference
	_jsii_.Get(
		j,
		"sparkEmrProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) SparkGlueProperties() DataAwsccDatazoneConnectionPropsSparkGluePropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsSparkGluePropertiesOutputReference
	_jsii_.Get(
		j,
		"sparkGlueProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) WorkflowsMwaaProperties() DataAwsccDatazoneConnectionPropsWorkflowsMwaaPropertiesOutputReference {
	var returns DataAwsccDatazoneConnectionPropsWorkflowsMwaaPropertiesOutputReference
	_jsii_.Get(
		j,
		"workflowsMwaaProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) WorkflowsServerlessProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowsServerlessProperties",
		&returns,
	)
	return returns
}


func NewDataAwsccDatazoneConnectionPropsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccDatazoneConnectionPropsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccDatazoneConnectionPropsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDatazoneConnection.DataAwsccDatazoneConnectionPropsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccDatazoneConnectionPropsOutputReference_Override(d DataAwsccDatazoneConnectionPropsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccDatazoneConnection.DataAwsccDatazoneConnectionPropsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference)SetInternalValue(val *DataAwsccDatazoneConnectionProps) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccDatazoneConnectionPropsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

