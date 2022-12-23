# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [members/v1/members.proto](#members_v1_members-proto)
    - [AppendMember](#members-v1-AppendMember)
    - [AppendRequest](#members-v1-AppendRequest)
    - [AppendResponse](#members-v1-AppendResponse)
    - [DropMember](#members-v1-DropMember)
    - [DropRequest](#members-v1-DropRequest)
    - [DropResponse](#members-v1-DropResponse)
    - [ListRequest](#members-v1-ListRequest)
    - [ListResponse](#members-v1-ListResponse)
    - [Member](#members-v1-Member)
  
    - [MembersService](#members-v1-MembersService)
  
- [teams/v1/teams.proto](#teams_v1_teams-proto)
    - [CreateRequest](#teams-v1-CreateRequest)
    - [CreateResponse](#teams-v1-CreateResponse)
    - [CreateTeam](#teams-v1-CreateTeam)
    - [DeleteRequest](#teams-v1-DeleteRequest)
    - [DeleteResponse](#teams-v1-DeleteResponse)
    - [ListRequest](#teams-v1-ListRequest)
    - [ListResponse](#teams-v1-ListResponse)
    - [ShowRequest](#teams-v1-ShowRequest)
    - [ShowResponse](#teams-v1-ShowResponse)
    - [Team](#teams-v1-Team)
    - [UpdateRequest](#teams-v1-UpdateRequest)
    - [UpdateResponse](#teams-v1-UpdateResponse)
    - [UpdateTeam](#teams-v1-UpdateTeam)
  
    - [TeamsService](#teams-v1-TeamsService)
  
- [users/v1/users.proto](#users_v1_users-proto)
    - [CreateRequest](#users-v1-CreateRequest)
    - [CreateResponse](#users-v1-CreateResponse)
    - [CreateUser](#users-v1-CreateUser)
    - [DeleteRequest](#users-v1-DeleteRequest)
    - [DeleteResponse](#users-v1-DeleteResponse)
    - [ListRequest](#users-v1-ListRequest)
    - [ListResponse](#users-v1-ListResponse)
    - [ShowRequest](#users-v1-ShowRequest)
    - [ShowResponse](#users-v1-ShowResponse)
    - [UpdateRequest](#users-v1-UpdateRequest)
    - [UpdateResponse](#users-v1-UpdateResponse)
    - [UpdateUser](#users-v1-UpdateUser)
    - [User](#users-v1-User)
  
    - [UsersService](#users-v1-UsersService)
  
- [Scalar Value Types](#scalar-value-types)



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
| members | [Member](#members-v1-Member) | repeated |  |






<a name="members-v1-Member"></a>

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





 

 

 


<a name="members-v1-MembersService"></a>

### MembersService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#members-v1-ListRequest) | [ListResponse](#members-v1-ListResponse) |  |
| Append | [AppendRequest](#members-v1-AppendRequest) | [AppendResponse](#members-v1-AppendResponse) |  |
| Drop | [DropRequest](#members-v1-DropRequest) | [DropResponse](#members-v1-DropResponse) |  |

 



<a name="teams_v1_teams-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## teams/v1/teams.proto



<a name="teams-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [CreateTeam](#teams-v1-CreateTeam) |  |  |






<a name="teams-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [Team](#teams-v1-Team) |  |  |






<a name="teams-v1-CreateTeam"></a>

### CreateTeam



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) |  |  |
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






<a name="teams-v1-ListRequest"></a>

### ListRequest







<a name="teams-v1-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| teams | [Team](#teams-v1-Team) | repeated |  |






<a name="teams-v1-ShowRequest"></a>

### ShowRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="teams-v1-ShowResponse"></a>

### ShowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| team | [Team](#teams-v1-Team) |  |  |






<a name="teams-v1-Team"></a>

### Team



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| slug | [string](#string) |  |  |
| name | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






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
| team | [Team](#teams-v1-Team) |  |  |






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

 



<a name="users_v1_users-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## users/v1/users.proto



<a name="users-v1-CreateRequest"></a>

### CreateRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [CreateUser](#users-v1-CreateUser) |  |  |






<a name="users-v1-CreateResponse"></a>

### CreateResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#users-v1-User) |  |  |






<a name="users-v1-CreateUser"></a>

### CreateUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slug | [string](#string) |  |  |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |
| email | [string](#string) |  |  |
| firstname | [string](#string) |  |  |
| lastname | [string](#string) |  |  |
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






<a name="users-v1-ListRequest"></a>

### ListRequest







<a name="users-v1-ListResponse"></a>

### ListResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#users-v1-User) | repeated |  |






<a name="users-v1-ShowRequest"></a>

### ShowRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="users-v1-ShowResponse"></a>

### ShowResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#users-v1-User) |  |  |






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
| user | [User](#users-v1-User) |  |  |






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






<a name="users-v1-User"></a>

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





 

 

 


<a name="users-v1-UsersService"></a>

### UsersService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| List | [ListRequest](#users-v1-ListRequest) | [ListResponse](#users-v1-ListResponse) |  |
| Create | [CreateRequest](#users-v1-CreateRequest) | [CreateResponse](#users-v1-CreateResponse) |  |
| Update | [UpdateRequest](#users-v1-UpdateRequest) | [UpdateResponse](#users-v1-UpdateResponse) |  |
| Show | [ShowRequest](#users-v1-ShowRequest) | [ShowResponse](#users-v1-ShowResponse) |  |
| Delete | [DeleteRequest](#users-v1-DeleteRequest) | [DeleteResponse](#users-v1-DeleteResponse) |  |

 



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

