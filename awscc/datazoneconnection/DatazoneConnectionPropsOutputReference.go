// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazoneconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/datazoneconnection/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatazoneConnectionPropsOutputReference interface {
	cdktn.ComplexObject
	AmazonQProperties() DatazoneConnectionPropsAmazonQPropertiesOutputReference
	AmazonQPropertiesInput() interface{}
	AthenaProperties() DatazoneConnectionPropsAthenaPropertiesOutputReference
	AthenaPropertiesInput() interface{}
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
	GlueProperties() DatazoneConnectionPropsGluePropertiesOutputReference
	GluePropertiesInput() interface{}
	HyperPodProperties() DatazoneConnectionPropsHyperPodPropertiesOutputReference
	HyperPodPropertiesInput() interface{}
	IamProperties() DatazoneConnectionPropsIamPropertiesOutputReference
	IamPropertiesInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LakehouseProperties() DatazoneConnectionPropsLakehousePropertiesOutputReference
	LakehousePropertiesInput() interface{}
	MlflowProperties() DatazoneConnectionPropsMlflowPropertiesOutputReference
	MlflowPropertiesInput() interface{}
	RedshiftProperties() DatazoneConnectionPropsRedshiftPropertiesOutputReference
	RedshiftPropertiesInput() interface{}
	S3Properties() DatazoneConnectionPropsS3PropertiesOutputReference
	S3PropertiesInput() interface{}
	SparkEmrProperties() DatazoneConnectionPropsSparkEmrPropertiesOutputReference
	SparkEmrPropertiesInput() interface{}
	SparkGlueProperties() DatazoneConnectionPropsSparkGluePropertiesOutputReference
	SparkGluePropertiesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WorkflowsMwaaProperties() DatazoneConnectionPropsWorkflowsMwaaPropertiesOutputReference
	WorkflowsMwaaPropertiesInput() interface{}
	WorkflowsServerlessProperties() *string
	SetWorkflowsServerlessProperties(val *string)
	WorkflowsServerlessPropertiesInput() *string
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
	PutAmazonQProperties(value *DatazoneConnectionPropsAmazonQProperties)
	PutAthenaProperties(value *DatazoneConnectionPropsAthenaProperties)
	PutGlueProperties(value *DatazoneConnectionPropsGlueProperties)
	PutHyperPodProperties(value *DatazoneConnectionPropsHyperPodProperties)
	PutIamProperties(value *DatazoneConnectionPropsIamProperties)
	PutLakehouseProperties(value *DatazoneConnectionPropsLakehouseProperties)
	PutMlflowProperties(value *DatazoneConnectionPropsMlflowProperties)
	PutRedshiftProperties(value *DatazoneConnectionPropsRedshiftProperties)
	PutS3Properties(value *DatazoneConnectionPropsS3Properties)
	PutSparkEmrProperties(value *DatazoneConnectionPropsSparkEmrProperties)
	PutSparkGlueProperties(value *DatazoneConnectionPropsSparkGlueProperties)
	PutWorkflowsMwaaProperties(value *DatazoneConnectionPropsWorkflowsMwaaProperties)
	ResetAmazonQProperties()
	ResetAthenaProperties()
	ResetGlueProperties()
	ResetHyperPodProperties()
	ResetIamProperties()
	ResetLakehouseProperties()
	ResetMlflowProperties()
	ResetRedshiftProperties()
	ResetS3Properties()
	ResetSparkEmrProperties()
	ResetSparkGlueProperties()
	ResetWorkflowsMwaaProperties()
	ResetWorkflowsServerlessProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazoneConnectionPropsOutputReference
