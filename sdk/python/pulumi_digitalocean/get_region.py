# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetRegionResult:
    """
    A collection of values returned by getRegion.
    """
    def __init__(__self__, available=None, features=None, id=None, name=None, sizes=None, slug=None):
        if available and not isinstance(available, bool):
            raise TypeError("Expected argument 'available' to be a bool")
        __self__.available = available
        """
        A boolean value that represents whether new Droplets can be created in this region.
        """
        if features and not isinstance(features, list):
            raise TypeError("Expected argument 'features' to be a list")
        __self__.features = features
        """
        A set of features available in this region.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The display name of the region.
        """
        if sizes and not isinstance(sizes, list):
            raise TypeError("Expected argument 'sizes' to be a list")
        __self__.sizes = sizes
        """
        A set of identifying slugs for the Droplet sizes available in this region.
        """
        if slug and not isinstance(slug, str):
            raise TypeError("Expected argument 'slug' to be a str")
        __self__.slug = slug
        """
        A human-readable string that is used as a unique identifier for each region.
        """
class AwaitableGetRegionResult(GetRegionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionResult(
            available=self.available,
            features=self.features,
            id=self.id,
            name=self.name,
            sizes=self.sizes,
            slug=self.slug)

def get_region(slug=None,opts=None):
    """
    Get information on a single DigitalOcean region. This is useful to find out 
    what Droplet sizes and features are supported within a region.



    > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/region.html.md.


    :param str slug: A human-readable string that is used as a unique identifier for each region.
    """
    __args__ = dict()


    __args__['slug'] = slug
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('digitalocean:index/getRegion:getRegion', __args__, opts=opts).value

    return AwaitableGetRegionResult(
        available=__ret__.get('available'),
        features=__ret__.get('features'),
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        sizes=__ret__.get('sizes'),
        slug=__ret__.get('slug'))
