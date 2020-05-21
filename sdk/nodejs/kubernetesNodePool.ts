// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

import {DropletSlug} from "./index";

/**
 * Provides a DigitalOcean Kubernetes node pool resource. While the default node pool must be defined in the `digitalocean..KubernetesCluster` resource, this resource can be used to add additional ones to a cluster.
 *
 * ## Example Usage
 *
 * ### Basic Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const foo = new digitalocean.KubernetesCluster("foo", {
 *     region: "nyc1",
 *     version: "1.15.5-do.1",
 *     node_pool: {
 *         name: "front-end-pool",
 *         size: "s-2vcpu-2gb",
 *         nodeCount: 3,
 *     },
 * });
 * const bar = new digitalocean.KubernetesNodePool("bar", {
 *     clusterId: foo.id,
 *     size: "c-2",
 *     nodeCount: 2,
 *     tags: ["backend"],
 *     labels: {
 *         service: "backend",
 *         priority: "high",
 *     },
 * });
 * ```
 *
 * ### Autoscaling Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const autoscale-pool-01 = new digitalocean.KubernetesNodePool("autoscale-pool-01", {
 *     clusterId: digitalocean_kubernetes_cluster.foo.id,
 *     size: "s-1vcpu-2gb",
 *     autoScale: true,
 *     minNodes: 0,
 *     maxNodes: 5,
 * });
 * ```
 */
export class KubernetesNodePool extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesNodePool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesNodePoolState, opts?: pulumi.CustomResourceOptions): KubernetesNodePool {
        return new KubernetesNodePool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/kubernetesNodePool:KubernetesNodePool';

    /**
     * Returns true if the given object is an instance of KubernetesNodePool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KubernetesNodePool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KubernetesNodePool.__pulumiType;
    }

    /**
     * A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
     */
    public /*out*/ readonly actualNodeCount!: pulumi.Output<number>;
    /**
     * Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
     */
    public readonly autoScale!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Kubernetes cluster to which the node pool is associated.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
     */
    public readonly maxNodes!: pulumi.Output<number | undefined>;
    /**
     * If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
     */
    public readonly minNodes!: pulumi.Output<number | undefined>;
    /**
     * A name for the node pool.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
     */
    public readonly nodeCount!: pulumi.Output<number | undefined>;
    /**
     * A list of nodes in the pool. Each node exports the following attributes:
     * - `id` -  A unique ID that can be used to identify and reference the node.
     * - `name` - The auto-generated name for the node.
     * - `status` -  A string indicating the current status of the individual node.
     * - `dropletId` - The id of the node's droplet
     * - `createdAt` - The date and time when the node was created.
     * - `updatedAt` - The date and time when the node was last updated.
     */
    public /*out*/ readonly nodes!: pulumi.Output<outputs.KubernetesNodePoolNode[]>;
    /**
     * The slug identifier for the type of Droplet to be used as workers in the node pool.
     */
    public readonly size!: pulumi.Output<DropletSlug>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a KubernetesNodePool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesNodePoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesNodePoolArgs | KubernetesNodePoolState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as KubernetesNodePoolState | undefined;
            inputs["actualNodeCount"] = state ? state.actualNodeCount : undefined;
            inputs["autoScale"] = state ? state.autoScale : undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["maxNodes"] = state ? state.maxNodes : undefined;
            inputs["minNodes"] = state ? state.minNodes : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodeCount"] = state ? state.nodeCount : undefined;
            inputs["nodes"] = state ? state.nodes : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as KubernetesNodePoolArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.size === undefined) {
                throw new Error("Missing required property 'size'");
            }
            inputs["autoScale"] = args ? args.autoScale : undefined;
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["maxNodes"] = args ? args.maxNodes : undefined;
            inputs["minNodes"] = args ? args.minNodes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeCount"] = args ? args.nodeCount : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["actualNodeCount"] = undefined /*out*/;
            inputs["nodes"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(KubernetesNodePool.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KubernetesNodePool resources.
 */
export interface KubernetesNodePoolState {
    /**
     * A computed field representing the actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
     */
    readonly actualNodeCount?: pulumi.Input<number>;
    /**
     * Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
     */
    readonly autoScale?: pulumi.Input<boolean>;
    /**
     * The ID of the Kubernetes cluster to which the node pool is associated.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
     */
    readonly maxNodes?: pulumi.Input<number>;
    /**
     * If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
     */
    readonly minNodes?: pulumi.Input<number>;
    /**
     * A name for the node pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * A list of nodes in the pool. Each node exports the following attributes:
     * - `id` -  A unique ID that can be used to identify and reference the node.
     * - `name` - The auto-generated name for the node.
     * - `status` -  A string indicating the current status of the individual node.
     * - `dropletId` - The id of the node's droplet
     * - `createdAt` - The date and time when the node was created.
     * - `updatedAt` - The date and time when the node was last updated.
     */
    readonly nodes?: pulumi.Input<pulumi.Input<inputs.KubernetesNodePoolNode>[]>;
    /**
     * The slug identifier for the type of Droplet to be used as workers in the node pool.
     */
    readonly size?: pulumi.Input<DropletSlug>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a KubernetesNodePool resource.
 */
export interface KubernetesNodePoolArgs {
    /**
     * Enable auto-scaling of the number of nodes in the node pool within the given min/max range.
     */
    readonly autoScale?: pulumi.Input<boolean>;
    /**
     * The ID of the Kubernetes cluster to which the node pool is associated.
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * A map of key/value pairs to apply to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
     */
    readonly maxNodes?: pulumi.Input<number>;
    /**
     * If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
     */
    readonly minNodes?: pulumi.Input<number>;
    /**
     * A name for the node pool.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The number of Droplet instances in the node pool. If auto-scaling is enabled, this should only be set if the desired result is to explicitly reset the number of nodes to this value. If auto-scaling is enabled, and the node count is outside of the given min/max range, it will use the min nodes value.
     */
    readonly nodeCount?: pulumi.Input<number>;
    /**
     * The slug identifier for the type of Droplet to be used as workers in the node pool.
     */
    readonly size: pulumi.Input<DropletSlug>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
}
