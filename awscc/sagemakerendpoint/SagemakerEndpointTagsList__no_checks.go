// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sagemakerendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SagemakerEndpointTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SagemakerEndpointTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SagemakerEndpointTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SagemakerEndpointTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SagemakerEndpointTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SagemakerEndpointTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SagemakerEndpointTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSagemakerEndpointTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

