// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sagemakerhub

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SagemakerHubTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SagemakerHubTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SagemakerHubTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SagemakerHubTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SagemakerHubTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SagemakerHubTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SagemakerHubTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSagemakerHubTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

