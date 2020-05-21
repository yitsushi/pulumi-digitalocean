// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

import {DatabaseSlug, Region} from "./index";

/**
 * Provides a DigitalOcean database cluster resource.
 *
 * ## Example Usage
 *
 * ### Create a new PostgreSQL database cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const postgresExample = new digitalocean.DatabaseCluster("postgres-example", {
 *     engine: "pg",
 *     nodeCount: 1,
 *     region: "nyc1",
 *     size: "db-s-1vcpu-1gb",
 *     version: "11",
 * });
 * ```
 *
 * ### Create a new MySQL database cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const mysqlExample = new digitalocean.DatabaseCluster("mysql-example", {
 *     engine: "mysql",
 *     nodeCount: 1,
 *     region: "nyc1",
 *     size: "db-s-1vcpu-1gb",
 *     version: "8",
 * });
 * ```
 *
 * ### Create a new Redis database cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const redisExample = new digitalocean.DatabaseCluster("redis-example", {
 *     engine: "redis",
 *     nodeCount: 1,
 *     region: "nyc1",
 *     size: "db-s-1vcpu-1gb",
 *     version: "5",
 * });
 * ```
 */
export class DatabaseCluster extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseClusterState, opts?: pulumi.CustomResourceOptions): DatabaseCluster {
        return new DatabaseCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/databaseCluster:DatabaseCluster';

