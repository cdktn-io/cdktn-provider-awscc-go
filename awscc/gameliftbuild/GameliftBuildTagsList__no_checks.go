// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gameliftbuild

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GameliftBuildTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GameliftBuildTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GameliftBuildTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GameliftBuildTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GameliftBuildTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GameliftBuildTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GameliftBuildTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGameliftBuildTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

