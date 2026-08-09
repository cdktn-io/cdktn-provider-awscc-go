// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3accessgrant

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3AccessGrantTagsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3AccessGrantTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3AccessGrantTagsList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantTagsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantTagsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3AccessGrantTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3AccessGrantTagsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

