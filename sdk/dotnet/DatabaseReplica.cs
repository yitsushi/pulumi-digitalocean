// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    /// <summary>
    /// Provides a DigitalOcean database replica resource.
    /// 
    /// ## Example Usage
    /// ### Create a new PostgreSQL database replica
    /// ```csharp
    /// using Pulumi;
    /// using DigitalOcean = Pulumi.DigitalOcean;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var postgres_example = new DigitalOcean.DatabaseCluster("postgres-example", new DigitalOcean.DatabaseClusterArgs
    ///         {
    ///             Engine = "pg",
    ///             Version = "11",
    ///             Size = "db-s-1vcpu-1gb",
    ///             Region = "nyc1",
    ///             NodeCount = 1,
    ///         });
    ///         var read_replica = new DigitalOcean.DatabaseReplica("read-replica", new DigitalOcean.DatabaseReplicaArgs
    ///         {
    ///             ClusterId = postgres_example.Id,
    ///             Size = "db-s-1vcpu-1gb",
    ///             Region = "nyc1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database replicas can be imported using the `id` of the source database cluster and the `name` of the replica joined with a comma. For example
    /// 
    /// ```sh
    ///  $ pulumi import digitalocean:index/databaseReplica:DatabaseReplica read-replica 245bcfd0-7f31-4ce6-a2bc-475a116cca97,read-replica
    /// ```
    /// </summary>
    [DigitalOceanResourceType("digitalocean:index/databaseReplica:DatabaseReplica")]
    public partial class DatabaseReplica : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the original source database cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Name of the replica's default database.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// Database replica's hostname.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// The name for the database replica.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password for the replica's default user.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Network port that the database replica is listening on.
        /// </summary>
        [Output("port")]
        public Output<int> Port { get; private set; } = null!;

        /// <summary>
        /// Same as `host`, but only accessible from resources within the account and in the same region.
        /// </summary>
        [Output("privateHost")]
        public Output<string> PrivateHost { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPC where the database replica will be located.
        /// </summary>
        [Output("privateNetworkUuid")]
        public Output<string> PrivateNetworkUuid { get; private set; } = null!;

        /// <summary>
        /// Same as `uri`, but only accessible from resources within the account and in the same region.
        /// </summary>
        [Output("privateUri")]
        public Output<string> PrivateUri { get; private set; } = null!;

        /// <summary>
        /// DigitalOcean region where the replica will reside.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
        /// </summary>
        [Output("size")]
        public Output<string?> Size { get; private set; } = null!;

        /// <summary>
        /// A list of tag names to be applied to the database replica.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        /// <summary>
        /// The full URI for connecting to the database replica.
        /// </summary>
        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;

        /// <summary>
        /// Username for the replica's default user.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseReplica resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseReplica(string name, DatabaseReplicaArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/databaseReplica:DatabaseReplica", name, args ?? new DatabaseReplicaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseReplica(string name, Input<string> id, DatabaseReplicaState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/databaseReplica:DatabaseReplica", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DatabaseReplica resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseReplica Get(string name, Input<string> id, DatabaseReplicaState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseReplica(name, id, state, options);
        }
    }

    public sealed class DatabaseReplicaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the original source database cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The name for the database replica.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the VPC where the database replica will be located.
        /// </summary>
        [Input("privateNetworkUuid")]
        public Input<string>? PrivateNetworkUuid { get; set; }

        /// <summary>
        /// DigitalOcean region where the replica will reside.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag names to be applied to the database replica.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public DatabaseReplicaArgs()
        {
        }
    }

    public sealed class DatabaseReplicaState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the original source database cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Name of the replica's default database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// Database replica's hostname.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The name for the database replica.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password for the replica's default user.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Network port that the database replica is listening on.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Same as `host`, but only accessible from resources within the account and in the same region.
        /// </summary>
        [Input("privateHost")]
        public Input<string>? PrivateHost { get; set; }

        /// <summary>
        /// The ID of the VPC where the database replica will be located.
        /// </summary>
        [Input("privateNetworkUuid")]
        public Input<string>? PrivateNetworkUuid { get; set; }

        /// <summary>
        /// Same as `uri`, but only accessible from resources within the account and in the same region.
        /// </summary>
        [Input("privateUri")]
        public Input<string>? PrivateUri { get; set; }

        /// <summary>
        /// DigitalOcean region where the replica will reside.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// A list of tag names to be applied to the database replica.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The full URI for connecting to the database replica.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        /// <summary>
        /// Username for the replica's default user.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public DatabaseReplicaState()
        {
        }
    }
}
