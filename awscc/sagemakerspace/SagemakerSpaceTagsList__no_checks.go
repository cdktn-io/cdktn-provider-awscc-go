// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sagemakerspace

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SagemakerSpaceTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SagemakerSpaceTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SagemakerSpaceTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SagemakerSpaceTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SagemakerSpaceTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SagemakerSpaceTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SagemakerSpaceTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSagemakerSpaceTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