    /**
     * Returns true if the given object is an instance of DatabaseCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseCluster.__pulumiType;
    }

    /**
     * Name of the cluster's default database.
     */
    public /*out*/ readonly database!: pulumi.Output<string>;
    /**
     * Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
     */
    public readonly evictionPolicy!: pulumi.Output<string | undefined>;
    /**
     * Database cluster's hostname.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * Defines when the automatic maintenance should be performed for the database cluster.
     */
    public readonly maintenanceWindows!: pulumi.Output<outputs.DatabaseClusterMaintenanceWindow[] | undefined>;
    /**
     * The name of the database cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of nodes that will be included in the cluster.
     */
    public readonly nodeCount!: pulumi.Output<number>;
    /**
     * Password for the cluster's default user.
     */
    public /*out*/ readonly password!: pulumi.Output<string>;
    /**
     * Network port that the database cluster is listening on.
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * Same as `host`, but only accessible from resources within the account and in the same region.
     */
    public /*out*/ readonly privateHost!: pulumi.Output<string>;
    /**
     * The ID of the VPC where the database cluster will be located.
     */
    public readonly privateNetworkUuid!: pulumi.Output<string>;
    /**
     * Same as `uri`, but only accessible from resources within the account and in the same region.
     */
    public /*out*/ readonly privateUri!: pulumi.Output<string>;
    /**
     * DigitalOcean region where the cluster will reside.
     */
    public readonly region!: pulumi.Output<Region>;
    /**
     * Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
     */
    public readonly size!: pulumi.Output<DatabaseSlug>;
    /**
     * A comma separated string specifying the  SQL modes for a MySQL cluster.
     */
    public readonly sqlMode!: pulumi.Output<string | undefined>;
    /**
     * A list of tag names to be applied to the database cluster.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The full URI for connecting to the database cluster.
     */
    public /*out*/ readonly uri!: pulumi.Output<string>;
    /**
     * The uniform resource name of the database cluster.
     */
    public /*out*/ readonly urn!: pulumi.Output<string>;
    /**
     * Username for the cluster's default user.
     */
    public /*out*/ readonly user!: pulumi.Output<string>;
    /**
     * Engine version used by the cluster (ex. `11` for PostgreSQL 11).
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a DatabaseCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseClusterArgs | DatabaseClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseClusterState | undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["engine"] = state ? state.engine : undefined;
            inputs["evictionPolicy"] = state ? state.evictionPolicy : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["maintenanceWindows"] = state ? state.maintenanceWindows : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeCount"] = state ? state.nodeCount : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["privateHost"] = state ? state.privateHost : undefined;
            inputs["privateNetworkUuid"] = state ? state.privateNetworkUuid : undefined;
            inputs["privateUri"] = state ? state.privateUri : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["sqlMode"] = state ? state.sqlMode : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["uri"] = state ? state.uri : undefined;
            inputs["urn"] = state ? state.urn : undefined;
            inputs["user"] = state ? state.user : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as DatabaseClusterArgs | undefined;
            if (!args || args.engine === undefined) {
                throw new Error("Missing required property 'engine'");
            }
            if (!args || args.nodeCount === undefined) {
                throw new Error("Missing required property 'nodeCount'");
            }
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            if (!args || args.size === undefined) {
                throw new Error("Missing required property 'size'");
            }
            inputs["engine"] = args ? args.engine : undefined;
            inputs["evictionPolicy"] = args ? args.evictionPolicy : undefined;
            inputs["maintenanceWindows"] = args ? args.maintenanceWindows : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeCount"] = args ? args.nodeCount : undefined;
            inputs["privateNetworkUuid"] = args ? args.privateNetworkUuid : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["sqlMode"] = args ? args.sqlMode : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["database"] = undefined /*out*/;
            inputs["host"] = undefined /*out*/;
            inputs["password"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["privateHost"] = undefined /*out*/;
            inputs["privateUri"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
            inputs["urn"] = undefined /*out*/;
            inputs["user"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatabaseCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseCluster resources.
 */
export interface DatabaseClusterState {
    /**
     * Name of the cluster's default database.
     */
    readonly database?: pulumi.Input<string>;
    /**
     * Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
     */
    readonly engine?: pulumi.Input<string>;
    /**
     * A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
     */
    readonly evictionPolicy?: pulumi.Input<string>;
    /**
     * Database cluster's hostname.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * Defines when the automatic maintenance should be performed for the database cluster.
     */
    readonly maintenanceWindows?: pulumi.Input<pulumi.Input<inputs.DatabaseClusterMaintenanceWindow>[]>;
    /**
     * The name of the database cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Number of nodes that will be included in the cluster.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * Password for the cluster's default user.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Network port that the database cluster is listening on.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Same as `host`, but only accessible from resources within the account and in the same region.
     */
    readonly privateHost?: pulumi.Input<string>;
    /**
     * The ID of the VPC where the database cluster will be located.
     */
    readonly privateNetworkUuid?: pulumi.Input<string>;
    /**
     * Same as `uri`, but only accessible from resources within the account and in the same region.
     */
    readonly privateUri?: pulumi.Input<string>;
    /**
     * DigitalOcean region where the cluster will reside.
     */
    readonly region?: pulumi.Input<Region>;
    /**
     * Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
     */
    readonly size?: pulumi.Input<DatabaseSlug>;
    /**
     * A comma separated string specifying the  SQL modes for a MySQL cluster.
     */
    readonly sqlMode?: pulumi.Input<string>;
    /**
     * A list of tag names to be applied to the database cluster.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The full URI for connecting to the database cluster.
     */
    readonly uri?: pulumi.Input<string>;
    /**
     * The uniform resource name of the database cluster.
     */
    readonly urn?: pulumi.Input<string>;
    /**
     * Username for the cluster's default user.
     */
    readonly user?: pulumi.Input<string>;
    /**
     * Engine version used by the cluster (ex. `11` for PostgreSQL 11).
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseCluster resource.
 */
export interface DatabaseClusterArgs {
    /**
     * Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
     */
    readonly engine: pulumi.Input<string>;
    /**
     * A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
     */
    readonly evictionPolicy?: pulumi.Input<string>;
    /**
     * Defines when the automatic maintenance should be performed for the database cluster.
     */
    readonly maintenanceWindows?: pulumi.Input<pulumi.Input<inputs.DatabaseClusterMaintenanceWindow>[]>;
    /**
     * The name of the database cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Number of nodes that will be included in the cluster.
     */
    readonly nodeCount: pulumi.Input<number>;
    /**
     * The ID of the VPC where the database cluster will be located.
     */
    readonly privateNetworkUuid?: pulumi.Input<string>;
    /**
     * DigitalOcean region where the cluster will reside.
     */
    readonly region: pulumi.Input<Region>;
    /**
     * Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
     */
    readonly size: pulumi.Input<DatabaseSlug>;
    /**
     * A comma separated string specifying the  SQL modes for a MySQL cluster.
     */
    readonly sqlMode?: pulumi.Input<string>;
    /**
     * A list of tag names to be applied to the database cluster.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Engine version used by the cluster (ex. `11` for PostgreSQL 11).
     */
    readonly version?: pulumi.Input<string>;
}
