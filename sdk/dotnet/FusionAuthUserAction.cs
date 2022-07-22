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
    /// ## # User Action Resource
    /// 
    /// [User Actions API](https://fusionauth.io/docs/v1/tech/apis/user-actions/)
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
    ///         var example = new Fusionauth.FusionAuthUserAction("example", new Fusionauth.FusionAuthUserActionArgs
    ///         {
    ///             PreventLogin = true,
    ///             Temporal = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthUserAction:FusionAuthUserAction")]
    public partial class FusionAuthUserAction : Pulumi.CustomResource
    {
        /// <summary>
        /// The Id of the Email Template that is used when User Actions are canceled.
        /// </summary>
        [Output("cancelEmailTemplateId")]
        public Output<string?> CancelEmailTemplateId { get; private set; } = null!;

        /// <summary>
        /// The Id of the Email Template that is used when User Actions expired automatically (end).
        /// </summary>
        [Output("endEmailTemplateId")]
        public Output<string?> EndEmailTemplateId { get; private set; } = null!;

        /// <summary>
        /// Whether to include the email information in the JSON that is sent to the Webhook when a user action is taken.
        /// </summary>
        [Output("includeEmailInEventJson")]
        public Output<bool?> IncludeEmailInEventJson { get; private set; } = null!;

        /// <summary>
        /// A mapping of localized names for this User Action Option. The key is the Locale and the value is the name of the User Action Option for that language.
        /// </summary>
        [Output("localizedNames")]
        public Output<ImmutableDictionary<string, object>?> LocalizedNames { get; private set; } = null!;

        /// <summary>
        /// The Id of the Email Template that is used when User Actions are modified.
        /// </summary>
        [Output("modifyEmailTemplateId")]
        public Output<string?> ModifyEmailTemplateId { get; private set; } = null!;

        /// <summary>
        /// The name of this User Action Option.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of User Action Options.
        /// </summary>
        [Output("options")]
        public Output<ImmutableArray<Outputs.FusionAuthUserActionOption>> Options { get; private set; } = null!;

        /// <summary>
        /// Whether or not this User Action will prevent user login. When this value is set to true the action must also be marked as a time based action. See `temporal`.
        /// </summary>
        [Output("preventLogin")]
        public Output<bool?> PreventLogin { get; private set; } = null!;

        /// <summary>
        /// Whether or not FusionAuth will send events to any registered Webhooks when this User Action expires.
        /// </summary>
        [Output("sendEndEvent")]
        public Output<bool?> SendEndEvent { get; private set; } = null!;

        /// <summary>
        /// The Id of the Email Template that is used when User Actions are started (created).
        /// </summary>
        [Output("startEmailTemplateId")]
        public Output<string?> StartEmailTemplateId { get; private set; } = null!;

        /// <summary>
        /// Whether or not this User Action is time-based (temporal).
        /// </summary>
        [Output("temporal")]
        public Output<bool?> Temporal { get; private set; } = null!;

        /// <summary>
        /// Whether or not email is enabled for this User Action.
        /// </summary>
        [Output("userEmailingEnabled")]
        public Output<bool?> UserEmailingEnabled { get; private set; } = null!;

        /// <summary>
        /// Whether or not user notifications are enabled for this User Action.
        /// </summary>
        [Output("userNotificationsEnabled")]
        public Output<bool?> UserNotificationsEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthUserAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthUserAction(string name, FusionAuthUserActionArgs? args = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthUserAction:FusionAuthUserAction", name, args ?? new FusionAuthUserActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthUserAction(string name, Input<string> id, FusionAuthUserActionState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthUserAction:FusionAuthUserAction", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FusionAuthUserAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthUserAction Get(string name, Input<string> id, FusionAuthUserActionState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthUserAction(name, id, state, options);
        }
    }

    public sealed class FusionAuthUserActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Id of the Email Template that is used when User Actions are canceled.
        /// </summary>
        [Input("cancelEmailTemplateId")]
        public Input<string>? CancelEmailTemplateId { get; set; }

        /// <summary>
        /// The Id of the Email Template that is used when User Actions expired automatically (end).
        /// </summary>
        [Input("endEmailTemplateId")]
        public Input<string>? EndEmailTemplateId { get; set; }

        /// <summary>
        /// Whether to include the email information in the JSON that is sent to the Webhook when a user action is taken.
        /// </summary>
        [Input("includeEmailInEventJson")]
        public Input<bool>? IncludeEmailInEventJson { get; set; }

        [Input("localizedNames")]
        private InputMap<object>? _localizedNames;

        /// <summary>
        /// A mapping of localized names for this User Action Option. The key is the Locale and the value is the name of the User Action Option for that language.
        /// </summary>
        public InputMap<object> LocalizedNames
        {
            get => _localizedNames ?? (_localizedNames = new InputMap<object>());
            set => _localizedNames = value;
        }

        /// <summary>
        /// The Id of the Email Template that is used when User Actions are modified.
        /// </summary>
        [Input("modifyEmailTemplateId")]
        public Input<string>? ModifyEmailTemplateId { get; set; }

        /// <summary>
        /// The name of this User Action Option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        private InputList<Inputs.FusionAuthUserActionOptionArgs>? _options;

        /// <summary>
        /// The list of User Action Options.
        /// </summary>
        public InputList<Inputs.FusionAuthUserActionOptionArgs> Options
        {
            get => _options ?? (_options = new InputList<Inputs.FusionAuthUserActionOptionArgs>());
            set => _options = value;
        }

        /// <summary>
        /// Whether or not this User Action will prevent user login. When this value is set to true the action must also be marked as a time based action. See `temporal`.
        /// </summary>
        [Input("preventLogin")]
        public Input<bool>? PreventLogin { get; set; }

        /// <summary>
        /// Whether or not FusionAuth will send events to any registered Webhooks when this User Action expires.
        /// </summary>
        [Input("sendEndEvent")]
        public Input<bool>? SendEndEvent { get; set; }

        /// <summary>
        /// The Id of the Email Template that is used when User Actions are started (created).
        /// </summary>
        [Input("startEmailTemplateId")]
        public Input<string>? StartEmailTemplateId { get; set; }

        /// <summary>
        /// Whether or not this User Action is time-based (temporal).
        /// </summary>
        [Input("temporal")]
        public Input<bool>? Temporal { get; set; }

        /// <summary>
        /// Whether or not email is enabled for this User Action.
        /// </summary>
        [Input("userEmailingEnabled")]
        public Input<bool>? UserEmailingEnabled { get; set; }

        /// <summary>
        /// Whether or not user notifications are enabled for this User Action.
        /// </summary>
        [Input("userNotificationsEnabled")]
        public Input<bool>? UserNotificationsEnabled { get; set; }

        public FusionAuthUserActionArgs()
        {
        }
    }

    public sealed class FusionAuthUserActionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Id of the Email Template that is used when User Actions are canceled.
        /// </summary>
        [Input("cancelEmailTemplateId")]
        public Input<string>? CancelEmailTemplateId { get; set; }

        /// <summary>
        /// The Id of the Email Template that is used when User Actions expired automatically (end).
        /// </summary>
        [Input("endEmailTemplateId")]
        public Input<string>? EndEmailTemplateId { get; set; }

        /// <summary>
        /// Whether to include the email information in the JSON that is sent to the Webhook when a user action is taken.
        /// </summary>
        [Input("includeEmailInEventJson")]
        public Input<bool>? IncludeEmailInEventJson { get; set; }

        [Input("localizedNames")]
        private InputMap<object>? _localizedNames;

        /// <summary>
        /// A mapping of localized names for this User Action Option. The key is the Locale and the value is the name of the User Action Option for that language.
        /// </summary>
        public InputMap<object> LocalizedNames
        {
            get => _localizedNames ?? (_localizedNames = new InputMap<object>());
            set => _localizedNames = value;
        }

        /// <summary>
        /// The Id of the Email Template that is used when User Actions are modified.
        /// </summary>
        [Input("modifyEmailTemplateId")]
        public Input<string>? ModifyEmailTemplateId { get; set; }

        /// <summary>
        /// The name of this User Action Option.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("options")]
        private InputList<Inputs.FusionAuthUserActionOptionGetArgs>? _options;

        /// <summary>
        /// The list of User Action Options.
        /// </summary>
        public InputList<Inputs.FusionAuthUserActionOptionGetArgs> Options
        {
            get => _options ?? (_options = new InputList<Inputs.FusionAuthUserActionOptionGetArgs>());
            set => _options = value;
        }

        /// <summary>
        /// Whether or not this User Action will prevent user login. When this value is set to true the action must also be marked as a time based action. See `temporal`.
        /// </summary>
        [Input("preventLogin")]
        public Input<bool>? PreventLogin { get; set; }

        /// <summary>
        /// Whether or not FusionAuth will send events to any registered Webhooks when this User Action expires.
        /// </summary>
        [Input("sendEndEvent")]
        public Input<bool>? SendEndEvent { get; set; }

        /// <summary>
        /// The Id of the Email Template that is used when User Actions are started (created).
        /// </summary>
        [Input("startEmailTemplateId")]
        public Input<string>? StartEmailTemplateId { get; set; }

        /// <summary>
        /// Whether or not this User Action is time-based (temporal).
        /// </summary>
        [Input("temporal")]
        public Input<bool>? Temporal { get; set; }

        /// <summary>
        /// Whether or not email is enabled for this User Action.
        /// </summary>
        [Input("userEmailingEnabled")]
        public Input<bool>? UserEmailingEnabled { get; set; }

        /// <summary>
        /// Whether or not user notifications are enabled for this User Action.
        /// </summary>
        [Input("userNotificationsEnabled")]
        public Input<bool>? UserNotificationsEnabled { get; set; }

        public FusionAuthUserActionState()
        {
        }
    }
}
