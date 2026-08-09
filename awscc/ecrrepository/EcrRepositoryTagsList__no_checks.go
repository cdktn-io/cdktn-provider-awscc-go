// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ecrrepository

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcrRepositoryTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EcrRepositoryTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcrRepositoryTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcrRepositoryTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcrRepositoryTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcrRepositoryTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcrRepositoryTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcrRepositoryTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

