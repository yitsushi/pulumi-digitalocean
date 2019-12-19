# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetKubernetesVersionsResult:
    """
    A collection of values returned by getKubernetesVersions.
    """
    def __init__(__self__, latest_version=None, valid_versions=None, version_prefix=None, id=None):
        if latest_version and not isinstance(latest_version, str):
            raise TypeError("Expected argument 'latest_version' to be a str")
        __self__.latest_version = latest_version
        """
        The most recent version available.
        """
        if valid_versions and not isinstance(valid_versions, list):
            raise TypeError("Expected argument 'valid_versions' to be a list")
        __self__.valid_versions = valid_versions
        """
        A list of available versions.
        """
        if version_prefix and not isinstance(version_prefix, str):
            raise TypeError("Expected argument 'version_prefix' to be a str")
        __self__.version_prefix = version_prefix
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetKubernetesVersionsResult(GetKubernetesVersionsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKubernetesVersionsResult(
            latest_version=self.latest_version,
            valid_versions=self.valid_versions,
            version_prefix=self.version_prefix,
            id=self.id)

def get_kubernetes_versions(version_prefix=None,opts=None):
    """
    Provides access to the available DigitalOcean Kubernetes Service versions.
    

    > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/kubernetes_versions.html.markdown.
    """
    __args__ = dict()

    __args__['versionPrefix'] = version_prefix
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getKubernetesVersions:getKubernetesVersions', __args__, opts=opts).value

    return AwaitableGetKubernetesVersionsResult(
        latest_version=__ret__.get('latestVersion'),
        valid_versions=__ret__.get('validVersions'),
        version_prefix=__ret__.get('versionPrefix'),
        id=__ret__.get('id'))
