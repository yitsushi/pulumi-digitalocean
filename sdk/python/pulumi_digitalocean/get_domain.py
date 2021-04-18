# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetDomainResult',
    'AwaitableGetDomainResult',
    'get_domain',
]

@pulumi.output_type
class GetDomainResult:
    """
    A collection of values returned by getDomain.
    """
    def __init__(__self__, domain_urn=None, id=None, name=None, ttl=None, zone_file=None):
        if domain_urn and not isinstance(domain_urn, str):
            raise TypeError("Expected argument 'domain_urn' to be a str")
        pulumi.set(__self__, "domain_urn", domain_urn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if ttl and not isinstance(ttl, int):
            raise TypeError("Expected argument 'ttl' to be a int")
        pulumi.set(__self__, "ttl", ttl)
        if zone_file and not isinstance(zone_file, str):
            raise TypeError("Expected argument 'zone_file' to be a str")
        pulumi.set(__self__, "zone_file", zone_file)

    @property
    @pulumi.getter(name="domainUrn")
    def domain_urn(self) -> str:
        """
        The uniform resource name of the domain
        * `zone_file`: The zone file of the domain.
        """
        return pulumi.get(self, "domain_urn")

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
    def ttl(self) -> int:
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="zoneFile")
    def zone_file(self) -> str:
        return pulumi.get(self, "zone_file")


class AwaitableGetDomainResult(GetDomainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDomainResult(
            domain_urn=self.domain_urn,
            id=self.id,
            name=self.name,
            ttl=self.ttl,
            zone_file=self.zone_file)


def get_domain(name: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDomainResult:
    """
    Get information on a domain. This data source provides the name, TTL, and zone
    file as configured on your DigitalOcean account. This is useful if the domain
    name in question is not managed by this provider or you need to utilize TTL or zone
    file data.

    An error is triggered if the provided domain name is not managed with your
    DigitalOcean account.


    :param str name: The name of the domain.
    """
    __args__ = dict()
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getDomain:getDomain', __args__, opts=opts, typ=GetDomainResult).value

    return AwaitableGetDomainResult(
        domain_urn=__ret__.domain_urn,
        id=__ret__.id,
        name=__ret__.name,
        ttl=__ret__.ttl,
        zone_file=__ret__.zone_file)
