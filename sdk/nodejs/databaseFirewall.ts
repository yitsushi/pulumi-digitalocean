// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a DigitalOcean database firewall resource allowing you to restrict
 * connections to your database to trusted sources. You may limit connections to
 * specific Droplets, Kubernetes clusters, or IP addresses.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/database_firewall.html.markdown.
 */
export class DatabaseFirewall extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseFirewall resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseFirewallState, opts?: pulumi.CustomResourceOptions): DatabaseFirewall {
        return new DatabaseFirewall(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index:DatabaseFirewall';

    /**
     * Returns true if the given object is an instance of DatabaseFirewall.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatabaseFirewall {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatabaseFirewall.__pulumiType;
    }

    /**
     * The ID of the target database cluster.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * A rule specifying a resource allowed to access the database cluster. The following arguments must be specified:
     * - `type` - (Required) The type of resource that the firewall rule allows to access the database cluster. The possible values are: `droplet`, `k8s`, `ipAddr`, or `tag`.
     * - `value` - (Required) The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.
     */
    public readonly rules!: pulumi.Output<outputs.DatabaseFirewallRule[]>;

    /**
     * Create a DatabaseFirewall resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseFirewallArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseFirewallArgs | DatabaseFirewallState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseFirewallState | undefined;
            inputs["clusterId"] = state ? state.clusterId : undefined;
            inputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as DatabaseFirewallArgs | undefined;
            if (!args || args.clusterId === undefined) {
                throw new Error("Missing required property 'clusterId'");
            }
            if (!args || args.rules === undefined) {
                throw new Error("Missing required property 'rules'");
            }
            inputs["clusterId"] = args ? args.clusterId : undefined;
            inputs["rules"] = args ? args.rules : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DatabaseFirewall.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseFirewall resources.
 */
export interface DatabaseFirewallState {
    /**
     * The ID of the target database cluster.
     */
    readonly clusterId?: pulumi.Input<string>;
    /**
     * A rule specifying a resource allowed to access the database cluster. The following arguments must be specified:
     * - `type` - (Required) The type of resource that the firewall rule allows to access the database cluster. The possible values are: `droplet`, `k8s`, `ipAddr`, or `tag`.
     * - `value` - (Required) The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.DatabaseFirewallRule>[]>;
}

/**
 * The set of arguments for constructing a DatabaseFirewall resource.
 */
export interface DatabaseFirewallArgs {
    /**
     * The ID of the target database cluster.
     */
    readonly clusterId: pulumi.Input<string>;
    /**
     * A rule specifying a resource allowed to access the database cluster. The following arguments must be specified:
     * - `type` - (Required) The type of resource that the firewall rule allows to access the database cluster. The possible values are: `droplet`, `k8s`, `ipAddr`, or `tag`.
     * - `value` - (Required) The ID of the specific resource, the name of a tag applied to a group of resources, or the IP address that the firewall rule allows to access the database cluster.
     */
    readonly rules: pulumi.Input<pulumi.Input<inputs.DatabaseFirewallRule>[]>;
}
