// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * > **NOTE:** DigitalOcean Kubernetes is currently in [Limited Availability](https://www.digitalocean.com/docs/platform/product-lifecycle/). In order to access its API, you must first enable Kubernetes on your account by opting-in via the [cloud control panel](https://cloud.digitalocean.com/kubernetes/clusters). While the Kubernetes Cluster functionality is currently in limited availability the structure of this resource may change over time. Please share any feedback you may have by [opening an issue on GitHub](https://github.com/terraform-providers/terraform-provider-digitalocean/issues).
 * 
 * Provides a DigitalOcean Kubernetes cluster resource. This can be used to create, delete, and modify clusters. For more information see the [official documentation](https://www.digitalocean.com/docs/kubernetes/).
 */
export class KubernetesCluster extends pulumi.CustomResource {
    /**
     * Get an existing KubernetesCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KubernetesClusterState, opts?: pulumi.CustomResourceOptions): KubernetesCluster {
        return new KubernetesCluster(name, <any>state, { ...opts, id: id });
    }

    /**
     * The range of IP addresses in the overlay network of the Kubernetes cluster.
     */
    public /*out*/ readonly clusterSubnet!: pulumi.Output<string>;
    /**
     * The date and time when the Kubernetes cluster was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The base URL of the API server on the Kubernetes master node.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The public IPv4 address of the Kubernetes master node.
     */
    public /*out*/ readonly ipv4Address!: pulumi.Output<string>;
    public /*out*/ readonly kubeConfigs!: pulumi.Output<{ clientCertificate: string, clientKey: string, clusterCaCertificate: string, host: string, rawConfig: string }[]>;
    /**
     * A name for the Kubernetes cluster.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean_kubernetes_node_pool` resource. The following arguments may be specified:
     * - `name` - (Required) A name for the node pool.
     * - `size` - (Required) The slug identifier for the type of Droplet to be used as workers in the node pool.
     * - `node_count` - (Required) The number of Droplet instances in the node pool.
     * - `tags` - (Optional) A list of tag names to be applied to the Kubernetes cluster.
     */
    public readonly nodePool!: pulumi.Output<{ id: string, name: string, nodeCount: number, nodes: { createdAt: string, id: string, name: string, status: string, updatedAt: string }[], size: string, tags?: string[] }>;
    /**
     * The slug identifier for the region where the Kubernetes cluster will be created.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The range of assignable IP addresses for services running in the Kubernetes cluster.
     */
    public /*out*/ readonly serviceSubnet!: pulumi.Output<string>;
    /**
     * A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * The date and time when the Kubernetes cluster was last updated.
     * * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
     * - `raw_config` - The full contents of the Kubernetes cluster's kubeconfig file.
     * - `host` - The URL of the API server on the Kubernetes master node.
     * - `client_key` - The base64 encoded private key used by clients to access the cluster.
     * - `client_certificate` - The base64 encoded public certificate used by clients to access the cluster.
     * - `cluster_ca_certificate` - The base64 encoded public certificate for the cluster's certificate authority.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The slug identifier for the version of Kubernetes used for the cluster.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a KubernetesCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KubernetesClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KubernetesClusterArgs | KubernetesClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as KubernetesClusterState | undefined;
            inputs["clusterSubnet"] = state ? state.clusterSubnet : undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["ipv4Address"] = state ? state.ipv4Address : undefined;
            inputs["kubeConfigs"] = state ? state.kubeConfigs : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["nodePool"] = state ? state.nodePool : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["serviceSubnet"] = state ? state.serviceSubnet : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["updatedAt"] = state ? state.updatedAt : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as KubernetesClusterArgs | undefined;
            if (!args || args.nodePool === undefined) {
                throw new Error("Missing required property 'nodePool'");
            }
            if (!args || args.region === undefined) {
                throw new Error("Missing required property 'region'");
            }
            if (!args || args.version === undefined) {
                throw new Error("Missing required property 'version'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["nodePool"] = args ? args.nodePool : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["clusterSubnet"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
            inputs["ipv4Address"] = undefined /*out*/;
            inputs["kubeConfigs"] = undefined /*out*/;
            inputs["serviceSubnet"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("digitalocean:index/kubernetesCluster:KubernetesCluster", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering KubernetesCluster resources.
 */
export interface KubernetesClusterState {
    /**
     * The range of IP addresses in the overlay network of the Kubernetes cluster.
     */
    readonly clusterSubnet?: pulumi.Input<string>;
    /**
     * The date and time when the Kubernetes cluster was created.
     */
    readonly createdAt?: pulumi.Input<string>;
    /**
     * The base URL of the API server on the Kubernetes master node.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * The public IPv4 address of the Kubernetes master node.
     */
    readonly ipv4Address?: pulumi.Input<string>;
    readonly kubeConfigs?: pulumi.Input<pulumi.Input<{ clientCertificate?: pulumi.Input<string>, clientKey?: pulumi.Input<string>, clusterCaCertificate?: pulumi.Input<string>, host?: pulumi.Input<string>, rawConfig?: pulumi.Input<string> }>[]>;
    /**
     * A name for the Kubernetes cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean_kubernetes_node_pool` resource. The following arguments may be specified:
     * - `name` - (Required) A name for the node pool.
     * - `size` - (Required) The slug identifier for the type of Droplet to be used as workers in the node pool.
     * - `node_count` - (Required) The number of Droplet instances in the node pool.
     * - `tags` - (Optional) A list of tag names to be applied to the Kubernetes cluster.
     */
    readonly nodePool?: pulumi.Input<{ id?: pulumi.Input<string>, name: pulumi.Input<string>, nodeCount: pulumi.Input<number>, nodes?: pulumi.Input<pulumi.Input<{ createdAt?: pulumi.Input<string>, id?: pulumi.Input<string>, name?: pulumi.Input<string>, status?: pulumi.Input<string>, updatedAt?: pulumi.Input<string> }>[]>, size: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * The slug identifier for the region where the Kubernetes cluster will be created.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The range of assignable IP addresses for services running in the Kubernetes cluster.
     */
    readonly serviceSubnet?: pulumi.Input<string>;
    /**
     * A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date and time when the Kubernetes cluster was last updated.
     * * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
     * - `raw_config` - The full contents of the Kubernetes cluster's kubeconfig file.
     * - `host` - The URL of the API server on the Kubernetes master node.
     * - `client_key` - The base64 encoded private key used by clients to access the cluster.
     * - `client_certificate` - The base64 encoded public certificate used by clients to access the cluster.
     * - `cluster_ca_certificate` - The base64 encoded public certificate for the cluster's certificate authority.
     */
    readonly updatedAt?: pulumi.Input<string>;
    /**
     * The slug identifier for the version of Kubernetes used for the cluster.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a KubernetesCluster resource.
 */
export interface KubernetesClusterArgs {
    /**
     * A name for the Kubernetes cluster.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A block representing the cluster's default node pool. Additional node pools may be added to the cluster using the `digitalocean_kubernetes_node_pool` resource. The following arguments may be specified:
     * - `name` - (Required) A name for the node pool.
     * - `size` - (Required) The slug identifier for the type of Droplet to be used as workers in the node pool.
     * - `node_count` - (Required) The number of Droplet instances in the node pool.
     * - `tags` - (Optional) A list of tag names to be applied to the Kubernetes cluster.
     */
    readonly nodePool: pulumi.Input<{ id?: pulumi.Input<string>, name: pulumi.Input<string>, nodeCount: pulumi.Input<number>, nodes?: pulumi.Input<pulumi.Input<{ createdAt?: pulumi.Input<string>, id?: pulumi.Input<string>, name?: pulumi.Input<string>, status?: pulumi.Input<string>, updatedAt?: pulumi.Input<string> }>[]>, size: pulumi.Input<string>, tags?: pulumi.Input<pulumi.Input<string>[]> }>;
    /**
     * The slug identifier for the region where the Kubernetes cluster will be created.
     */
    readonly region: pulumi.Input<string>;
    /**
     * A list of tag names to be applied to the Kubernetes cluster.
     */
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The slug identifier for the version of Kubernetes used for the cluster.
     */
    readonly version: pulumi.Input<string>;
}
