# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class SpacesBucket(pulumi.CustomResource):
    acl: pulumi.Output[str]
    """
    Canned ACL applied on bucket creation (`private` or `public-read`)
    """
    bucket_domain_name: pulumi.Output[str]
    """
    The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
    """
    cors_rules: pulumi.Output[list]
    """
    A container holding a list of elements describing allowed methods for a specific origin.

      * `allowedHeaders` (`list`) - A list of headers that will be included in the CORS preflight request's `Access-Control-Request-Headers`. A header may contain one wildcard (e.g. `x-amz-*`).
      * `allowedMethods` (`list`) - A list of HTTP methods (e.g. `GET`) which are allowed from the specified origin.
      * `allowedOrigins` (`list`) - A list of hosts from which requests using the specified methods are allowed. A host may contain one wildcard (e.g. http://*.example.com).
      * `maxAgeSeconds` (`float`) - The time in seconds that browser can cache the response for a preflight request.
    """
    force_destroy: pulumi.Output[bool]
    """
    Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
    """
    name: pulumi.Output[str]
    """
    The name of the bucket
    """
    region: pulumi.Output[str]
    """
    The region where the bucket resides (Defaults to `nyc3`)
    """
    urn: pulumi.Output[str]
    """
    The uniform resource name for the bucket
    """
    def __init__(__self__, resource_name, opts=None, acl=None, cors_rules=None, force_destroy=None, name=None, region=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a bucket resource for Spaces, DigitalOcean's object storage product.

        The [Spaces API](https://developers.digitalocean.com/documentation/spaces/) was
        designed to be interoperable with Amazon's AWS S3 API. This allows users to
        interact with the service while using the tools they already know. Spaces
        mirrors S3's authentication framework and requests to Spaces require a key pair
        similar to Amazon's Access ID and Secret Key.

        The authentication requirement can be met by either setting the
        `SPACES_ACCESS_KEY_ID` and `SPACES_SECRET_ACCESS_KEY` environment variables or
        the provider's `spaces_access_id` and `spaces_secret_key` arguments to the
        access ID and secret you generate via the DigitalOcean control panel. For
        example:


        For more information, See [An Introduction to DigitalOcean Spaces](https://www.digitalocean.com/community/tutorials/an-introduction-to-digitalocean-spaces)



        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/spaces_bucket.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl: Canned ACL applied on bucket creation (`private` or `public-read`)
        :param pulumi.Input[list] cors_rules: A container holding a list of elements describing allowed methods for a specific origin.
        :param pulumi.Input[bool] force_destroy: Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
        :param pulumi.Input[str] name: The name of the bucket
        :param pulumi.Input[str] region: The region where the bucket resides (Defaults to `nyc3`)

        The **cors_rules** object supports the following:

          * `allowedHeaders` (`pulumi.Input[list]`) - A list of headers that will be included in the CORS preflight request's `Access-Control-Request-Headers`. A header may contain one wildcard (e.g. `x-amz-*`).
          * `allowedMethods` (`pulumi.Input[list]`) - A list of HTTP methods (e.g. `GET`) which are allowed from the specified origin.
          * `allowedOrigins` (`pulumi.Input[list]`) - A list of hosts from which requests using the specified methods are allowed. A host may contain one wildcard (e.g. http://*.example.com).
          * `maxAgeSeconds` (`pulumi.Input[float]`) - The time in seconds that browser can cache the response for a preflight request.
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

            __props__['acl'] = acl
            __props__['cors_rules'] = cors_rules
            __props__['force_destroy'] = force_destroy
            __props__['name'] = name
            __props__['region'] = region
            __props__['bucket_domain_name'] = None
            __props__['urn'] = None
        super(SpacesBucket, __self__).__init__(
            'digitalocean:index/spacesBucket:SpacesBucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, acl=None, bucket_domain_name=None, cors_rules=None, force_destroy=None, name=None, region=None, urn=None):
        """
        Get an existing SpacesBucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] acl: Canned ACL applied on bucket creation (`private` or `public-read`)
        :param pulumi.Input[str] bucket_domain_name: The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
        :param pulumi.Input[list] cors_rules: A container holding a list of elements describing allowed methods for a specific origin.
        :param pulumi.Input[bool] force_destroy: Unless `true`, the bucket will only be destroyed if empty (Defaults to `false`)
        :param pulumi.Input[str] name: The name of the bucket
        :param pulumi.Input[str] region: The region where the bucket resides (Defaults to `nyc3`)
        :param pulumi.Input[str] urn: The uniform resource name for the bucket

        The **cors_rules** object supports the following:

          * `allowedHeaders` (`pulumi.Input[list]`) - A list of headers that will be included in the CORS preflight request's `Access-Control-Request-Headers`. A header may contain one wildcard (e.g. `x-amz-*`).
          * `allowedMethods` (`pulumi.Input[list]`) - A list of HTTP methods (e.g. `GET`) which are allowed from the specified origin.
          * `allowedOrigins` (`pulumi.Input[list]`) - A list of hosts from which requests using the specified methods are allowed. A host may contain one wildcard (e.g. http://*.example.com).
          * `maxAgeSeconds` (`pulumi.Input[float]`) - The time in seconds that browser can cache the response for a preflight request.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["acl"] = acl
        __props__["bucket_domain_name"] = bucket_domain_name
        __props__["cors_rules"] = cors_rules
        __props__["force_destroy"] = force_destroy
        __props__["name"] = name
        __props__["region"] = region
        __props__["urn"] = urn
        return SpacesBucket(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

