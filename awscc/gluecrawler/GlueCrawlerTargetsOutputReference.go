// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluecrawler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gluecrawler/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueCrawlerTargetsOutputReference interface {
	cdktn.ComplexObject
	CatalogTargets() GlueCrawlerTargetsCatalogTargetsList
	CatalogTargetsInput() interface{}
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
	DeltaTargets() GlueCrawlerTargetsDeltaTargetsList
	DeltaTargetsInput() interface{}
	DynamoDbTargets() GlueCrawlerTargetsDynamoDbTargetsList
	DynamoDbTargetsInput() interface{}
	// Experimental.
	Fqn() *string
	HudiTargets() GlueCrawlerTargetsHudiTargetsList
	HudiTargetsInput() interface{}
	IcebergTargets() GlueCrawlerTargetsIcebergTargetsList
	IcebergTargetsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	JdbcTargets() GlueCrawlerTargetsJdbcTargetsList
	JdbcTargetsInput() interface{}
	MongoDbTargets() GlueCrawlerTargetsMongoDbTargetsList
	MongoDbTargetsInput() interface{}
	S3Targets() GlueCrawlerTargetsS3TargetsList
	S3TargetsInput() interface{}
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
	PutCatalogTargets(value interface{})
	PutDeltaTargets(value interface{})
	PutDynamoDbTargets(value interface{})
	PutHudiTargets(value interface{})
	PutIcebergTargets(value interface{})
	PutJdbcTargets(value interface{})
	PutMongoDbTargets(value interface{})
	PutS3Targets(value interface{})
	ResetCatalogTargets()
	ResetDeltaTargets()
	ResetDynamoDbTargets()
	ResetHudiTargets()
	ResetIcebergTargets()
	ResetJdbcTargets()
	ResetMongoDbTargets()
	ResetS3Targets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GlueCrawlerTargetsOutputReference
type jsiiProxy_GlueCrawlerTargetsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) CatalogTargets() GlueCrawlerTargetsCatalogTargetsList {
	var returns GlueCrawlerTargetsCatalogTargetsList
	_jsii_.Get(
		j,
		"catalogTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) CatalogTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) DeltaTargets() GlueCrawlerTargetsDeltaTargetsList {
	var returns GlueCrawlerTargetsDeltaTargetsList
	_jsii_.Get(
		j,
		"deltaTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) DeltaTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deltaTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) DynamoDbTargets() GlueCrawlerTargetsDynamoDbTargetsList {
	var returns GlueCrawlerTargetsDynamoDbTargetsList
	_jsii_.Get(
		j,
		"dynamoDbTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) DynamoDbTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) HudiTargets() GlueCrawlerTargetsHudiTargetsList {
	var returns GlueCrawlerTargetsHudiTargetsList
	_jsii_.Get(
		j,
		"hudiTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) HudiTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hudiTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) IcebergTargets() GlueCrawlerTargetsIcebergTargetsList {
	var returns GlueCrawlerTargetsIcebergTargetsList
	_jsii_.Get(
		j,
		"icebergTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) IcebergTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"icebergTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) JdbcTargets() GlueCrawlerTargetsJdbcTargetsList {
	var returns GlueCrawlerTargetsJdbcTargetsList
	_jsii_.Get(
		j,
		"jdbcTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) JdbcTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jdbcTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) MongoDbTargets() GlueCrawlerTargetsMongoDbTargetsList {
	var returns GlueCrawlerTargetsMongoDbTargetsList
	_jsii_.Get(
		j,
		"mongoDbTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) MongoDbTargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongoDbTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) S3Targets() GlueCrawlerTargetsS3TargetsList {
	var returns GlueCrawlerTargetsS3TargetsList
	_jsii_.Get(
		j,
		"s3Targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) S3TargetsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3TargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGlueCrawlerTargetsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GlueCrawlerTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewGlueCrawlerTargetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCrawlerTargetsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueCrawler.GlueCrawlerTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGlueCrawlerTargetsOutputReference_Override(g GlueCrawlerTargetsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueCrawler.GlueCrawlerTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GlueCrawlerTargetsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) PutCatalogTargets(value interface{}) {
	if err := g.validatePutCatalogTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCatalogTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) PutDeltaTargets(value interface{}) {
	if err := g.validatePutDeltaTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeltaTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) PutDynamoDbTargets(value interface{}) {
	if err := g.validatePutDynamoDbTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDynamoDbTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) PutHudiTargets(value interface{}) {
	if err := g.validatePutHudiTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHudiTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) PutIcebergTargets(value interface{}) {
	if err := g.validatePutIcebergTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIcebergTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) PutJdbcTargets(value interface{}) {
	if err := g.validatePutJdbcTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJdbcTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) PutMongoDbTargets(value interface{}) {
	if err := g.validatePutMongoDbTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMongoDbTargets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) PutS3Targets(value interface{}) {
	if err := g.validatePutS3TargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putS3Targets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ResetCatalogTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ResetDeltaTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetDeltaTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ResetDynamoDbTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetDynamoDbTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ResetHudiTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetHudiTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ResetIcebergTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetIcebergTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ResetJdbcTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetJdbcTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ResetMongoDbTargets() {
	_jsii_.InvokeVoid(
		g,
		"resetMongoDbTargets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ResetS3Targets() {
	_jsii_.InvokeVoid(
		g,
		"resetS3Targets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCrawlerTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

