package discord

import (
	"golang.org/x/oauth2"
)

// All scope constants that can be used.
const (
	ScopeActivitiesRead                       = "activities.read"  // Whitelist only
	ScopeActivitiesWrite                      = "activities.write" // Whitelist only
	ScopeApplicationsBuildsRead               = "applications.builds.read"
	ScopeApplicationsBuildsUpload             = "applications.builds.upload" // Whitelist only
	ScopeApplicationsCommands                 = "applications.commands"
	ScopeApplicationsCommandsPermissionUpdate = "applications.commands.permissions.update" // Whitelist only
	ScopeApplicationsCommandsUpdate           = "applications.commands.update"
	ScopeApplicationsEntitlements             = "applications.entitlements"
	ScopeApplicationsStoreUpdate              = "applications.store.update"
	ScopeBot                                  = "bot"
	ScopeConnections                          = "connections"
	ScopeDMChannelsRead                       = "dm_channels.read" // Whitelist only
	ScopeEmail                                = "email"
	ScopeGroupDMJoin                          = "gdm.join"
	ScopeGuilds                               = "guilds"
	ScopeGuildsJoin                           = "guilds.join"
	ScopeGuildsMembersRead                    = "guilds.members.read"
	ScopeIdentify                             = "identify"
	ScopeMessagesRead                         = "messages.read"
	ScopeRelationshipsRead                    = "relationships.read"     // Whitelist only
	ScopeRPC                                  = "rpc"                    // Whitelist only
	ScopeRPCActivitiesWrite                   = "rpc.activities.write"   // Whitelist only
	ScopeRPCNotificationsRead                 = "rpc.notifications.read" // Whitelist only
	ScopeRPCRead                              = "rpc.voice.read"         // Whitelist only
	ScopeRPCWrite                             = "rpc.voice.write"        // Whitelist only
	ScopeVoice                                = "voice"                  // Whitelist only
	ScopeWebhookIncoming                      = "webhook.Incoming"
)

// Endpoint is Discord's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:   "https://discord.com/api/oauth2/authorize",
	TokenURL:  "https://discord.com/api/oauth2/token",
	AuthStyle: oauth2.AuthStyleInParams,
}
