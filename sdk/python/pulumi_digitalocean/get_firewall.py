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
    'GetFirewallResult',
    'AwaitableGetFirewallResult',
    'get_firewall',
]

@pulumi.output_type
class GetFirewallResult:
    """
    A collection of values returned by getFirewall.
    """
    def __init__(__self__, created_at=None, droplet_ids=None, firewall_id=None, id=None, inbound_rules=None, name=None, outbound_rules=None, pending_changes=None, status=None, tags=None):
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        pulumi.set(__self__, "created_at", created_at)
        if droplet_ids and not isinstance(droplet_ids, list):
            raise TypeError("Expected argument 'droplet_ids' to be a list")
        pulumi.set(__self__, "droplet_ids", droplet_ids)
        if firewall_id and not isinstance(firewall_id, str):
            raise TypeError("Expected argument 'firewall_id' to be a str")
        pulumi.set(__self__, "firewall_id", firewall_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inbound_rules and not isinstance(inbound_rules, list):
            raise TypeError("Expected argument 'inbound_rules' to be a list")
        pulumi.set(__self__, "inbound_rules", inbound_rules)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if outbound_rules and not isinstance(outbound_rules, list):
            raise TypeError("Expected argument 'outbound_rules' to be a list")
        pulumi.set(__self__, "outbound_rules", outbound_rules)
        if pending_changes and not isinstance(pending_changes, list):
            raise TypeError("Expected argument 'pending_changes' to be a list")
        pulumi.set(__self__, "pending_changes", pending_changes)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        """
        A time value given in ISO8601 combined date and time format
        that represents when the Firewall was created.
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="dropletIds")
    def droplet_ids(self) -> Sequence[int]:
        """
        The list of the IDs of the Droplets assigned to
        the Firewall.
        """
        return pulumi.get(self, "droplet_ids")

    @property
    @pulumi.getter(name="firewallId")
    def firewall_id(self) -> str:
        return pulumi.get(self, "firewall_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inboundRules")
    def inbound_rules(self) -> Sequence['outputs.GetFirewallInboundRuleResult']:
        return pulumi.get(self, "inbound_rules")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the Firewall.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outboundRules")
    def outbound_rules(self) -> Sequence['outputs.GetFirewallOutboundRuleResult']:
        return pulumi.get(self, "outbound_rules")

    @property
    @pulumi.getter(name="pendingChanges")
    def pending_changes(self) -> Sequence['outputs.GetFirewallPendingChangeResult']:
        """
        A set of object containing the fields, `droplet_id`,
        `removing`, and `status`.  It is provided to detail exactly which Droplets
        are having their security policies updated.  When empty, all changes
        have been successfully applied.
        """
        return pulumi.get(self, "pending_changes")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        A status string indicating the current state of the Firewall.
        This can be "waiting", "succeeded", or "failed".
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        The names of the Tags assigned to the Firewall.
        """
        return pulumi.get(self, "tags")


class AwaitableGetFirewallResult(GetFirewallResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFirewallResult(
            created_at=self.created_at,
            droplet_ids=self.droplet_ids,
            firewall_id=self.firewall_id,
            id=self.id,
            inbound_rules=self.inbound_rules,
            name=self.name,
            outbound_rules=self.outbound_rules,
            pending_changes=self.pending_changes,
            status=self.status,
            tags=self.tags)


def get_firewall(droplet_ids: Optional[Sequence[int]] = None,
                 firewall_id: Optional[str] = None,
                 inbound_rules: Optional[Sequence[pulumi.InputType['GetFirewallInboundRuleArgs']]] = None,
                 outbound_rules: Optional[Sequence[pulumi.InputType['GetFirewallOutboundRuleArgs']]] = None,
                 tags: Optional[Sequence[str]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFirewallResult:
    """
    Get information on a DigitalOcean Firewall.

    ## Example Usage

    Get the firewall:

    ```python
    import pulumi
    import pulumi_digitalocean as digitalocean

    example = digitalocean.get_firewall(firewall_id="1df48973-6eef-4214-854f-fa7726e7e583")
    pulumi.export("exampleFirewallName", example.name)
    ```


    :param Sequence[int] droplet_ids: The list of the IDs of the Droplets assigned to
           the Firewall.
    :param str firewall_id: The ID of the firewall to retrieve information
           about.
    :param Sequence[str] tags: The names of the Tags assigned to the Firewall.
    """
    __args__ = dict()
    __args__['dropletIds'] = droplet_ids
    __args__['firewallId'] = firewall_id
    __args__['inboundRules'] = inbound_rules
    __args__['outboundRules'] = outbound_rules
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getFirewall:getFirewall', __args__, opts=opts, typ=GetFirewallResult).value

    return AwaitableGetFirewallResult(
        created_at=__ret__.created_at,
        droplet_ids=__ret__.droplet_ids,
        firewall_id=__ret__.firewall_id,
        id=__ret__.id,
        inbound_rules=__ret__.inbound_rules,
        name=__ret__.name,
        outbound_rules=__ret__.outbound_rules,
        pending_changes=__ret__.pending_changes,
        status=__ret__.status,
        tags=__ret__.tags)
