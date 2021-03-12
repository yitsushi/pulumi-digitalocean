// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import {DatabaseSlug, Region} from "./index";

/**
 * Provides a DigitalOcean database replica resource.
 *
 * ## Example Usage
 * ### Create a new PostgreSQL database replica
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const postgres_example = new digitalocean.DatabaseCluster("postgres-example", {
 *     engine: "pg",
 *     version: "11",
 *     size: "db-s-1vcpu-1gb",
 *     region: "nyc1",
 *     nodeCount: 1,
 * });
 * const read_replica = new digitalocean.DatabaseReplica("read-replica", {
 *     clusterId: postgres_example.id,
 *     size: "db-s-1vcpu-1gb",
 *     region: "nyc1",
 * });
 * ```
 *
 * ## Import
 *
 * Database replicas can be imported using the `id` of the source database cluster and the `name` of the replica joined with a comma. For example
 *
 * ```sh
 *  $ pulumi import digitalocean:index/databaseReplica:DatabaseReplica read-replica 245bcfd0-7f31-4ce6-a2bc-475a116cca97,read-replica
 * ```
 */
export class DatabaseReplica extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseReplica resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseReplicaState, opts?: pulumi.CustomResourceOptions): DatabaseReplica {
        return new DatabaseReplica(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/databaseReplica:DatabaseReplica';

    /**
     * Returns true if the given object is an instance of DatabaseReplica.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseReplica {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseReplica.__pulumiType;
    }

    /**
     * The ID of the original source database cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Name of the replica's default database.
     */
    public /*out*/ readonly database!: pulumi.Output<string>;
    /**
     * Database replica's hostname.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * The name for the database replica.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password for the replica's default user.
     */
    public /*out*/ readonly password!: pulumi.Output<string>;
    /**
     * Network port that the database replica is listening on.
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * Same as `host`, but only accessible from resources within the account and in the same region.
     */
    public /*out*/ readonly privateHost!: pulumi.Output<string>;
    /**
     * The ID of the VPC where the database replica will be located.
     */
    public readonly privateNetworkUuid!: pulumi.Output<string>;
    /**
     * Same as `uri`, but only accessible from resources within the account and in the same region.
     */
    public /*out*/ readonly privateUri!: pulumi.Output<string>;
    /**
     * DigitalOcean region where the replica will reside.
     */
    public readonly region!: pulumi.Output<Region | undefined>;
    /**
     * Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
     */
    public readonly size!: pulumi.Output<DatabaseSlug | undefined>;
    /**
     * A list of tag names to be applied to the database replica.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The full URI for connecting to the database replica.
     */
    public /*out*/ readonly uri!: pulumi.Output<string>;
    /**
     * Username for the replica's default user.
     */
    public /*out*/ readonly user!: pulumi.Output<string>;

    /**
     * Create a DatabaseReplica resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseReplicaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseReplicaArgs | DatabaseReplicaState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatabaseReplicaState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["database"] = state ? state.database : undefined;
            inputs["host"] = state ? state.host : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["password"] = state ? state.password : undefined;
            inputs["port"] = state ? state.port : undefined;
            inputs["privateHost"] = state ? state.privateHost : undefined;
            inputs["privateNetworkUuid"] = state ? state.privateNetworkUuid : undefined;
            inputs["privateUri"] = state ? state.privateUri : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["uri"] = state ? state.uri : undefined;
            inputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as DatabaseReplicaArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["privateNetworkUuid"] = args ? args.privateNetworkUuid : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["database"] = undefined /*out*/;
            inputs["host"] = undefined /*out*/;
            inputs["password"] = undefined /*out*/;
            inputs["port"] = undefined /*out*/;
            inputs["privateHost"] = undefined /*out*/;
            inputs["privateUri"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
            inputs["user"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DatabaseReplica.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseReplica resources.
 */
export interface DatabaseReplicaState {
    /**
     * The ID of the original source database cluster.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * Name of the replica's default database.
     */
    readonly database?: pulumi.Input<string>;
    /**
     * Database replica's hostname.
     */
    readonly host?: pulumi.Input<string>;
    /**
     * The name for the database replica.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Password for the replica's default user.
     */
    readonly password?: pulumi.Input<string>;
    /**
     * Network port that the database replica is listening on.
     */
    readonly port?: pulumi.Input<number>;
    /**
     * Same as `host`, but only accessible from resources within the account and in the same region.
     */
    readonly privateHost?: pulumi.Input<string>;
    /**
     * The ID of the VPC where the database replica will be located.
     */
    readonly privateNetworkUuid?: pulumi.Input<string>;
    /**
     * Same as `uri`, but only accessible from resources within the account and in the same region.
     */
    readonly privateUri?: pulumi.Input<string>;
    /**
     * DigitalOcean region where the replica will reside.
     */
    readonly region?: pulumi.Input<Region>;
    /**
     * Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
     */
    readonly size?: pulumi.Input<DatabaseSlug>;
    /**
     * A list of tag names to be applied to the database replica.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The full URI for connecting to the database replica.
     */
    readonly uri?: pulumi.Input<string>;
    /**
     * Username for the replica's default user.
     */
    readonly user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatabaseReplica resource.
 */
export interface DatabaseReplicaArgs {
    /**
     * The ID of the original source database cluster.
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * The name for the database replica.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the VPC where the database replica will be located.
     */
    readonly privateNetworkUuid?: pulumi.Input<string>;
    /**
     * DigitalOcean region where the replica will reside.
     */
    readonly region?: pulumi.Input<Region>;
    /**
     * Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
     */
    readonly size?: pulumi.Input<DatabaseSlug>;
    /**
     * A list of tag names to be applied to the database replica.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
