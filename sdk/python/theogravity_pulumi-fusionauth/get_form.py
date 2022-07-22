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

__all__ = [
    'GetFormResult',
    'AwaitableGetFormResult',
    'get_form',
    'get_form_output',
]

@pulumi.output_type
class GetFormResult:
    """
    A collection of values returned by getForm.
    """
    def __init__(__self__, data=None, form_id=None, id=None, name=None, steps=None, type=None):
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)
        if form_id and not isinstance(form_id, str):
            raise TypeError("Expected argument 'form_id' to be a str")
        pulumi.set(__self__, "form_id", form_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if steps and not isinstance(steps, list):
            raise TypeError("Expected argument 'steps' to be a list")
        pulumi.set(__self__, "steps", steps)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def data(self) -> Optional[Mapping[str, Any]]:
        """
        An object that can hold any information about the Form that should be persisted.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="formId")
    def form_id(self) -> str:
        return pulumi.get(self, "form_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The unique name of the Form.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def steps(self) -> Optional[Sequence['outputs.GetFormStepResult']]:
        """
        An ordered list of objects containing one or more Form Fields.
        """
        return pulumi.get(self, "steps")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        The form type. The possible values are:
        """
        return pulumi.get(self, "type")


class AwaitableGetFormResult(GetFormResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFormResult(
            data=self.data,
            form_id=self.form_id,
            id=self.id,
            name=self.name,
            steps=self.steps,
            type=self.type)


def get_form(data: Optional[Mapping[str, Any]] = None,
             form_id: Optional[str] = None,
             name: Optional[str] = None,
             steps: Optional[Sequence[pulumi.InputType['GetFormStepArgs']]] = None,
             type: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFormResult:
    """
    ## # Form Resource

    A FusionAuth Form is a customizable object that contains one-to-many ordered steps. Each step is comprised of one or more Form Fields.

    [Forms API](https://fusionauth.io/docs/v1/tech/apis/forms)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    default = fusionauth.get_form(name="Default User Self Service provided by FusionAuth")
    ```


    :param Mapping[str, Any] data: An object that can hold any information about the Form that should be persisted.
    :param str form_id: The unique id of the Form. Either `form_id` or `name` must be specified.
    :param str name: The name of the Form. Either `form_id` or `name` must be specified.
    :param Sequence[pulumi.InputType['GetFormStepArgs']] steps: An ordered list of objects containing one or more Form Fields.
    :param str type: The form type. The possible values are:
    """
    __args__ = dict()
    __args__['data'] = data
    __args__['formId'] = form_id
    __args__['name'] = name
    __args__['steps'] = steps
    __args__['type'] = type
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
    __ret__ = pulumi.runtime.invoke('fusionauth:index/getForm:getForm', __args__, opts=opts, typ=GetFormResult).value

    return AwaitableGetFormResult(
        data=__ret__.data,
        form_id=__ret__.form_id,
        id=__ret__.id,
        name=__ret__.name,
        steps=__ret__.steps,
        type=__ret__.type)


@_utilities.lift_output_func(get_form)
def get_form_output(data: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                    form_id: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    steps: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetFormStepArgs']]]]] = None,
                    type: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFormResult]:
    """
    ## # Form Resource

    A FusionAuth Form is a customizable object that contains one-to-many ordered steps. Each step is comprised of one or more Form Fields.

    [Forms API](https://fusionauth.io/docs/v1/tech/apis/forms)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_fusionauth as fusionauth

    default = fusionauth.get_form(name="Default User Self Service provided by FusionAuth")
    ```


    :param Mapping[str, Any] data: An object that can hold any information about the Form that should be persisted.
    :param str form_id: The unique id of the Form. Either `form_id` or `name` must be specified.
    :param str name: The name of the Form. Either `form_id` or `name` must be specified.
    :param Sequence[pulumi.InputType['GetFormStepArgs']] steps: An ordered list of objects containing one or more Form Fields.
    :param str type: The form type. The possible values are:
    """
    ...
