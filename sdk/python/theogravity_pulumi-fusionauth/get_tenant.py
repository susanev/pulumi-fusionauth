# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetTenantResult',
    'AwaitableGetTenantResult',
    'get_tenant',
    'get_tenant_output',
]

@pulumi.output_type
class GetTenantResult:
    """
    A collection of values returned by getTenant.
    """
    def __init__(__self__, id=None, name=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetTenantResult(GetTenantResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTenantResult(
            id=self.id,
            name=self.name)


def get_tenant(name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTenantResult:
    """
    ## # Tenant Resource

    A FusionAuth Tenant is a named object that represents a discrete namespace for Users, Applications and Groups. A user is unique by email address or username within a tenant.

    Tenants may be useful to support a multi-tenant application where you wish to use a single instance of FusionAuth but require the ability to have duplicate users across the tenants in your own application. In this scenario a user may exist multiple times with the same email address and different passwords across tenants.

    Tenants may also be useful in a test or staging environment to allow multiple users to call APIs and create and modify users without possibility of collision.

    [Tenants API](https://fusionauth.io/docs/v1/tech/apis/tenants)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    default = fusionauth.get_tenant(name="Default")
    ```


    :param str name: The name of the Tenant.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fusionauth:index/getTenant:getTenant', __args__, opts=opts, typ=GetTenantResult).value

    return AwaitableGetTenantResult(
        id=__ret__.id,
        name=__ret__.name)


@_utilities.lift_output_func(get_tenant)
def get_tenant_output(name: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTenantResult]:
    """
    ## # Tenant Resource

    A FusionAuth Tenant is a named object that represents a discrete namespace for Users, Applications and Groups. A user is unique by email address or username within a tenant.

    Tenants may be useful to support a multi-tenant application where you wish to use a single instance of FusionAuth but require the ability to have duplicate users across the tenants in your own application. In this scenario a user may exist multiple times with the same email address and different passwords across tenants.

    Tenants may also be useful in a test or staging environment to allow multiple users to call APIs and create and modify users without possibility of collision.

    [Tenants API](https://fusionauth.io/docs/v1/tech/apis/tenants)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    default = fusionauth.get_tenant(name="Default")
    ```


    :param str name: The name of the Tenant.
    """
    ...
