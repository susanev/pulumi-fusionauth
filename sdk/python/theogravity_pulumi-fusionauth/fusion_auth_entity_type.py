# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FusionAuthEntityTypeArgs', 'FusionAuthEntityType']

@pulumi.input_type
class FusionAuthEntityTypeArgs:
    def __init__(__self__, *,
                 data: Optional[pulumi.Input[str]] = None,
                 entity_type_id: Optional[pulumi.Input[str]] = None,
                 jwt_configuration: Optional[pulumi.Input['FusionAuthEntityTypeJwtConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FusionAuthEntityType resource.
        :param pulumi.Input[str] data: An object that can hold any information about the Entity Type that should be persisted. Must be a
               JSON string.
        :param pulumi.Input[str] entity_type_id: The ID to use for the new Entity Type. If not specified a secure random UUID will be
               generated.
        :param pulumi.Input['FusionAuthEntityTypeJwtConfigurationArgs'] jwt_configuration: A block to configure JSON Web Token (JWT) options.
        :param pulumi.Input[str] name: A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        """
        if data is not None:
            pulumi.set(__self__, "data", data)
        if entity_type_id is not None:
            pulumi.set(__self__, "entity_type_id", entity_type_id)
        if jwt_configuration is not None:
            pulumi.set(__self__, "jwt_configuration", jwt_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        An object that can hold any information about the Entity Type that should be persisted. Must be a
        JSON string.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="entityTypeId")
    def entity_type_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID to use for the new Entity Type. If not specified a secure random UUID will be
        generated.
        """
        return pulumi.get(self, "entity_type_id")

    @entity_type_id.setter
    def entity_type_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_type_id", value)

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> Optional[pulumi.Input['FusionAuthEntityTypeJwtConfigurationArgs']]:
        """
        A block to configure JSON Web Token (JWT) options.
        """
        return pulumi.get(self, "jwt_configuration")

    @jwt_configuration.setter
    def jwt_configuration(self, value: Optional[pulumi.Input['FusionAuthEntityTypeJwtConfigurationArgs']]):
        pulumi.set(self, "jwt_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _FusionAuthEntityTypeState:
    def __init__(__self__, *,
                 data: Optional[pulumi.Input[str]] = None,
                 entity_type_id: Optional[pulumi.Input[str]] = None,
                 jwt_configuration: Optional[pulumi.Input['FusionAuthEntityTypeJwtConfigurationArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FusionAuthEntityType resources.
        :param pulumi.Input[str] data: An object that can hold any information about the Entity Type that should be persisted. Must be a
               JSON string.
        :param pulumi.Input[str] entity_type_id: The ID to use for the new Entity Type. If not specified a secure random UUID will be
               generated.
        :param pulumi.Input['FusionAuthEntityTypeJwtConfigurationArgs'] jwt_configuration: A block to configure JSON Web Token (JWT) options.
        :param pulumi.Input[str] name: A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        """
        if data is not None:
            pulumi.set(__self__, "data", data)
        if entity_type_id is not None:
            pulumi.set(__self__, "entity_type_id", entity_type_id)
        if jwt_configuration is not None:
            pulumi.set(__self__, "jwt_configuration", jwt_configuration)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        An object that can hold any information about the Entity Type that should be persisted. Must be a
        JSON string.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="entityTypeId")
    def entity_type_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID to use for the new Entity Type. If not specified a secure random UUID will be
        generated.
        """
        return pulumi.get(self, "entity_type_id")

    @entity_type_id.setter
    def entity_type_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_type_id", value)

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> Optional[pulumi.Input['FusionAuthEntityTypeJwtConfigurationArgs']]:
        """
        A block to configure JSON Web Token (JWT) options.
        """
        return pulumi.get(self, "jwt_configuration")

    @jwt_configuration.setter
    def jwt_configuration(self, value: Optional[pulumi.Input['FusionAuthEntityTypeJwtConfigurationArgs']]):
        pulumi.set(self, "jwt_configuration", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class FusionAuthEntityType(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 entity_type_id: Optional[pulumi.Input[str]] = None,
                 jwt_configuration: Optional[pulumi.Input[pulumi.InputType['FusionAuthEntityTypeJwtConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Entity Type Resource

        Entity Types categorize Entities. For example, an Entity Type could be `Device`, `API` or `Company`.

        [Entity Type API](https://fusionauth.io/docs/v1/tech/apis/entity-management/entity-types/#create-an-entity-type)

        ## Example Usage

        ```python
        import pulumi
        import json
        import theogravity_pulumi-fusionauth as fusionauth

        company = fusionauth.FusionAuthEntityType("company",
            data=json.dumps({
                "createdBy": "jared@fusionauth.io",
            }),
            jwt_configuration=fusionauth.FusionAuthEntityTypeJwtConfigurationArgs(
                access_token_key_id="a7516c7c-6234-4021-b0b4-8870c807aeb2",
                enabled=True,
                time_to_live_in_seconds=3600,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: An object that can hold any information about the Entity Type that should be persisted. Must be a
               JSON string.
        :param pulumi.Input[str] entity_type_id: The ID to use for the new Entity Type. If not specified a secure random UUID will be
               generated.
        :param pulumi.Input[pulumi.InputType['FusionAuthEntityTypeJwtConfigurationArgs']] jwt_configuration: A block to configure JSON Web Token (JWT) options.
        :param pulumi.Input[str] name: A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FusionAuthEntityTypeArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Entity Type Resource

        Entity Types categorize Entities. For example, an Entity Type could be `Device`, `API` or `Company`.

        [Entity Type API](https://fusionauth.io/docs/v1/tech/apis/entity-management/entity-types/#create-an-entity-type)

        ## Example Usage

        ```python
        import pulumi
        import json
        import theogravity_pulumi-fusionauth as fusionauth

        company = fusionauth.FusionAuthEntityType("company",
            data=json.dumps({
                "createdBy": "jared@fusionauth.io",
            }),
            jwt_configuration=fusionauth.FusionAuthEntityTypeJwtConfigurationArgs(
                access_token_key_id="a7516c7c-6234-4021-b0b4-8870c807aeb2",
                enabled=True,
                time_to_live_in_seconds=3600,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param FusionAuthEntityTypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FusionAuthEntityTypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 entity_type_id: Optional[pulumi.Input[str]] = None,
                 jwt_configuration: Optional[pulumi.Input[pulumi.InputType['FusionAuthEntityTypeJwtConfigurationArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FusionAuthEntityTypeArgs.__new__(FusionAuthEntityTypeArgs)

            __props__.__dict__["data"] = data
            __props__.__dict__["entity_type_id"] = entity_type_id
            __props__.__dict__["jwt_configuration"] = jwt_configuration
            __props__.__dict__["name"] = name
        super(FusionAuthEntityType, __self__).__init__(
            'fusionauth:index/fusionAuthEntityType:FusionAuthEntityType',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data: Optional[pulumi.Input[str]] = None,
            entity_type_id: Optional[pulumi.Input[str]] = None,
            jwt_configuration: Optional[pulumi.Input[pulumi.InputType['FusionAuthEntityTypeJwtConfigurationArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'FusionAuthEntityType':
        """
        Get an existing FusionAuthEntityType resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data: An object that can hold any information about the Entity Type that should be persisted. Must be a
               JSON string.
        :param pulumi.Input[str] entity_type_id: The ID to use for the new Entity Type. If not specified a secure random UUID will be
               generated.
        :param pulumi.Input[pulumi.InputType['FusionAuthEntityTypeJwtConfigurationArgs']] jwt_configuration: A block to configure JSON Web Token (JWT) options.
        :param pulumi.Input[str] name: A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FusionAuthEntityTypeState.__new__(_FusionAuthEntityTypeState)

        __props__.__dict__["data"] = data
        __props__.__dict__["entity_type_id"] = entity_type_id
        __props__.__dict__["jwt_configuration"] = jwt_configuration
        __props__.__dict__["name"] = name
        return FusionAuthEntityType(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Optional[str]]:
        """
        An object that can hold any information about the Entity Type that should be persisted. Must be a
        JSON string.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="entityTypeId")
    def entity_type_id(self) -> pulumi.Output[str]:
        """
        The ID to use for the new Entity Type. If not specified a secure random UUID will be
        generated.
        """
        return pulumi.get(self, "entity_type_id")

    @property
    @pulumi.getter(name="jwtConfiguration")
    def jwt_configuration(self) -> pulumi.Output['outputs.FusionAuthEntityTypeJwtConfiguration']:
        """
        A block to configure JSON Web Token (JWT) options.
        """
        return pulumi.get(self, "jwt_configuration")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        """
        return pulumi.get(self, "name")

