// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    public static partial class Invokes
    {
        /// <summary>
        /// Provides information on a DigitalOcean database cluster resource.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/database_cluster.html.markdown.
        /// </summary>
        [Obsolete("Use GetDatabaseCluster.InvokeAsync() instead")]
        public static Task<GetDatabaseClusterResult> GetDatabaseCluster(GetDatabaseClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseClusterResult>("digitalocean:index/getDatabaseCluster:getDatabaseCluster", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetDatabaseCluster
    {
        /// <summary>
        /// Provides information on a DigitalOcean database cluster resource.
        /// 
        /// 
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/d/database_cluster.html.markdown.
        /// </summary>
        public static Task<GetDatabaseClusterResult> InvokeAsync(GetDatabaseClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseClusterResult>("digitalocean:index/getDatabaseCluster:getDatabaseCluster", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetDatabaseClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database cluster.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("tags")]
        private List<string>? _tags;
        public List<string> Tags
        {
            get => _tags ?? (_tags = new List<string>());
            set => _tags = value;
        }

        public GetDatabaseClusterArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetDatabaseClusterResult
    {
        /// <summary>
        /// Name of the cluster's default database.
        /// </summary>
        public readonly string Database;
        /// <summary>
        /// Database engine used by the cluster (ex. `pg` for PostreSQL).
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// Database cluster's hostname.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Defines when the automatic maintenance should be performed for the database cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDatabaseClusterMaintenanceWindowsResult> MaintenanceWindows;
        public readonly string Name;
        /// <summary>
        /// Number of nodes that will be included in the cluster.
        /// </summary>
        public readonly int NodeCount;
        /// <summary>
        /// Password for the cluster's default user.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Network port that the database cluster is listening on.
        /// </summary>
        public readonly int Port;
        /// <summary>
        /// Same as `host`, but only accessible from resources within the account and in the same region.
        /// </summary>
        public readonly string PrivateHost;
        /// <summary>
        /// Same as `uri`, but only accessible from resources within the account and in the same region.
        /// </summary>
        public readonly string PrivateUri;
        /// <summary>
        /// DigitalOcean region where the cluster will reside.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Database droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
        /// </summary>
        public readonly string Size;
        public readonly ImmutableArray<string> Tags;
        /// <summary>
        /// The full URI for connecting to the database cluster.
        /// </summary>
        public readonly string Uri;
        /// <summary>
        /// The uniform resource name of the database cluster.
        /// </summary>
        public readonly string Urn;
        /// <summary>
        /// Username for the cluster's default user.
        /// </summary>
        public readonly string User;
        /// <summary>
        /// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
        /// </summary>
        public readonly string Version;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetDatabaseClusterResult(
            string database,
            string engine,
            string host,
            ImmutableArray<Outputs.GetDatabaseClusterMaintenanceWindowsResult> maintenanceWindows,
            string name,
            int nodeCount,
            string password,
            int port,
            string privateHost,
            string privateUri,
            string region,
            string size,
            ImmutableArray<string> tags,
            string uri,
            string urn,
            string user,
            string version,
            string id)
        {
            Database = database;
            Engine = engine;
            Host = host;
            MaintenanceWindows = maintenanceWindows;
            Name = name;
            NodeCount = nodeCount;
            Password = password;
            Port = port;
            PrivateHost = privateHost;
            PrivateUri = privateUri;
            Region = region;
            Size = size;
            Tags = tags;
            Uri = uri;
            Urn = urn;
            User = user;
            Version = version;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetDatabaseClusterMaintenanceWindowsResult
    {
        /// <summary>
        /// The day of the week on which to apply maintenance updates.
        /// </summary>
        public readonly string Day;
        /// <summary>
        /// The hour in UTC at which maintenance updates will be applied in 24 hour format.
        /// </summary>
        public readonly string Hour;

        [OutputConstructor]
        private GetDatabaseClusterMaintenanceWindowsResult(
            string day,
            string hour)
        {
            Day = day;
            Hour = hour;
        }
    }
    }
}
