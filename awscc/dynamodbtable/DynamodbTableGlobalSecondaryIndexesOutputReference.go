// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dynamodbtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DynamodbTableGlobalSecondaryIndexesOutputReference interface {
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
	ContributorInsightsSpecification() DynamodbTableGlobalSecondaryIndexesContributorInsightsSpecificationOutputReference
	ContributorInsightsSpecificationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IndexName() *string
	SetIndexName(val *string)
	IndexNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeySchema() DynamodbTableGlobalSecondaryIndexesKeySchemaList
	KeySchemaInput() interface{}
	OnDemandThroughput() DynamodbTableGlobalSecondaryIndexesOnDemandThroughputOutputReference
	OnDemandThroughputInput() interface{}
	Projection() DynamodbTableGlobalSecondaryIndexesProjectionOutputReference
	ProjectionInput() interface{}
	ProvisionedThroughput() DynamodbTableGlobalSecondaryIndexesProvisionedThroughputOutputReference
	ProvisionedThroughputInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	WarmThroughput() DynamodbTableGlobalSecondaryIndexesWarmThroughputOutputReference
	WarmThroughputInput() interface{}
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
	PutContributorInsightsSpecification(value *DynamodbTableGlobalSecondaryIndexesContributorInsightsSpecification)
	PutKeySchema(value interface{})
	PutOnDemandThroughput(value *DynamodbTableGlobalSecondaryIndexesOnDemandThroughput)
	PutProjection(value *DynamodbTableGlobalSecondaryIndexesProjection)
	PutProvisionedThroughput(value *DynamodbTableGlobalSecondaryIndexesProvisionedThroughput)
	PutWarmThroughput(value *DynamodbTableGlobalSecondaryIndexesWarmThroughput)
	ResetContributorInsightsSpecification()
	ResetIndexName()
	ResetKeySchema()
	ResetOnDemandThroughput()
	ResetProjection()
	ResetProvisionedThroughput()
	ResetWarmThroughput()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamodbTableGlobalSecondaryIndexesOutputReference
type jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ContributorInsightsSpecification() DynamodbTableGlobalSecondaryIndexesContributorInsightsSpecificationOutputReference {
	var returns DynamodbTableGlobalSecondaryIndexesContributorInsightsSpecificationOutputReference
	_jsii_.Get(
		j,
		"contributorInsightsSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ContributorInsightsSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contributorInsightsSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) IndexName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) IndexNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) KeySchema() DynamodbTableGlobalSecondaryIndexesKeySchemaList {
	var returns DynamodbTableGlobalSecondaryIndexesKeySchemaList
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) KeySchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) OnDemandThroughput() DynamodbTableGlobalSecondaryIndexesOnDemandThroughputOutputReference {
	var returns DynamodbTableGlobalSecondaryIndexesOnDemandThroughputOutputReference
	_jsii_.Get(
		j,
		"onDemandThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) OnDemandThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) Projection() DynamodbTableGlobalSecondaryIndexesProjectionOutputReference {
	var returns DynamodbTableGlobalSecondaryIndexesProjectionOutputReference
	_jsii_.Get(
		j,
		"projection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ProjectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"projectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ProvisionedThroughput() DynamodbTableGlobalSecondaryIndexesProvisionedThroughputOutputReference {
	var returns DynamodbTableGlobalSecondaryIndexesProvisionedThroughputOutputReference
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ProvisionedThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) WarmThroughput() DynamodbTableGlobalSecondaryIndexesWarmThroughputOutputReference {
	var returns DynamodbTableGlobalSecondaryIndexesWarmThroughputOutputReference
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) WarmThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmThroughputInput",
		&returns,
	)
	return returns
}


func NewDynamodbTableGlobalSecondaryIndexesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DynamodbTableGlobalSecondaryIndexesOutputReference {
	_init_.Initialize()

	if err := validateNewDynamodbTableGlobalSecondaryIndexesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTableGlobalSecondaryIndexesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDynamodbTableGlobalSecondaryIndexesOutputReference_Override(d DynamodbTableGlobalSecondaryIndexesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTableGlobalSecondaryIndexesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference)SetIndexName(val *string) {
	if err := j.validateSetIndexNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexName",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) PutContributorInsightsSpecification(value *DynamodbTableGlobalSecondaryIndexesContributorInsightsSpecification) {
	if err := d.validatePutContributorInsightsSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContributorInsightsSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) PutKeySchema(value interface{}) {
	if err := d.validatePutKeySchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKeySchema",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) PutOnDemandThroughput(value *DynamodbTableGlobalSecondaryIndexesOnDemandThroughput) {
	if err := d.validatePutOnDemandThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOnDemandThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) PutProjection(value *DynamodbTableGlobalSecondaryIndexesProjection) {
	if err := d.validatePutProjectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProjection",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) PutProvisionedThroughput(value *DynamodbTableGlobalSecondaryIndexesProvisionedThroughput) {
	if err := d.validatePutProvisionedThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvisionedThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) PutWarmThroughput(value *DynamodbTableGlobalSecondaryIndexesWarmThroughput) {
	if err := d.validatePutWarmThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWarmThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ResetContributorInsightsSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetContributorInsightsSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ResetIndexName() {
	_jsii_.InvokeVoid(
		d,
		"resetIndexName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ResetKeySchema() {
	_jsii_.InvokeVoid(
		d,
		"resetKeySchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ResetOnDemandThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetOnDemandThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ResetProjection() {
	_jsii_.InvokeVoid(
		d,
		"resetProjection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ResetWarmThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetWarmThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

