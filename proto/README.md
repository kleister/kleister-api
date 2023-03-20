# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [types/v1/build.proto](#types_v1_build-proto)
    - [Build](#types-v1-Build)
  
- [types/v1/version.proto](#types_v1_version-proto)
    - [Version](#types-v1-Version)
    - [VersionFile](#types-v1-VersionFile)
  
- [builds/v1/builds.proto](#builds_v1_builds-proto)
    - [AttachVersionRequest](#builds-v1-AttachVersionRequest)
    - [AttachVersionResponse](#builds-v1-AttachVersionResponse)
    - [CreateBuild](#builds-v1-CreateBuild)
    - [CreateRequest](#builds-v1-CreateRequest)
    - [CreateResponse](#builds-v1-CreateResponse)
    - [DeleteRequest](#builds-v1-DeleteRequest)
    - [DeleteResponse](#builds-v1-DeleteResponse)
    - [DropVersionRequest](#builds-v1-DropVersionRequest)
    - [DropVersionResponse](#builds-v1-DropVersionResponse)
    - [ListRequest](#builds-v1-ListRequest)
    - [ListResponse](#builds-v1-ListResponse)
    - [ListVersionsRequest](#builds-v1-ListVersionsRequest)
    - [ListVersionsResponse](#builds-v1-ListVersionsResponse)
    - [ShowRequest](#builds-v1-ShowRequest)
    - [ShowResponse](#builds-v1-ShowResponse)
    - [UpdateBuild](#builds-v1-UpdateBuild)
    - [UpdateRequest](#builds-v1-UpdateRequest)
    - [UpdateResponse](#builds-v1-UpdateResponse)
  
    - [BuildsService](#builds-v1-BuildsService)
  
- [types/v1/forge.proto](#types_v1_forge-proto)
    - [Forge](#types-v1-Forge)
  
- [forge/v1/forge.proto](#forge_v1_forge-proto)
    - [ListBuildsRequest](#forge-v1-ListBuildsRequest)
    - [ListBuildsResponse](#forge-v1-ListBuildsResponse)
    - [SearchRequest](#forge-v1-SearchRequest)
    - [SearchResponse](#forge-v1-SearchResponse)
    - [UpdateRequest](#forge-v1-UpdateRequest)
    - [UpdateResponse](#forge-v1-UpdateResponse)
  
    - [ForgeService](#forge-v1-ForgeService)
  
- [types/v1/member.proto](#types_v1_member-proto)
    - [Member](#types-v1-Member)
  
- [members/v1/members.proto](#members_v1_members-proto)
    - [AppendMember](#members-v1-AppendMember)
    - [AppendRequest](#members-v1-AppendRequest)
    - [AppendResponse](#members-v1-AppendResponse)
    - [DropMember](#members-v1-DropMember)
    - [DropRequest](#members-v1-DropRequest)
    - [DropResponse](#members-v1-DropResponse)
    - [ListRequest](#members-v1-ListRequest)
    - [ListResponse](#members-v1-ListResponse)
  
    - [MembersService](#members-v1-MembersService)
  
- [types/v1/minecraft.proto](#types_v1_minecraft-proto)
    - [Minecraft](#types-v1-Minecraft)
  
- [minecraft/v1/minecraft.proto](#minecraft_v1_minecraft-proto)
    - [ListBuildsRequest](#minecraft-v1-ListBuildsRequest)
    - [ListBuildsResponse](#minecraft-v1-ListBuildsResponse)
    - [SearchRequest](#minecraft-v1-SearchRequest)
    - [SearchResponse](#minecraft-v1-SearchResponse)
    - [UpdateRequest](#minecraft-v1-UpdateRequest)
    - [UpdateResponse](#minecraft-v1-UpdateResponse)
  
    - [MinecraftService](#minecraft-v1-MinecraftService)
  
- [types/v1/mod.proto](#types_v1_mod-proto)
    - [Mod](#types-v1-Mod)
  
- [types/v1/team.proto](#types_v1_team-proto)
    - [Team](#types-v1-Team)
  
- [types/v1/user.proto](#types_v1_user-proto)
    - [User](#types-v1-User)
  
- [mods/v1/mods.proto](#mods_v1_mods-proto)
    - [AttachTeamRequest](#mods-v1-AttachTeamRequest)
    - [AttachTeamResponse](#mods-v1-AttachTeamResponse)
    - [AttachUserRequest](#mods-v1-AttachUserRequest)
    - [AttachUserResponse](#mods-v1-AttachUserResponse)
    - [CreateMod](#mods-v1-CreateMod)
    - [CreateRequest](#mods-v1-CreateRequest)
    - [CreateResponse](#mods-v1-CreateResponse)
    - [DeleteRequest](#mods-v1-DeleteRequest)
    - [DeleteResponse](#mods-v1-DeleteResponse)
    - [DropTeamRequest](#mods-v1-DropTeamRequest)
    - [DropTeamResponse](#mods-v1-DropTeamResponse)
    - [DropUserRequest](#mods-v1-DropUserRequest)
    - [DropUserResponse](#mods-v1-DropUserResponse)
    - [ListRequest](#mods-v1-ListRequest)
    - [ListResponse](#mods-v1-ListResponse)
    - [ListTeamsRequest](#mods-v1-ListTeamsRequest)
    - [ListTeamsResponse](#mods-v1-ListTeamsResponse)
    - [ListUsersRequest](#mods-v1-ListUsersRequest)
    - [ListUsersResponse](#mods-v1-ListUsersResponse)
    - [ShowRequest](#mods-v1-ShowRequest)
    - [ShowResponse](#mods-v1-ShowResponse)
    - [UpdateMod](#mods-v1-UpdateMod)
    - [UpdateRequest](#mods-v1-UpdateRequest)
    - [UpdateResponse](#mods-v1-UpdateResponse)
  
    - [ModsService](#mods-v1-ModsService)
  
- [types/v1/pack.proto](#types_v1_pack-proto)
    - [Pack](#types-v1-Pack)
    - [PackBack](#types-v1-PackBack)
    - [PackIcon](#types-v1-PackIcon)
    - [PackLogo](#types-v1-PackLogo)
  
- [packs/v1/packs.proto](#packs_v1_packs-proto)
    - [AttachTeamRequest](#packs-v1-AttachTeamRequest)
    - [AttachTeamResponse](#packs-v1-AttachTeamResponse)
    - [AttachUserRequest](#packs-v1-AttachUserRequest)
    - [AttachUserResponse](#packs-v1-AttachUserResponse)
    - [CreatePack](#packs-v1-CreatePack)
    - [CreateRequest](#packs-v1-CreateRequest)
    - [CreateResponse](#packs-v1-CreateResponse)
    - [DeleteRequest](#packs-v1-DeleteRequest)
    - [DeleteResponse](#packs-v1-DeleteResponse)
    - [DropTeamRequest](#packs-v1-DropTeamRequest)
    - [DropTeamResponse](#packs-v1-DropTeamResponse)
    - [DropUserRequest](#packs-v1-DropUserRequest)
    - [DropUserResponse](#packs-v1-DropUserResponse)
    - [ListRequest](#packs-v1-ListRequest)
    - [ListResponse](#packs-v1-ListResponse)
    - [ListTeamsRequest](#packs-v1-ListTeamsRequest)
    - [ListTeamsResponse](#packs-v1-ListTeamsResponse)
    - [ListUsersRequest](#packs-v1-ListUsersRequest)
    - [ListUsersResponse](#packs-v1-ListUsersResponse)
    - [ShowRequest](#packs-v1-ShowRequest)
    - [ShowResponse](#packs-v1-ShowResponse)
    - [UpdatePack](#packs-v1-UpdatePack)
    - [UpdateRequest](#packs-v1-UpdateRequest)
    - [UpdateResponse](#packs-v1-UpdateResponse)
  
    - [PacksService](#packs-v1-PacksService)
  
- [profile/v1/profile.proto](#profile_v1_profile-proto)
    - [ListModsRequest](#profile-v1-ListModsRequest)
    - [ListModsResponse](#profile-v1-ListModsResponse)
    - [ListPacksRequest](#profile-v1-ListPacksRequest)
    - [ListPacksResponse](#profile-v1-ListPacksResponse)
    - [ListTeamsRequest](#profile-v1-ListTeamsRequest)
    - [ListTeamsResponse](#profile-v1-ListTeamsResponse)
    - [ShowRequest](#profile-v1-ShowRequest)
    - [ShowResponse](#profile-v1-ShowResponse)
    - [UpdateRequest](#profile-v1-UpdateRequest)
    - [UpdateResponse](#profile-v1-UpdateResponse)
  
    - [ProfileService](#profile-v1-ProfileService)
  
- [teams/v1/teams.proto](#teams_v1_teams-proto)
    - [AttachModRequest](#teams-v1-AttachModRequest)
    - [AttachModResponse](#teams-v1-AttachModResponse)
    - [AttachPackRequest](#teams-v1-AttachPackRequest)
    - [AttachPackResponse](#teams-v1-AttachPackResponse)
    - [AttachUserRequest](#teams-v1-AttachUserRequest)
    - [AttachUserResponse](#teams-v1-AttachUserResponse)
    - [CreateRequest](#teams-v1-CreateRequest)
    - [CreateResponse](#teams-v1-CreateResponse)
    - [CreateTeam](#teams-v1-CreateTeam)
    - [DeleteRequest](#teams-v1-DeleteRequest)
    - [DeleteResponse](#teams-v1-DeleteResponse)
    - [DropModRequest](#teams-v1-DropModRequest)
    - [DropModResponse](#teams-v1-DropModResponse)
    - [DropPackRequest](#teams-v1-DropPackRequest)
    - [DropPackResponse](#teams-v1-DropPackResponse)
    - [DropUserRequest](#teams-v1-DropUserRequest)
    - [DropUserResponse](#teams-v1-DropUserResponse)
    - [ListModsRequest](#teams-v1-ListModsRequest)
    - [ListModsResponse](#teams-v1-ListModsResponse)
    - [ListPacksRequest](#teams-v1-ListPacksRequest)
    - [ListPacksResponse](#teams-v1-ListPacksResponse)
    - [ListRequest](#teams-v1-ListRequest)
    - [ListResponse](#teams-v1-ListResponse)
    - [ListUsersRequest](#teams-v1-ListUsersRequest)
    - [ListUsersResponse](#teams-v1-ListUsersResponse)
    - [ShowRequest](#teams-v1-ShowRequest)
    - [ShowResponse](#teams-v1-ShowResponse)
    - [UpdateRequest](#teams-v1-UpdateRequest)
    - [UpdateResponse](#teams-v1-UpdateResponse)
    - [UpdateTeam](#teams-v1-UpdateTeam)
  
    - [TeamsService](#teams-v1-TeamsService)
  
- [users/v1/users.proto](#users_v1_users-proto)
    - [AttachModRequest](#users-v1-AttachModRequest)
    - [AttachModResponse](#users-v1-AttachModResponse)
    - [AttachPackRequest](#users-v1-AttachPackRequest)
    - [AttachPackResponse](#users-v1-AttachPackResponse)
    - [AttachTeamRequest](#users-v1-AttachTeamRequest)
    - [AttachTeamResponse](#users-v1-AttachTeamResponse)
    - [CreateRequest](#users-v1-CreateRequest)
    - [CreateResponse](#users-v1-CreateResponse)
    - [CreateUser](#users-v1-CreateUser)
    - [DeleteRequest](#users-v1-DeleteRequest)
    - [DeleteResponse](#users-v1-DeleteResponse)
    - [DropModRequest](#users-v1-DropModRequest)
    - [DropModResponse](#users-v1-DropModResponse)
    - [DropPackRequest](#users-v1-DropPackRequest)
    - [DropPackResponse](#users-v1-DropPackResponse)
    - [DropTeamRequest](#users-v1-DropTeamRequest)
    - [DropTeamResponse](#users-v1-DropTeamResponse)
    - [ListModsRequest](#users-v1-ListModsRequest)
    - [ListModsResponse](#users-v1-ListModsResponse)
    - [ListPacksRequest](#users-v1-ListPacksRequest)
    - [ListPacksResponse](#users-v1-ListPacksResponse)
    - [ListRequest](#users-v1-ListRequest)
    - [ListResponse](#users-v1-ListResponse)
    - [ListTeamsRequest](#users-v1-ListTeamsRequest)
    - [ListTeamsResponse](#users-v1-ListTeamsResponse)
    - [ShowRequest](#users-v1-ShowRequest)
    - [ShowResponse](#users-v1-ShowResponse)
    - [UpdateRequest](#users-v1-UpdateRequest)
    - [UpdateResponse](#users-v1-UpdateResponse)
    - [UpdateUser](#users-v1-UpdateUser)
  
    - [UsersService](#users-v1-UsersService)
  
- [versions/v1/versions.proto](#versions_v1_versions-proto)
    - [AttachBuildRequest](#versions-v1-AttachBuildRequest)
    - [AttachBuildResponse](#versions-v1-AttachBuildResponse)
    - [CreateRequest](#versions-v1-CreateRequest)
    - [CreateResponse](#versions-v1-CreateResponse)
    - [CreateVersion](#versions-v1-CreateVersion)
    - [DeleteRequest](#versions-v1-DeleteRequest)
    - [DeleteResponse](#versions-v1-DeleteResponse)
    - [DropBuildRequest](#versions-v1-DropBuildRequest)
    - [DropBuildResponse](#versions-v1-DropBuildResponse)
    - [ListBuildsRequest](#versions-v1-ListBuildsRequest)
    - [ListBuildsResponse](#versions-v1-ListBuildsResponse)
    - [ListRequest](#versions-v1-ListRequest)
    - [ListResponse](#versions-v1-ListResponse)
    - [ShowRequest](#versions-v1-ShowRequest)
    - [ShowResponse](#versions-v1-ShowResponse)
    - [UpdateRequest](#versions-v1-UpdateRequest)
    - [UpdateResponse](#versions-v1-UpdateResponse)
    - [UpdateVersion](#versions-v1-UpdateVersion)
  
    - [VersionsService](#versions-v1-VersionsService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="types_v1_build-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/v1/build.proto



<a name="types-v1-Build"></a>

### Build



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| name | [string](#string) |  |  |
| minecraft | [string](#string) | optional |  |
| forge | [string](#string) | optional |  |
| java | [string](#string) |  |  |
| memory | [string](#string) |  |  |
| published | [bool](#bool) |  |  |
| private | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="types_v1_version-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/v1/version.proto



<a name="types-v1-Version"></a>

### Version



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| name | [string](#string) |  |  |
| file | [VersionFile](#types-v1-VersionFile) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="types-v1-VersionFile"></a>

### VersionFile



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| version_id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| content_type | [string](#string) |  |  |
| md5 | [string](#string) |  |  |
| path | [string](#string) |  |  |
| url | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="builds_v1_builds-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## builds/v1/builds.proto



<a name="builds-v1-AttachVersionRequest"></a>

### AttachVersionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| build | [string](#string) |  |  |
| mod | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="builds-v1-AttachVersionResponse"></a>

### AttachVersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="builds-v1-CreateBuild"></a>

### CreateBuild



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) |  |  |
| minecraft | [string](#string) | optional |  |
| forge | [string](#string) | optional |  |
| java | [string](#string) | optional |  |
| memory | [string](#string) | optional |  |
| published | [bool](#bool) | optional |  |
| private | [bool](#bool) | optional |  |






<a name="builds-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| build | [CreateBuild](#builds-v1-CreateBuild) |  |  |






<a name="builds-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| build | [types.v1.Build](#types-v1-Build) |  |  |






<a name="builds-v1-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="builds-v1-DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="builds-v1-DropVersionRequest"></a>

### DropVersionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| build | [string](#string) |  |  |
| mod | [string](#string) |  |  |
| version | [string](#string) |  |  |






<a name="builds-v1-DropVersionResponse"></a>

### DropVersionResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="builds-v1-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="builds-v1-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| builds | [types.v1.Build](#types-v1-Build) | repeated |  |






<a name="builds-v1-ListVersionsRequest"></a>

### ListVersionsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| build | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="builds-v1-ListVersionsResponse"></a>

### ListVersionsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| versions | [types.v1.Version](#types-v1-Version) | repeated |  |






<a name="builds-v1-ShowRequest"></a>

### ShowRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="builds-v1-ShowResponse"></a>

### ShowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| build | [types.v1.Build](#types-v1-Build) |  |  |






<a name="builds-v1-UpdateBuild"></a>

### UpdateBuild



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) | optional |  |
| minecraft | [string](#string) | optional |  |
| forge | [string](#string) | optional |  |
| java | [string](#string) | optional |  |
| memory | [string](#string) | optional |  |
| published | [bool](#bool) | optional |  |
| private | [bool](#bool) | optional |  |






<a name="builds-v1-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| id | [string](#string) |  |  |
| build | [UpdateBuild](#builds-v1-UpdateBuild) |  |  |






<a name="builds-v1-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| build | [types.v1.Build](#types-v1-Build) |  |  |





 

 

 


<a name="builds-v1-BuildsService"></a>

### BuildsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#builds-v1-ListRequest) | [ListResponse](#builds-v1-ListResponse) |  |
| Create | [CreateRequest](#builds-v1-CreateRequest) | [CreateResponse](#builds-v1-CreateResponse) |  |
| Update | [UpdateRequest](#builds-v1-UpdateRequest) | [UpdateResponse](#builds-v1-UpdateResponse) |  |
| Show | [ShowRequest](#builds-v1-ShowRequest) | [ShowResponse](#builds-v1-ShowResponse) |  |
| Delete | [DeleteRequest](#builds-v1-DeleteRequest) | [DeleteResponse](#builds-v1-DeleteResponse) |  |
| ListVersions | [ListVersionsRequest](#builds-v1-ListVersionsRequest) | [ListVersionsResponse](#builds-v1-ListVersionsResponse) |  |
| AttachVersion | [AttachVersionRequest](#builds-v1-AttachVersionRequest) | [AttachVersionResponse](#builds-v1-AttachVersionResponse) |  |
| DropVersion | [DropVersionRequest](#builds-v1-DropVersionRequest) | [DropVersionResponse](#builds-v1-DropVersionResponse) |  |

 



<a name="types_v1_forge-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/v1/forge.proto



<a name="types-v1-Forge"></a>

### Forge



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| minecraft | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="forge_v1_forge-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## forge/v1/forge.proto



<a name="forge-v1-ListBuildsRequest"></a>

### ListBuildsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forge | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="forge-v1-ListBuildsResponse"></a>

### ListBuildsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| builds | [types.v1.Build](#types-v1-Build) | repeated |  |






<a name="forge-v1-SearchRequest"></a>

### SearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="forge-v1-SearchResponse"></a>

### SearchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [types.v1.Forge](#types-v1-Forge) | repeated |  |






<a name="forge-v1-UpdateRequest"></a>

### UpdateRequest







<a name="forge-v1-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |





 

 

 


<a name="forge-v1-ForgeService"></a>

### ForgeService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Search | [SearchRequest](#forge-v1-SearchRequest) | [SearchResponse](#forge-v1-SearchResponse) |  |
| Update | [UpdateRequest](#forge-v1-UpdateRequest) | [UpdateResponse](#forge-v1-UpdateResponse) |  |
| ListBuilds | [ListBuildsRequest](#forge-v1-ListBuildsRequest) | [ListBuildsResponse](#forge-v1-ListBuildsResponse) |  |

 



<a name="types_v1_member-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/v1/member.proto



<a name="types-v1-Member"></a>

### Member



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team_id | [string](#string) |  |  |
| team_slug | [string](#string) |  |  |
| team_name | [string](#string) |  |  |
| user_id | [string](#string) |  |  |
| user_slug | [string](#string) |  |  |
| user_name | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="members_v1_members-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## members/v1/members.proto



<a name="members-v1-AppendMember"></a>

### AppendMember



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="members-v1-AppendRequest"></a>

### AppendRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member | [AppendMember](#members-v1-AppendMember) |  |  |






<a name="members-v1-AppendResponse"></a>

### AppendResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="members-v1-DropMember"></a>

### DropMember



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="members-v1-DropRequest"></a>

### DropRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member | [DropMember](#members-v1-DropMember) |  |  |






<a name="members-v1-DropResponse"></a>

### DropResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="members-v1-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="members-v1-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| members | [types.v1.Member](#types-v1-Member) | repeated |  |





 

 

 


<a name="members-v1-MembersService"></a>

### MembersService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#members-v1-ListRequest) | [ListResponse](#members-v1-ListResponse) |  |
| Append | [AppendRequest](#members-v1-AppendRequest) | [AppendResponse](#members-v1-AppendResponse) |  |
| Drop | [DropRequest](#members-v1-DropRequest) | [DropResponse](#members-v1-DropResponse) |  |

 



<a name="types_v1_minecraft-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/v1/minecraft.proto



<a name="types-v1-Minecraft"></a>

### Minecraft



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| type | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="minecraft_v1_minecraft-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## minecraft/v1/minecraft.proto



<a name="minecraft-v1-ListBuildsRequest"></a>

### ListBuildsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| minecraft | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="minecraft-v1-ListBuildsResponse"></a>

### ListBuildsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| builds | [types.v1.Build](#types-v1-Build) | repeated |  |






<a name="minecraft-v1-SearchRequest"></a>

### SearchRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="minecraft-v1-SearchResponse"></a>

### SearchResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [types.v1.Minecraft](#types-v1-Minecraft) | repeated |  |






<a name="minecraft-v1-UpdateRequest"></a>

### UpdateRequest







<a name="minecraft-v1-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |





 

 

 


<a name="minecraft-v1-MinecraftService"></a>

### MinecraftService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Search | [SearchRequest](#minecraft-v1-SearchRequest) | [SearchResponse](#minecraft-v1-SearchResponse) |  |
| Update | [UpdateRequest](#minecraft-v1-UpdateRequest) | [UpdateResponse](#minecraft-v1-UpdateResponse) |  |
| ListBuilds | [ListBuildsRequest](#minecraft-v1-ListBuildsRequest) | [ListBuildsResponse](#minecraft-v1-ListBuildsResponse) |  |

 



<a name="types_v1_mod-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/v1/mod.proto



<a name="types-v1-Mod"></a>

### Mod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| name | [string](#string) |  |  |
| side | [string](#string) |  |  |
| description | [string](#string) |  |  |
| author | [string](#string) |  |  |
| website | [string](#string) |  |  |
| donate | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| versions | [Version](#types-v1-Version) | repeated |  |





 

 

 

 



<a name="types_v1_team-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/v1/team.proto



<a name="types-v1-Team"></a>

### Team



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| name | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="types_v1_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/v1/user.proto



<a name="types-v1-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| username | [string](#string) |  |  |
| email | [string](#string) |  |  |
| firstname | [string](#string) |  |  |
| lastname | [string](#string) |  |  |
| admin | [bool](#bool) |  |  |
| active | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="mods_v1_mods-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## mods/v1/mods.proto



<a name="mods-v1-AttachTeamRequest"></a>

### AttachTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| team | [string](#string) |  |  |






<a name="mods-v1-AttachTeamResponse"></a>

### AttachTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="mods-v1-AttachUserRequest"></a>

### AttachUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="mods-v1-AttachUserResponse"></a>

### AttachUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="mods-v1-CreateMod"></a>

### CreateMod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) |  |  |
| side | [string](#string) | optional |  |
| description | [string](#string) | optional |  |
| author | [string](#string) | optional |  |
| website | [string](#string) | optional |  |
| donate | [string](#string) | optional |  |






<a name="mods-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [CreateMod](#mods-v1-CreateMod) |  |  |






<a name="mods-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [types.v1.Mod](#types-v1-Mod) |  |  |






<a name="mods-v1-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="mods-v1-DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="mods-v1-DropTeamRequest"></a>

### DropTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| team | [string](#string) |  |  |






<a name="mods-v1-DropTeamResponse"></a>

### DropTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="mods-v1-DropUserRequest"></a>

### DropUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="mods-v1-DropUserResponse"></a>

### DropUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="mods-v1-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="mods-v1-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mods | [types.v1.Mod](#types-v1-Mod) | repeated |  |






<a name="mods-v1-ListTeamsRequest"></a>

### ListTeamsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="mods-v1-ListTeamsResponse"></a>

### ListTeamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| teams | [types.v1.Team](#types-v1-Team) | repeated |  |






<a name="mods-v1-ListUsersRequest"></a>

### ListUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="mods-v1-ListUsersResponse"></a>

### ListUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [types.v1.User](#types-v1-User) | repeated |  |






<a name="mods-v1-ShowRequest"></a>

### ShowRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="mods-v1-ShowResponse"></a>

### ShowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [types.v1.Mod](#types-v1-Mod) |  |  |






<a name="mods-v1-UpdateMod"></a>

### UpdateMod



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) | optional |  |
| side | [string](#string) | optional |  |
| description | [string](#string) | optional |  |
| author | [string](#string) | optional |  |
| website | [string](#string) | optional |  |
| donate | [string](#string) | optional |  |






<a name="mods-v1-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| mod | [UpdateMod](#mods-v1-UpdateMod) |  |  |






<a name="mods-v1-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [types.v1.Mod](#types-v1-Mod) |  |  |





 

 

 


<a name="mods-v1-ModsService"></a>

### ModsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#mods-v1-ListRequest) | [ListResponse](#mods-v1-ListResponse) |  |
| Create | [CreateRequest](#mods-v1-CreateRequest) | [CreateResponse](#mods-v1-CreateResponse) |  |
| Update | [UpdateRequest](#mods-v1-UpdateRequest) | [UpdateResponse](#mods-v1-UpdateResponse) |  |
| Show | [ShowRequest](#mods-v1-ShowRequest) | [ShowResponse](#mods-v1-ShowResponse) |  |
| Delete | [DeleteRequest](#mods-v1-DeleteRequest) | [DeleteResponse](#mods-v1-DeleteResponse) |  |
| ListUsers | [ListUsersRequest](#mods-v1-ListUsersRequest) | [ListUsersResponse](#mods-v1-ListUsersResponse) |  |
| AttachUser | [AttachUserRequest](#mods-v1-AttachUserRequest) | [AttachUserResponse](#mods-v1-AttachUserResponse) |  |
| DropUser | [DropUserRequest](#mods-v1-DropUserRequest) | [DropUserResponse](#mods-v1-DropUserResponse) |  |
| ListTeams | [ListTeamsRequest](#mods-v1-ListTeamsRequest) | [ListTeamsResponse](#mods-v1-ListTeamsResponse) |  |
| AttachTeam | [AttachTeamRequest](#mods-v1-AttachTeamRequest) | [AttachTeamResponse](#mods-v1-AttachTeamResponse) |  |
| DropTeam | [DropTeamRequest](#mods-v1-DropTeamRequest) | [DropTeamResponse](#mods-v1-DropTeamResponse) |  |

 



<a name="types_v1_pack-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## types/v1/pack.proto



<a name="types-v1-Pack"></a>

### Pack



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| name | [string](#string) |  |  |
| website | [string](#string) |  |  |
| back | [PackBack](#types-v1-PackBack) | optional |  |
| icon | [PackIcon](#types-v1-PackIcon) | optional |  |
| logo | [PackLogo](#types-v1-PackLogo) | optional |  |
| published | [bool](#bool) |  |  |
| private | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| builds | [Build](#types-v1-Build) | repeated |  |






<a name="types-v1-PackBack"></a>

### PackBack



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pack_id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| content_type | [string](#string) |  |  |
| md5 | [string](#string) |  |  |
| path | [string](#string) |  |  |
| url | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="types-v1-PackIcon"></a>

### PackIcon



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pack_id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| content_type | [string](#string) |  |  |
| md5 | [string](#string) |  |  |
| path | [string](#string) |  |  |
| url | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="types-v1-PackLogo"></a>

### PackLogo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pack_id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| content_type | [string](#string) |  |  |
| md5 | [string](#string) |  |  |
| path | [string](#string) |  |  |
| url | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 

 



<a name="packs_v1_packs-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## packs/v1/packs.proto



<a name="packs-v1-AttachTeamRequest"></a>

### AttachTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| team | [string](#string) |  |  |






<a name="packs-v1-AttachTeamResponse"></a>

### AttachTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="packs-v1-AttachUserRequest"></a>

### AttachUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="packs-v1-AttachUserResponse"></a>

### AttachUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="packs-v1-CreatePack"></a>

### CreatePack



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) |  |  |
| website | [string](#string) | optional |  |
| icon | [string](#string) | optional |  |
| logo | [string](#string) | optional |  |
| bg | [string](#string) | optional |  |
| published | [bool](#bool) | optional |  |
| private | [bool](#bool) | optional |  |






<a name="packs-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [CreatePack](#packs-v1-CreatePack) |  |  |






<a name="packs-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [types.v1.Pack](#types-v1-Pack) |  |  |






<a name="packs-v1-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="packs-v1-DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="packs-v1-DropTeamRequest"></a>

### DropTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| team | [string](#string) |  |  |






<a name="packs-v1-DropTeamResponse"></a>

### DropTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="packs-v1-DropUserRequest"></a>

### DropUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="packs-v1-DropUserResponse"></a>

### DropUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="packs-v1-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="packs-v1-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| packs | [types.v1.Pack](#types-v1-Pack) | repeated |  |






<a name="packs-v1-ListTeamsRequest"></a>

### ListTeamsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="packs-v1-ListTeamsResponse"></a>

### ListTeamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| teams | [types.v1.Team](#types-v1-Team) | repeated |  |






<a name="packs-v1-ListUsersRequest"></a>

### ListUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="packs-v1-ListUsersResponse"></a>

### ListUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [types.v1.User](#types-v1-User) | repeated |  |






<a name="packs-v1-ShowRequest"></a>

### ShowRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="packs-v1-ShowResponse"></a>

### ShowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [types.v1.Pack](#types-v1-Pack) |  |  |






<a name="packs-v1-UpdatePack"></a>

### UpdatePack



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) | optional |  |
| website | [string](#string) | optional |  |
| icon | [string](#string) | optional |  |
| logo | [string](#string) | optional |  |
| bg | [string](#string) | optional |  |
| recommended | [string](#string) | optional |  |
| latest | [string](#string) | optional |  |
| published | [bool](#bool) | optional |  |
| private | [bool](#bool) | optional |  |






<a name="packs-v1-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| pack | [UpdatePack](#packs-v1-UpdatePack) |  |  |






<a name="packs-v1-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [types.v1.Pack](#types-v1-Pack) |  |  |





 

 

 


<a name="packs-v1-PacksService"></a>

### PacksService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#packs-v1-ListRequest) | [ListResponse](#packs-v1-ListResponse) |  |
| Create | [CreateRequest](#packs-v1-CreateRequest) | [CreateResponse](#packs-v1-CreateResponse) |  |
| Update | [UpdateRequest](#packs-v1-UpdateRequest) | [UpdateResponse](#packs-v1-UpdateResponse) |  |
| Show | [ShowRequest](#packs-v1-ShowRequest) | [ShowResponse](#packs-v1-ShowResponse) |  |
| Delete | [DeleteRequest](#packs-v1-DeleteRequest) | [DeleteResponse](#packs-v1-DeleteResponse) |  |
| ListUsers | [ListUsersRequest](#packs-v1-ListUsersRequest) | [ListUsersResponse](#packs-v1-ListUsersResponse) |  |
| AttachUser | [AttachUserRequest](#packs-v1-AttachUserRequest) | [AttachUserResponse](#packs-v1-AttachUserResponse) |  |
| DropUser | [DropUserRequest](#packs-v1-DropUserRequest) | [DropUserResponse](#packs-v1-DropUserResponse) |  |
| ListTeams | [ListTeamsRequest](#packs-v1-ListTeamsRequest) | [ListTeamsResponse](#packs-v1-ListTeamsResponse) |  |
| AttachTeam | [AttachTeamRequest](#packs-v1-AttachTeamRequest) | [AttachTeamResponse](#packs-v1-AttachTeamResponse) |  |
| DropTeam | [DropTeamRequest](#packs-v1-DropTeamRequest) | [DropTeamResponse](#packs-v1-DropTeamResponse) |  |

 



<a name="profile_v1_profile-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## profile/v1/profile.proto



<a name="profile-v1-ListModsRequest"></a>

### ListModsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="profile-v1-ListModsResponse"></a>

### ListModsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mods | [types.v1.Mod](#types-v1-Mod) | repeated |  |






<a name="profile-v1-ListPacksRequest"></a>

### ListPacksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="profile-v1-ListPacksResponse"></a>

### ListPacksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| packs | [types.v1.Pack](#types-v1-Pack) | repeated |  |






<a name="profile-v1-ListTeamsRequest"></a>

### ListTeamsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="profile-v1-ListTeamsResponse"></a>

### ListTeamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| teams | [types.v1.Team](#types-v1-Team) | repeated |  |






<a name="profile-v1-ShowRequest"></a>

### ShowRequest







<a name="profile-v1-ShowResponse"></a>

### ShowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| profile | [types.v1.User](#types-v1-User) |  |  |






<a name="profile-v1-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| username | [string](#string) | optional |  |
| password | [string](#string) | optional |  |
| email | [string](#string) | optional |  |
| firstname | [string](#string) | optional |  |
| lastname | [string](#string) | optional |  |






<a name="profile-v1-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |





 

 

 


<a name="profile-v1-ProfileService"></a>

### ProfileService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Show | [ShowRequest](#profile-v1-ShowRequest) | [ShowResponse](#profile-v1-ShowResponse) |  |
| Update | [UpdateRequest](#profile-v1-UpdateRequest) | [UpdateResponse](#profile-v1-UpdateResponse) |  |
| ListTeams | [ListTeamsRequest](#profile-v1-ListTeamsRequest) | [ListTeamsResponse](#profile-v1-ListTeamsResponse) |  |
| ListPacks | [ListPacksRequest](#profile-v1-ListPacksRequest) | [ListPacksResponse](#profile-v1-ListPacksResponse) |  |
| ListMods | [ListModsRequest](#profile-v1-ListModsRequest) | [ListModsResponse](#profile-v1-ListModsResponse) |  |

 



<a name="teams_v1_teams-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## teams/v1/teams.proto



<a name="teams-v1-AttachModRequest"></a>

### AttachModRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| mod | [string](#string) |  |  |






<a name="teams-v1-AttachModResponse"></a>

### AttachModResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="teams-v1-AttachPackRequest"></a>

### AttachPackRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| pack | [string](#string) |  |  |






<a name="teams-v1-AttachPackResponse"></a>

### AttachPackResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="teams-v1-AttachUserRequest"></a>

### AttachUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="teams-v1-AttachUserResponse"></a>

### AttachUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="teams-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [CreateTeam](#teams-v1-CreateTeam) |  |  |






<a name="teams-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [types.v1.Team](#types-v1-Team) |  |  |






<a name="teams-v1-CreateTeam"></a>

### CreateTeam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) |  |  |






<a name="teams-v1-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="teams-v1-DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="teams-v1-DropModRequest"></a>

### DropModRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| mod | [string](#string) |  |  |






<a name="teams-v1-DropModResponse"></a>

### DropModResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="teams-v1-DropPackRequest"></a>

### DropPackRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| pack | [string](#string) |  |  |






<a name="teams-v1-DropPackResponse"></a>

### DropPackResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="teams-v1-DropUserRequest"></a>

### DropUserRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| user | [string](#string) |  |  |






<a name="teams-v1-DropUserResponse"></a>

### DropUserResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="teams-v1-ListModsRequest"></a>

### ListModsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="teams-v1-ListModsResponse"></a>

### ListModsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mods | [types.v1.Mod](#types-v1-Mod) | repeated |  |






<a name="teams-v1-ListPacksRequest"></a>

### ListPacksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="teams-v1-ListPacksResponse"></a>

### ListPacksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| packs | [types.v1.Pack](#types-v1-Pack) | repeated |  |






<a name="teams-v1-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="teams-v1-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| teams | [types.v1.Team](#types-v1-Team) | repeated |  |






<a name="teams-v1-ListUsersRequest"></a>

### ListUsersRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="teams-v1-ListUsersResponse"></a>

### ListUsersResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [types.v1.User](#types-v1-User) | repeated |  |






<a name="teams-v1-ShowRequest"></a>

### ShowRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="teams-v1-ShowResponse"></a>

### ShowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [types.v1.Team](#types-v1-Team) |  |  |






<a name="teams-v1-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| team | [UpdateTeam](#teams-v1-UpdateTeam) |  |  |






<a name="teams-v1-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [types.v1.Team](#types-v1-Team) |  |  |






<a name="teams-v1-UpdateTeam"></a>

### UpdateTeam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) | optional |  |





 

 

 


<a name="teams-v1-TeamsService"></a>

### TeamsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#teams-v1-ListRequest) | [ListResponse](#teams-v1-ListResponse) |  |
| Create | [CreateRequest](#teams-v1-CreateRequest) | [CreateResponse](#teams-v1-CreateResponse) |  |
| Update | [UpdateRequest](#teams-v1-UpdateRequest) | [UpdateResponse](#teams-v1-UpdateResponse) |  |
| Show | [ShowRequest](#teams-v1-ShowRequest) | [ShowResponse](#teams-v1-ShowResponse) |  |
| Delete | [DeleteRequest](#teams-v1-DeleteRequest) | [DeleteResponse](#teams-v1-DeleteResponse) |  |
| ListUsers | [ListUsersRequest](#teams-v1-ListUsersRequest) | [ListUsersResponse](#teams-v1-ListUsersResponse) |  |
| AttachUser | [AttachUserRequest](#teams-v1-AttachUserRequest) | [AttachUserResponse](#teams-v1-AttachUserResponse) |  |
| DropUser | [DropUserRequest](#teams-v1-DropUserRequest) | [DropUserResponse](#teams-v1-DropUserResponse) |  |
| ListPacks | [ListPacksRequest](#teams-v1-ListPacksRequest) | [ListPacksResponse](#teams-v1-ListPacksResponse) |  |
| AttachPack | [AttachPackRequest](#teams-v1-AttachPackRequest) | [AttachPackResponse](#teams-v1-AttachPackResponse) |  |
| DropPack | [DropPackRequest](#teams-v1-DropPackRequest) | [DropPackResponse](#teams-v1-DropPackResponse) |  |
| ListMods | [ListModsRequest](#teams-v1-ListModsRequest) | [ListModsResponse](#teams-v1-ListModsResponse) |  |
| AttachMod | [AttachModRequest](#teams-v1-AttachModRequest) | [AttachModResponse](#teams-v1-AttachModResponse) |  |
| DropMod | [DropModRequest](#teams-v1-DropModRequest) | [DropModResponse](#teams-v1-DropModResponse) |  |

 



<a name="users_v1_users-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## users/v1/users.proto



<a name="users-v1-AttachModRequest"></a>

### AttachModRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| mod | [string](#string) |  |  |






<a name="users-v1-AttachModResponse"></a>

### AttachModResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="users-v1-AttachPackRequest"></a>

### AttachPackRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| pack | [string](#string) |  |  |






<a name="users-v1-AttachPackResponse"></a>

### AttachPackResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="users-v1-AttachTeamRequest"></a>

### AttachTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| team | [string](#string) |  |  |






<a name="users-v1-AttachTeamResponse"></a>

### AttachTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="users-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [CreateUser](#users-v1-CreateUser) |  |  |






<a name="users-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [types.v1.User](#types-v1-User) |  |  |






<a name="users-v1-CreateUser"></a>

### CreateUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| email | [string](#string) |  |  |
| firstname | [string](#string) | optional |  |
| lastname | [string](#string) | optional |  |
| admin | [bool](#bool) |  |  |
| active | [bool](#bool) |  |  |






<a name="users-v1-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="users-v1-DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="users-v1-DropModRequest"></a>

### DropModRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| mod | [string](#string) |  |  |






<a name="users-v1-DropModResponse"></a>

### DropModResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="users-v1-DropPackRequest"></a>

### DropPackRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| pack | [string](#string) |  |  |






<a name="users-v1-DropPackResponse"></a>

### DropPackResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="users-v1-DropTeamRequest"></a>

### DropTeamRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| team | [string](#string) |  |  |






<a name="users-v1-DropTeamResponse"></a>

### DropTeamResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="users-v1-ListModsRequest"></a>

### ListModsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="users-v1-ListModsResponse"></a>

### ListModsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mods | [types.v1.Mod](#types-v1-Mod) | repeated |  |






<a name="users-v1-ListPacksRequest"></a>

### ListPacksRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="users-v1-ListPacksResponse"></a>

### ListPacksResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| packs | [types.v1.Pack](#types-v1-Pack) | repeated |  |






<a name="users-v1-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="users-v1-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [types.v1.User](#types-v1-User) | repeated |  |






<a name="users-v1-ListTeamsRequest"></a>

### ListTeamsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="users-v1-ListTeamsResponse"></a>

### ListTeamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| teams | [types.v1.Team](#types-v1-Team) | repeated |  |






<a name="users-v1-ShowRequest"></a>

### ShowRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="users-v1-ShowResponse"></a>

### ShowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [types.v1.User](#types-v1-User) |  |  |






<a name="users-v1-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| user | [UpdateUser](#users-v1-UpdateUser) |  |  |






<a name="users-v1-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [types.v1.User](#types-v1-User) |  |  |






<a name="users-v1-UpdateUser"></a>

### UpdateUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| username | [string](#string) | optional |  |
| password | [string](#string) | optional |  |
| email | [string](#string) | optional |  |
| firstname | [string](#string) | optional |  |
| lastname | [string](#string) | optional |  |
| admin | [bool](#bool) | optional |  |
| active | [bool](#bool) | optional |  |





 

 

 


<a name="users-v1-UsersService"></a>

### UsersService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#users-v1-ListRequest) | [ListResponse](#users-v1-ListResponse) |  |
| Create | [CreateRequest](#users-v1-CreateRequest) | [CreateResponse](#users-v1-CreateResponse) |  |
| Update | [UpdateRequest](#users-v1-UpdateRequest) | [UpdateResponse](#users-v1-UpdateResponse) |  |
| Show | [ShowRequest](#users-v1-ShowRequest) | [ShowResponse](#users-v1-ShowResponse) |  |
| Delete | [DeleteRequest](#users-v1-DeleteRequest) | [DeleteResponse](#users-v1-DeleteResponse) |  |
| ListTeams | [ListTeamsRequest](#users-v1-ListTeamsRequest) | [ListTeamsResponse](#users-v1-ListTeamsResponse) |  |
| AttachTeam | [AttachTeamRequest](#users-v1-AttachTeamRequest) | [AttachTeamResponse](#users-v1-AttachTeamResponse) |  |
| DropTeam | [DropTeamRequest](#users-v1-DropTeamRequest) | [DropTeamResponse](#users-v1-DropTeamResponse) |  |
| ListPacks | [ListPacksRequest](#users-v1-ListPacksRequest) | [ListPacksResponse](#users-v1-ListPacksResponse) |  |
| AttachPack | [AttachPackRequest](#users-v1-AttachPackRequest) | [AttachPackResponse](#users-v1-AttachPackResponse) |  |
| DropPack | [DropPackRequest](#users-v1-DropPackRequest) | [DropPackResponse](#users-v1-DropPackResponse) |  |
| ListMods | [ListModsRequest](#users-v1-ListModsRequest) | [ListModsResponse](#users-v1-ListModsResponse) |  |
| AttachMod | [AttachModRequest](#users-v1-AttachModRequest) | [AttachModResponse](#users-v1-AttachModResponse) |  |
| DropMod | [DropModRequest](#users-v1-DropModRequest) | [DropModResponse](#users-v1-DropModResponse) |  |

 



<a name="versions_v1_versions-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## versions/v1/versions.proto



<a name="versions-v1-AttachBuildRequest"></a>

### AttachBuildRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| version | [string](#string) |  |  |
| pack | [string](#string) |  |  |
| build | [string](#string) |  |  |






<a name="versions-v1-AttachBuildResponse"></a>

### AttachBuildResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="versions-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| version | [CreateVersion](#versions-v1-CreateVersion) |  |  |






<a name="versions-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [types.v1.Version](#types-v1-Version) |  |  |






<a name="versions-v1-CreateVersion"></a>

### CreateVersion



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) |  |  |
| file | [string](#string) |  |  |






<a name="versions-v1-DeleteRequest"></a>

### DeleteRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="versions-v1-DeleteResponse"></a>

### DeleteResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="versions-v1-DropBuildRequest"></a>

### DropBuildRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| version | [string](#string) |  |  |
| pack | [string](#string) |  |  |
| build | [string](#string) |  |  |






<a name="versions-v1-DropBuildResponse"></a>

### DropBuildResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message | [string](#string) |  |  |






<a name="versions-v1-ListBuildsRequest"></a>

### ListBuildsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| version | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="versions-v1-ListBuildsResponse"></a>

### ListBuildsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| builds | [types.v1.Build](#types-v1-Build) | repeated |  |






<a name="versions-v1-ListRequest"></a>

### ListRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| query | [string](#string) |  |  |






<a name="versions-v1-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| versions | [types.v1.Version](#types-v1-Version) | repeated |  |






<a name="versions-v1-ShowRequest"></a>

### ShowRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| id | [string](#string) |  |  |






<a name="versions-v1-ShowResponse"></a>

### ShowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [types.v1.Version](#types-v1-Version) |  |  |






<a name="versions-v1-UpdateRequest"></a>

### UpdateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mod | [string](#string) |  |  |
| id | [string](#string) |  |  |
| version | [UpdateVersion](#versions-v1-UpdateVersion) |  |  |






<a name="versions-v1-UpdateResponse"></a>

### UpdateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| version | [types.v1.Version](#types-v1-Version) |  |  |






<a name="versions-v1-UpdateVersion"></a>

### UpdateVersion



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) | optional |  |
| name | [string](#string) | optional |  |
| file | [string](#string) | optional |  |





 

 

 


<a name="versions-v1-VersionsService"></a>

### VersionsService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#versions-v1-ListRequest) | [ListResponse](#versions-v1-ListResponse) |  |
| Create | [CreateRequest](#versions-v1-CreateRequest) | [CreateResponse](#versions-v1-CreateResponse) |  |
| Update | [UpdateRequest](#versions-v1-UpdateRequest) | [UpdateResponse](#versions-v1-UpdateResponse) |  |
| Show | [ShowRequest](#versions-v1-ShowRequest) | [ShowResponse](#versions-v1-ShowResponse) |  |
| Delete | [DeleteRequest](#versions-v1-DeleteRequest) | [DeleteResponse](#versions-v1-DeleteResponse) |  |
| ListBuilds | [ListBuildsRequest](#versions-v1-ListBuildsRequest) | [ListBuildsResponse](#versions-v1-ListBuildsResponse) |  |
| AttachBuild | [AttachBuildRequest](#versions-v1-AttachBuildRequest) | [AttachBuildResponse](#versions-v1-AttachBuildResponse) |  |
| DropBuild | [DropBuildRequest](#versions-v1-DropBuildRequest) | [DropBuildResponse](#versions-v1-DropBuildResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

