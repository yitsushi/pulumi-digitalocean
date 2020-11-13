// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource which can be used to create a custom Image from a URL
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as digitalocean from "@pulumi/digitalocean";
 *
 * const flatcar = new digitalocean.CustomImage("flatcar", {
 *     regions: ["nyc3"],
 *     url: "https://stable.release.flatcar-linux.net/amd64-usr/2605.7.0/flatcar_production_digitalocean_image.bin.bz2",
 * });
 * ```
 */
export class CustomImage extends pulumi.CustomResource {
    /**
     * Get an existing CustomImage resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomImageState, opts?: pulumi.CustomResourceOptions): CustomImage {
        return new CustomImage(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'digitalocean:index/customImage:CustomImage';

    /**
     * Returns true if the given object is an instance of CustomImage.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomImage {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomImage.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly distribution!: pulumi.Output<string | undefined>;
    public /*out*/ readonly imageId!: pulumi.Output<number>;
    public /*out*/ readonly minDiskSize!: pulumi.Output<number>;
    /**
     * A name for the Custom Image.
     */
    public readonly name!: pulumi.Output<string>;
    public /*out*/ readonly public!: pulumi.Output<boolean>;
    /**
     * A list of regions. (Currently only one is supported)
     */
    public readonly regions!: pulumi.Output<string[]>;
    public /*out*/ readonly sizeGigabytes!: pulumi.Output<number>;
    public /*out*/ readonly slug!: pulumi.Output<string>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A URL from which the custom Linux virtual machine image may be retrieved.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a CustomImage resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomImageArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomImageArgs | CustomImageState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as CustomImageState | undefined;
            inputs["createdAt"] = state ? state.createdAt : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["distribution"] = state ? state.distribution : undefined;
            inputs["imageId"] = state ? state.imageId : undefined;
            inputs["minDiskSize"] = state ? state.minDiskSize : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["public"] = state ? state.public : undefined;
            inputs["regions"] = state ? state.regions : undefined;
            inputs["sizeGigabytes"] = state ? state.sizeGigabytes : undefined;
            inputs["slug"] = state ? state.slug : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as CustomImageArgs | undefined;
            if (!args || args.regions === undefined) {
                throw new Error("Missing required property 'regions'");
            }
            if (!args || args.url === undefined) {
                throw new Error("Missing required property 'url'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["distribution"] = args ? args.distribution : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["regions"] = args ? args.regions : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["createdAt"] = undefined /*out*/;
            inputs["imageId"] = undefined /*out*/;
            inputs["minDiskSize"] = undefined /*out*/;
            inputs["public"] = undefined /*out*/;
            inputs["sizeGigabytes"] = undefined /*out*/;
            inputs["slug"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CustomImage.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomImage resources.
 */
export interface CustomImageState {
    readonly createdAt?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly distribution?: pulumi.Input<string>;
    readonly imageId?: pulumi.Input<number>;
    readonly minDiskSize?: pulumi.Input<number>;
    /**
     * A name for the Custom Image.
     */
    readonly name?: pulumi.Input<string>;
    readonly public?: pulumi.Input<boolean>;
    /**
     * A list of regions. (Currently only one is supported)
     */
    readonly regions?: pulumi.Input<pulumi.Input<string>[]>;
    readonly sizeGigabytes?: pulumi.Input<number>;
    readonly slug?: pulumi.Input<string>;
    readonly status?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    readonly type?: pulumi.Input<string>;
    /**
     * A URL from which the custom Linux virtual machine image may be retrieved.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomImage resource.
 */
export interface CustomImageArgs {
    readonly description?: pulumi.Input<string>;
    readonly distribution?: pulumi.Input<string>;
    /**
     * A name for the Custom Image.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of regions. (Currently only one is supported)
     */
    readonly regions: pulumi.Input<pulumi.Input<string>[]>;
    readonly tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A URL from which the custom Linux virtual machine image may be retrieved.
     */
    readonly url: pulumi.Input<string>;
}
