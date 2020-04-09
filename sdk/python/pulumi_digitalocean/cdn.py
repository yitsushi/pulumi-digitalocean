# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Cdn(pulumi.CustomResource):
    certificate_id: pulumi.Output[str]
    """
    ID of a DigitalOcean managed TLS certificate for use with custom domains
    """
    created_at: pulumi.Output[str]
    """
    The date and time when the CDN Endpoint was created.
    """
    custom_domain: pulumi.Output[str]
    """
    The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint.
    """
    endpoint: pulumi.Output[str]
    """
    The fully qualified domain name (FQDN) from which the CDN-backed content is served.
    """
    origin: pulumi.Output[str]
    """
    The fully qualified domain name, (FQDN) for a Space.
    """
    ttl: pulumi.Output[float]
    """
    The time to live for the CDN Endpoint, in seconds. Default is 3600 seconds.
    * `certificate_id`- (Optional) The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
    """
    def __init__(__self__, resource_name, opts=None, certificate_id=None, custom_domain=None, origin=None, ttl=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DigitalOcean CDN Endpoint resource for use with Spaces.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/cdn.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_id: ID of a DigitalOcean managed TLS certificate for use with custom domains
        :param pulumi.Input[str] custom_domain: The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint.
        :param pulumi.Input[str] origin: The fully qualified domain name, (FQDN) for a Space.
        :param pulumi.Input[float] ttl: The time to live for the CDN Endpoint, in seconds. Default is 3600 seconds.
               * `certificate_id`- (Optional) The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['certificate_id'] = certificate_id
            __props__['custom_domain'] = custom_domain
            if origin is None:
                raise TypeError("Missing required property 'origin'")
            __props__['origin'] = origin
            __props__['ttl'] = ttl
            __props__['created_at'] = None
            __props__['endpoint'] = None
        super(Cdn, __self__).__init__(
            'digitalocean:index/cdn:Cdn',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, certificate_id=None, created_at=None, custom_domain=None, endpoint=None, origin=None, ttl=None):
        """
        Get an existing Cdn resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_id: ID of a DigitalOcean managed TLS certificate for use with custom domains
        :param pulumi.Input[str] created_at: The date and time when the CDN Endpoint was created.
        :param pulumi.Input[str] custom_domain: The fully qualified domain name (FQDN) of the custom subdomain used with the CDN Endpoint.
        :param pulumi.Input[str] endpoint: The fully qualified domain name (FQDN) from which the CDN-backed content is served.
        :param pulumi.Input[str] origin: The fully qualified domain name, (FQDN) for a Space.
        :param pulumi.Input[float] ttl: The time to live for the CDN Endpoint, in seconds. Default is 3600 seconds.
               * `certificate_id`- (Optional) The ID of a DigitalOcean managed TLS certificate used for SSL when a custom subdomain is provided.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate_id"] = certificate_id
        __props__["created_at"] = created_at
        __props__["custom_domain"] = custom_domain
        __props__["endpoint"] = endpoint
        __props__["origin"] = origin
        __props__["ttl"] = ttl
        return Cdn(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