type jsiiProxy_DatazoneConnectionPropsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) AmazonQProperties() DatazoneConnectionPropsAmazonQPropertiesOutputReference {
	var returns DatazoneConnectionPropsAmazonQPropertiesOutputReference
	_jsii_.Get(
		j,
		"amazonQProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) AmazonQPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonQPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) AthenaProperties() DatazoneConnectionPropsAthenaPropertiesOutputReference {
	var returns DatazoneConnectionPropsAthenaPropertiesOutputReference
	_jsii_.Get(
		j,
		"athenaProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) AthenaPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"athenaPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) GlueProperties() DatazoneConnectionPropsGluePropertiesOutputReference {
	var returns DatazoneConnectionPropsGluePropertiesOutputReference
	_jsii_.Get(
		j,
		"glueProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) GluePropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gluePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) HyperPodProperties() DatazoneConnectionPropsHyperPodPropertiesOutputReference {
	var returns DatazoneConnectionPropsHyperPodPropertiesOutputReference
	_jsii_.Get(
		j,
		"hyperPodProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) HyperPodPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hyperPodPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) IamProperties() DatazoneConnectionPropsIamPropertiesOutputReference {
	var returns DatazoneConnectionPropsIamPropertiesOutputReference
	_jsii_.Get(
		j,
		"iamProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) IamPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) LakehouseProperties() DatazoneConnectionPropsLakehousePropertiesOutputReference {
	var returns DatazoneConnectionPropsLakehousePropertiesOutputReference
	_jsii_.Get(
		j,
		"lakehouseProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) LakehousePropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lakehousePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) MlflowProperties() DatazoneConnectionPropsMlflowPropertiesOutputReference {
	var returns DatazoneConnectionPropsMlflowPropertiesOutputReference
	_jsii_.Get(
		j,
		"mlflowProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) MlflowPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlflowPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) RedshiftProperties() DatazoneConnectionPropsRedshiftPropertiesOutputReference {
	var returns DatazoneConnectionPropsRedshiftPropertiesOutputReference
	_jsii_.Get(
		j,
		"redshiftProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) RedshiftPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) S3Properties() DatazoneConnectionPropsS3PropertiesOutputReference {
	var returns DatazoneConnectionPropsS3PropertiesOutputReference
	_jsii_.Get(
		j,
		"s3Properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) S3PropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3PropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) SparkEmrProperties() DatazoneConnectionPropsSparkEmrPropertiesOutputReference {
	var returns DatazoneConnectionPropsSparkEmrPropertiesOutputReference
	_jsii_.Get(
		j,
		"sparkEmrProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) SparkEmrPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparkEmrPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) SparkGlueProperties() DatazoneConnectionPropsSparkGluePropertiesOutputReference {
	var returns DatazoneConnectionPropsSparkGluePropertiesOutputReference
	_jsii_.Get(
		j,
		"sparkGlueProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) SparkGluePropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparkGluePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) WorkflowsMwaaProperties() DatazoneConnectionPropsWorkflowsMwaaPropertiesOutputReference {
	var returns DatazoneConnectionPropsWorkflowsMwaaPropertiesOutputReference
	_jsii_.Get(
		j,
		"workflowsMwaaProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) WorkflowsMwaaPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowsMwaaPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) WorkflowsServerlessProperties() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowsServerlessProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference) WorkflowsServerlessPropertiesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowsServerlessPropertiesInput",
		&returns,
	)
	return returns
}


func NewDatazoneConnectionPropsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DatazoneConnectionPropsOutputReference {
	_init_.Initialize()

	if err := validateNewDatazoneConnectionPropsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneConnectionPropsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatazoneConnectionPropsOutputReference_Override(d DatazoneConnectionPropsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.datazoneConnection.DatazoneConnectionPropsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatazoneConnectionPropsOutputReference)SetWorkflowsServerlessProperties(val *string) {
	if err := j.validateSetWorkflowsServerlessPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflowsServerlessProperties",
		val,
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutAmazonQProperties(value *DatazoneConnectionPropsAmazonQProperties) {
	if err := d.validatePutAmazonQPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAmazonQProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutAthenaProperties(value *DatazoneConnectionPropsAthenaProperties) {
	if err := d.validatePutAthenaPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAthenaProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutGlueProperties(value *DatazoneConnectionPropsGlueProperties) {
	if err := d.validatePutGluePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlueProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutHyperPodProperties(value *DatazoneConnectionPropsHyperPodProperties) {
	if err := d.validatePutHyperPodPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHyperPodProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutIamProperties(value *DatazoneConnectionPropsIamProperties) {
	if err := d.validatePutIamPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIamProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutLakehouseProperties(value *DatazoneConnectionPropsLakehouseProperties) {
	if err := d.validatePutLakehousePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLakehouseProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutMlflowProperties(value *DatazoneConnectionPropsMlflowProperties) {
	if err := d.validatePutMlflowPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMlflowProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutRedshiftProperties(value *DatazoneConnectionPropsRedshiftProperties) {
	if err := d.validatePutRedshiftPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedshiftProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutS3Properties(value *DatazoneConnectionPropsS3Properties) {
	if err := d.validatePutS3PropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3Properties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutSparkEmrProperties(value *DatazoneConnectionPropsSparkEmrProperties) {
	if err := d.validatePutSparkEmrPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkEmrProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutSparkGlueProperties(value *DatazoneConnectionPropsSparkGlueProperties) {
	if err := d.validatePutSparkGluePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkGlueProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) PutWorkflowsMwaaProperties(value *DatazoneConnectionPropsWorkflowsMwaaProperties) {
	if err := d.validatePutWorkflowsMwaaPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkflowsMwaaProperties",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetAmazonQProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetAmazonQProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetAthenaProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetAthenaProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetGlueProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetGlueProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetHyperPodProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetHyperPodProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetIamProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetIamProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetLakehouseProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetLakehouseProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetMlflowProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetMlflowProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetRedshiftProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetRedshiftProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetS3Properties() {
	_jsii_.InvokeVoid(
		d,
		"resetS3Properties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetSparkEmrProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkEmrProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetSparkGlueProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkGlueProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetWorkflowsMwaaProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkflowsMwaaProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ResetWorkflowsServerlessProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkflowsServerlessProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatazoneConnectionPropsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

