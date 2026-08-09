// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package elasticacheuser

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElasticacheUserTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_ElasticacheUserTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElasticacheUserTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElasticacheUserTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ElasticacheUserTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElasticacheUserTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElasticacheUserTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElasticacheUserTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

