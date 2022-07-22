# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetIdpResult',
    'AwaitableGetIdpResult',
    'get_idp',
    'get_idp_output',
]

@pulumi.output_type
class GetIdpResult:
    """
    A collection of values returned by getIdp.
    """
    def __init__(__self__, id=None, name=None, type=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

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

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


class AwaitableGetIdpResult(GetIdpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdpResult(
            id=self.id,
            name=self.name,
            type=self.type)


def get_idp(name: Optional[str] = None,
            type: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdpResult:
    """
    ## # Application Resource

    [Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    fusion_auth = fusionauth.get_idp(name="Apple",
        type="Apple")
    ```


    :param str name: The name of the identity provider. This is only used for display purposes. Will be the type for types: `Apple`, `Facebook`, `Google`, `HYPR`, `Twitter`
    :param str type: The type of the identity provider.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fusionauth:index/getIdp:getIdp', __args__, opts=opts, typ=GetIdpResult).value

    return AwaitableGetIdpResult(
        id=__ret__.id,
        name=__ret__.name,
        type=__ret__.type)


@_utilities.lift_output_func(get_idp)
def get_idp_output(name: Optional[pulumi.Input[str]] = None,
                   type: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIdpResult]:
    """
    ## # Application Resource

    [Identity Providers API](https://fusionauth.io/docs/v1/tech/apis/identity-providers/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    fusion_auth = fusionauth.get_idp(name="Apple",
        type="Apple")
    ```


    :param str name: The name of the identity provider. This is only used for display purposes. Will be the type for types: `Apple`, `Facebook`, `Google`, `HYPR`, `Twitter`
    :param str type: The type of the identity provider.
    """
    ...
