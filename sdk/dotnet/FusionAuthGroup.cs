// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth
{
    /// <summary>
    /// ## # Group Resource
    /// 
    /// A FusionAuth Group is a named object that optionally contains one to many Application Roles.
    /// 
    /// When a Group does not contain any Application Roles it can still be utilized to logically associate users. Assigning Application Roles to a group allow it to be used to dynamically manage Role assignment to registered Users. In this second scenario as long as a User is registered to an Application the Group membership will allow them to inherit the corresponding Roles from the Group.
    /// 
    /// [Groups API](https://fusionauth.io/docs/v1/tech/apis/groups)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Fusionauth = theogravity.Fusionauth;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myGroup = new Fusionauth.FusionAuthGroup("myGroup", new Fusionauth.FusionAuthGroupArgs
    ///         {
    ///             TenantId = fusionauth_tenant.My_tenant.Id,
    ///             RoleIds = 
    ///             {
    ///                 fusionauth_application_role.Admins.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthGroup:FusionAuthGroup")]
    public partial class FusionAuthGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// An object that can hold any information about the Group that should be persisted.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, object>?> Data { get; private set; } = null!;

        /// <summary>
        /// The Id to use for the new Group. If not specified a secure random UUID will be generated.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the Group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Application Roles to assign to this group.
        /// </summary>
        [Output("roleIds")]
        public Output<ImmutableArray<string>> RoleIds { get; private set; } = null!;

        /// <summary>
        /// The unique Id of the tenant used to scope this API request.
        /// </summary>
        [Output("tenantId")]
        public Output<string> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthGroup(string name, FusionAuthGroupArgs args, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthGroup:FusionAuthGroup", name, args ?? new FusionAuthGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthGroup(string name, Input<string> id, FusionAuthGroupState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthGroup:FusionAuthGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/theogravity/pulumi-fusionauth/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FusionAuthGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthGroup Get(string name, Input<string> id, FusionAuthGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthGroup(name, id, state, options);
        }
    }

    public sealed class FusionAuthGroupArgs : Pulumi.ResourceArgs
    {
        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// An object that can hold any information about the Group that should be persisted.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        /// <summary>
        /// The Id to use for the new Group. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("roleIds")]
        private InputList<string>? _roleIds;

        /// <summary>
        /// The Application Roles to assign to this group.
        /// </summary>
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        /// <summary>
        /// The unique Id of the tenant used to scope this API request.
        /// </summary>
        [Input("tenantId", required: true)]
        public Input<string> TenantId { get; set; } = null!;

        public FusionAuthGroupArgs()
        {
        }
    }

    public sealed class FusionAuthGroupState : Pulumi.ResourceArgs
    {
        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// An object that can hold any information about the Group that should be persisted.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        /// <summary>
        /// The Id to use for the new Group. If not specified a secure random UUID will be generated.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("roleIds")]
        private InputList<string>? _roleIds;

        /// <summary>
        /// The Application Roles to assign to this group.
        /// </summary>
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        /// <summary>
        /// The unique Id of the tenant used to scope this API request.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public FusionAuthGroupState()
        {
        }
    }
}
