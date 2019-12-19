# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Project(pulumi.CustomResource):
    created_at: pulumi.Output[str]
    """
    the date and time when the project was created, (ISO8601)
    """
    description: pulumi.Output[str]
    """
    the description of the project
    """
    environment: pulumi.Output[str]
    """
    the environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`)
    """
    name: pulumi.Output[str]
    """
    The name of the Project
    """
    owner_id: pulumi.Output[float]
    """
    the id of the project owner.
    """
    owner_uuid: pulumi.Output[str]
    """
    the unique universal identifier of the project owner.
    """
    purpose: pulumi.Output[str]
    """
    the purpose of the project, (Default: "Web Application")
    """
    resources: pulumi.Output[list]
    """
    a list of uniform resource names (URNs) for the resources associated with the project
    """
    updated_at: pulumi.Output[str]
    """
    the date and time when the project was last updated, (ISO8601)
    """
    def __init__(__self__, resource_name, opts=None, description=None, environment=None, name=None, purpose=None, resources=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a DigitalOcean Project resource.
        
        Projects allow you to organize your resources into groups that fit the way you work.
        You can group resources (like Droplets, Spaces, Load Balancers, domains, and Floating IPs)
        in ways that align with the applications you host on DigitalOcean.
        
        The following resource types can be associated with a project:
        
        * Database Clusters
        * Domains
        * Droplets
        * Floating IP
        * Load Balancers
        * Spaces Bucket
        * Volume
        
        **Note:** A Terrafrom managed project cannot be set as a default project.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: the description of the project
        :param pulumi.Input[str] environment: the environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`)
        :param pulumi.Input[str] name: The name of the Project
        :param pulumi.Input[str] purpose: the purpose of the project, (Default: "Web Application")
        :param pulumi.Input[list] resources: a list of uniform resource names (URNs) for the resources associated with the project

        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/project.html.markdown.
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

            __props__['description'] = description
            __props__['environment'] = environment
            __props__['name'] = name
            __props__['purpose'] = purpose
            __props__['resources'] = resources
            __props__['created_at'] = None
            __props__['owner_id'] = None
            __props__['owner_uuid'] = None
            __props__['updated_at'] = None
        super(Project, __self__).__init__(
            'digitalocean:index/project:Project',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, created_at=None, description=None, environment=None, name=None, owner_id=None, owner_uuid=None, purpose=None, resources=None, updated_at=None):
        """
        Get an existing Project resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: the date and time when the project was created, (ISO8601)
        :param pulumi.Input[str] description: the description of the project
        :param pulumi.Input[str] environment: the environment of the project's resources. The possible values are: `Development`, `Staging`, `Production`)
        :param pulumi.Input[str] name: The name of the Project
        :param pulumi.Input[float] owner_id: the id of the project owner.
        :param pulumi.Input[str] owner_uuid: the unique universal identifier of the project owner.
        :param pulumi.Input[str] purpose: the purpose of the project, (Default: "Web Application")
        :param pulumi.Input[list] resources: a list of uniform resource names (URNs) for the resources associated with the project
        :param pulumi.Input[str] updated_at: the date and time when the project was last updated, (ISO8601)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/project.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["created_at"] = created_at
        __props__["description"] = description
        __props__["environment"] = environment
        __props__["name"] = name
        __props__["owner_id"] = owner_id
        __props__["owner_uuid"] = owner_uuid
        __props__["purpose"] = purpose
        __props__["resources"] = resources
        __props__["updated_at"] = updated_at
        return Project(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

