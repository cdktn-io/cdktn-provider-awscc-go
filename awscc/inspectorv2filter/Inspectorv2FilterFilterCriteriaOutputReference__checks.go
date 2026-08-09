// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package inspectorv2filter

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutAwsAccountIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaAwsAccountId:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaAwsAccountId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaAwsAccountId:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaAwsAccountId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaAwsAccountId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutCodeVulnerabilityDetectorNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorName:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorName:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutCodeVulnerabilityDetectorTagsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorTags:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorTags)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorTags:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorTags)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityDetectorTags; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutCodeVulnerabilityFilePathParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityFilePath:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityFilePath)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaCodeVulnerabilityFilePath:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityFilePath)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaCodeVulnerabilityFilePath; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutComponentIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaComponentId:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaComponentId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaComponentId:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaComponentId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaComponentId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutComponentTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaComponentType:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaComponentType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaComponentType:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaComponentType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaComponentType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEc2InstanceImageIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEc2InstanceImageId:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEc2InstanceImageId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEc2InstanceImageId:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEc2InstanceImageId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEc2InstanceImageId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEc2InstanceSubnetIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEc2InstanceSubnetId:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEc2InstanceSubnetId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEc2InstanceSubnetId:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEc2InstanceSubnetId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEc2InstanceSubnetId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEc2InstanceVpcIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEc2InstanceVpcId:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEc2InstanceVpcId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEc2InstanceVpcId:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEc2InstanceVpcId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEc2InstanceVpcId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEcrImageArchitectureParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEcrImageArchitecture:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEcrImageArchitecture)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEcrImageArchitecture:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEcrImageArchitecture)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEcrImageArchitecture; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEcrImageHashParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEcrImageHash:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEcrImageHash)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEcrImageHash:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEcrImageHash)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEcrImageHash; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEcrImagePushedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEcrImagePushedAt:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEcrImagePushedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEcrImagePushedAt:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEcrImagePushedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEcrImagePushedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEcrImageRegistryParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEcrImageRegistry:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEcrImageRegistry)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEcrImageRegistry:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEcrImageRegistry)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEcrImageRegistry; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEcrImageRepositoryNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEcrImageRepositoryName:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEcrImageRepositoryName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEcrImageRepositoryName:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEcrImageRepositoryName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEcrImageRepositoryName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEcrImageTagsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEcrImageTags:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEcrImageTags)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEcrImageTags:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEcrImageTags)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEcrImageTags; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutEpssScoreParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaEpssScore:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaEpssScore)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaEpssScore:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaEpssScore)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaEpssScore; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutExploitAvailableParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaExploitAvailable:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaExploitAvailable)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaExploitAvailable:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaExploitAvailable)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaExploitAvailable; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutFindingArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaFindingArn:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaFindingArn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaFindingArn:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaFindingArn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaFindingArn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutFindingStatusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaFindingStatus:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaFindingStatus)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaFindingStatus:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaFindingStatus)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaFindingStatus; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutFindingTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaFindingType:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaFindingType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaFindingType:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaFindingType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaFindingType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutFirstObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaFirstObservedAt:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaFirstObservedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaFirstObservedAt:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaFirstObservedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaFirstObservedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutFixAvailableParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaFixAvailable:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaFixAvailable)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaFixAvailable:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaFixAvailable)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaFixAvailable; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutInspectorScoreParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaInspectorScore:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaInspectorScore)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaInspectorScore:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaInspectorScore)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaInspectorScore; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionExecutionRoleArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionExecutionRoleArn:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaLambdaFunctionExecutionRoleArn)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaLambdaFunctionExecutionRoleArn:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaLambdaFunctionExecutionRoleArn)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionExecutionRoleArn; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionLastModifiedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionLastModifiedAt:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaLambdaFunctionLastModifiedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaLambdaFunctionLastModifiedAt:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaLambdaFunctionLastModifiedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionLastModifiedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionLayersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionLayers:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaLambdaFunctionLayers)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaLambdaFunctionLayers:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaLambdaFunctionLayers)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionLayers; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionName:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaLambdaFunctionName)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaLambdaFunctionName:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaLambdaFunctionName)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionName; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutLambdaFunctionRuntimeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionRuntime:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaLambdaFunctionRuntime)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaLambdaFunctionRuntime:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaLambdaFunctionRuntime)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaLambdaFunctionRuntime; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutLastObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaLastObservedAt:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaLastObservedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaLastObservedAt:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaLastObservedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaLastObservedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutNetworkProtocolParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaNetworkProtocol:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaNetworkProtocol)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaNetworkProtocol:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaNetworkProtocol)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaNetworkProtocol; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutPortRangeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaPortRange:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaPortRange)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaPortRange:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaPortRange)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaPortRange; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutRelatedVulnerabilitiesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaRelatedVulnerabilities:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaRelatedVulnerabilities)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaRelatedVulnerabilities:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaRelatedVulnerabilities)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaRelatedVulnerabilities; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutResourceIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaResourceId:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaResourceId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaResourceId:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaResourceId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaResourceId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutResourceTagsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaResourceTags:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaResourceTags)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaResourceTags:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaResourceTags)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaResourceTags; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutResourceTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaResourceType:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaResourceType)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaResourceType:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaResourceType)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaResourceType; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutSeverityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaSeverity:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaSeverity)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaSeverity:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaSeverity)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaSeverity; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutTitleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaTitle:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaTitle)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaTitle:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaTitle)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaTitle; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutUpdatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaUpdatedAt:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaUpdatedAt)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaUpdatedAt:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaUpdatedAt)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaUpdatedAt; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutVendorSeverityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaVendorSeverity:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaVendorSeverity)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaVendorSeverity:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaVendorSeverity)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaVendorSeverity; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutVulnerabilityIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaVulnerabilityId:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaVulnerabilityId)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaVulnerabilityId:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaVulnerabilityId)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaVulnerabilityId; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutVulnerabilitySourceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaVulnerabilitySource:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaVulnerabilitySource)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaVulnerabilitySource:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaVulnerabilitySource)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaVulnerabilitySource; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validatePutVulnerablePackagesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Inspectorv2FilterFilterCriteriaVulnerablePackages:
		value := value.(*[]*Inspectorv2FilterFilterCriteriaVulnerablePackages)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Inspectorv2FilterFilterCriteriaVulnerablePackages:
		value_ := value.([]*Inspectorv2FilterFilterCriteriaVulnerablePackages)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Inspectorv2FilterFilterCriteriaVulnerablePackages; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (i *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *Inspectorv2FilterFilterCriteria:
		val := val.(*Inspectorv2FilterFilterCriteria)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Inspectorv2FilterFilterCriteria:
		val_ := val.(Inspectorv2FilterFilterCriteria)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *Inspectorv2FilterFilterCriteria; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Inspectorv2FilterFilterCriteriaOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewInspectorv2FilterFilterCriteriaOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

