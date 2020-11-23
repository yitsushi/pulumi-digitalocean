# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['DatabaseReplica']


class DatabaseReplica(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_network_uuid: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a DigitalOcean database replica resource.

        ## Example Usage
        ### Create a new PostgreSQL database replica
        ```python
        import pulumi
        import pulumi_digitalocean as digitalocean

        postgres_example = digitalocean.DatabaseCluster("postgres-example",
            engine="pg",
            version="11",
            size="db-s-1vcpu-1gb",
            region="nyc1",
            node_count=1)
        read_replica = digitalocean.DatabaseReplica("read-replica",
            cluster_id=postgres_example.id,
            size="db-s-1vcpu-1gb",
            region="nyc1")
        ```

        ## Import

        Database replicas can be imported using the `id` of the source database cluster and the `name` of the replica joined with a comma. For example

        ```sh
         $ pulumi import digitalocean:index/databaseReplica:DatabaseReplica read-replica 245bcfd0-7f31-4ce6-a2bc-475a116cca97,read-replica
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The ID of the original source database cluster.
        :param pulumi.Input[str] name: The name for the database replica.
        :param pulumi.Input[str] region: DigitalOcean region where the replica will reside.
        :param pulumi.Input[str] size: Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if cluster_id is None:
                raise TypeError("Missing required property 'cluster_id'")
            __props__['cluster_id'] = cluster_id
            __props__['name'] = name
            __props__['private_network_uuid'] = private_network_uuid
            __props__['region'] = region
            __props__['size'] = size
            __props__['tags'] = tags
            __props__['database'] = None
            __props__['host'] = None
            __props__['password'] = None
            __props__['port'] = None
            __props__['private_host'] = None
            __props__['private_uri'] = None
            __props__['uri'] = None
            __props__['user'] = None
        super(DatabaseReplica, __self__).__init__(
            'digitalocean:index/databaseReplica:DatabaseReplica',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            database: Optional[pulumi.Input[str]] = None,
            host: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            private_host: Optional[pulumi.Input[str]] = None,
            private_network_uuid: Optional[pulumi.Input[str]] = None,
            private_uri: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            uri: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'DatabaseReplica':
        """
        Get an existing DatabaseReplica resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The ID of the original source database cluster.
        :param pulumi.Input[str] database: Name of the replica's default database.
        :param pulumi.Input[str] host: Database replica's hostname.
        :param pulumi.Input[str] name: The name for the database replica.
        :param pulumi.Input[str] password: Password for the replica's default user.
        :param pulumi.Input[int] port: Network port that the database replica is listening on.
        :param pulumi.Input[str] private_host: Same as `host`, but only accessible from resources within the account and in the same region.
        :param pulumi.Input[str] private_uri: Same as `uri`, but only accessible from resources within the account and in the same region.
        :param pulumi.Input[str] region: DigitalOcean region where the replica will reside.
        :param pulumi.Input[str] size: Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
        :param pulumi.Input[str] uri: The full URI for connecting to the database replica.
        :param pulumi.Input[str] user: Username for the replica's default user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cluster_id"] = cluster_id
        __props__["database"] = database
        __props__["host"] = host
        __props__["name"] = name
        __props__["password"] = password
        __props__["port"] = port
        __props__["private_host"] = private_host
        __props__["private_network_uuid"] = private_network_uuid
        __props__["private_uri"] = private_uri
        __props__["region"] = region
        __props__["size"] = size
        __props__["tags"] = tags
        __props__["uri"] = uri
        __props__["user"] = user
        return DatabaseReplica(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The ID of the original source database cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output[str]:
        """
        Name of the replica's default database.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def host(self) -> pulumi.Output[str]:
        """
        Database replica's hostname.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name for the database replica.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        Password for the replica's default user.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        Network port that the database replica is listening on.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="privateHost")
    def private_host(self) -> pulumi.Output[str]:
        """
        Same as `host`, but only accessible from resources within the account and in the same region.
        """
        return pulumi.get(self, "private_host")

    @property
    @pulumi.getter(name="privateNetworkUuid")
    def private_network_uuid(self) -> pulumi.Output[str]:
        return pulumi.get(self, "private_network_uuid")

    @property
    @pulumi.getter(name="privateUri")
    def private_uri(self) -> pulumi.Output[str]:
        """
        Same as `uri`, but only accessible from resources within the account and in the same region.
        """
        return pulumi.get(self, "private_uri")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        DigitalOcean region where the replica will reside.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[Optional[str]]:
        """
        Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def uri(self) -> pulumi.Output[str]:
        """
        The full URI for connecting to the database replica.
        """
        return pulumi.get(self, "uri")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        Username for the replica's default user.
        """
        return pulumi.get(self, "user")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

