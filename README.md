
# UralCTF API.
Бэкенд сайта UralCTF.
  

## Informations

### Version

0.1.0

## Content negotiation

### URI Schemes
  * http
  * https

### Consumes
  * application/json

### Produces
  * application/json

## All endpoints

###  search

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/api/search/city | [search cities](#search-cities) |  |
  


###  teams

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/api/teams/check-name | [check team name](#check-team-name) | Проверка уникальности имени команды. |
| POST | /api/api/teams | [create team](#create-team) |  |
| GET | /api/api/teams | [list teams](#list-teams) | Получение списка команд. |
  


###  universities

| Method  | URI     | Name   | Summary |
|---------|---------|--------|---------|
| GET | /api/api/search/university | [search universities](#search-universities) |  |
  


## Paths

### <span id="check-team-name"></span> Проверка уникальности имени команды. (*checkTeamName*)

```
GET /api/api/teams/check-name
```

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| name | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#check-team-name-200) | OK |  |  | [schema](#check-team-name-200-schema) |
| [400](#check-team-name-400) | Bad Request |  |  | [schema](#check-team-name-400-schema) |
| [500](#check-team-name-500) | Internal Server Error |  |  | [schema](#check-team-name-500-schema) |

#### Responses


##### <span id="check-team-name-200"></span> 200
Status: OK

###### <span id="check-team-name-200-schema"></span> Schema
   
  

[CheckTeamNameOKBody](#check-team-name-o-k-body)

##### <span id="check-team-name-400"></span> 400
Status: Bad Request

###### <span id="check-team-name-400-schema"></span> Schema
   
  

[CheckTeamNameBadRequestBody](#check-team-name-bad-request-body)

##### <span id="check-team-name-500"></span> 500
Status: Internal Server Error

###### <span id="check-team-name-500-schema"></span> Schema
   
  

[CheckTeamNameInternalServerErrorBody](#check-team-name-internal-server-error-body)

###### Inlined models

**<span id="check-team-name-bad-request-body"></span> CheckTeamNameBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



**<span id="check-team-name-internal-server-error-body"></span> CheckTeamNameInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



**<span id="check-team-name-o-k-body"></span> CheckTeamNameOKBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Available | boolean| `bool` |  | |  |  |



### <span id="create-team"></span> create team (*createTeam*)

```
POST /api/api/teams
```

Создание новой команды и получение ID команды

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| Body | `body` | [CreateTeamRequest](#create-team-request) | `models.CreateTeamRequest` | |  | |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [201](#create-team-201) | Created |  |  | [schema](#create-team-201-schema) |
| [400](#create-team-400) | Bad Request |  |  | [schema](#create-team-400-schema) |
| [500](#create-team-500) | Internal Server Error |  |  | [schema](#create-team-500-schema) |

#### Responses


##### <span id="create-team-201"></span> 201
Status: Created

###### <span id="create-team-201-schema"></span> Schema
   
  

[CreateTeamCreatedBody](#create-team-created-body)

##### <span id="create-team-400"></span> 400
Status: Bad Request

###### <span id="create-team-400-schema"></span> Schema
   
  

[CreateTeamBadRequestBody](#create-team-bad-request-body)

##### <span id="create-team-500"></span> 500
Status: Internal Server Error

###### <span id="create-team-500-schema"></span> Schema
   
  

[CreateTeamInternalServerErrorBody](#create-team-internal-server-error-body)

###### Inlined models

**<span id="create-team-bad-request-body"></span> CreateTeamBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



**<span id="create-team-created-body"></span> CreateTeamCreatedBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| TeamID | int64 (formatted integer)| `int64` |  | |  |  |



**<span id="create-team-internal-server-error-body"></span> CreateTeamInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



### <span id="list-teams"></span> Получение списка команд. (*listTeams*)

```
GET /api/api/teams
```

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#list-teams-200) | OK |  |  | [schema](#list-teams-200-schema) |
| [404](#list-teams-404) | Not Found |  |  | [schema](#list-teams-404-schema) |
| [500](#list-teams-500) | Internal Server Error |  |  | [schema](#list-teams-500-schema) |

#### Responses


##### <span id="list-teams-200"></span> 200
Status: OK

###### <span id="list-teams-200-schema"></span> Schema
   
  

[][Team](#team)

##### <span id="list-teams-404"></span> 404
Status: Not Found

###### <span id="list-teams-404-schema"></span> Schema
   
  

[ListTeamsNotFoundBody](#list-teams-not-found-body)

##### <span id="list-teams-500"></span> 500
Status: Internal Server Error

###### <span id="list-teams-500-schema"></span> Schema
   
  

[ListTeamsInternalServerErrorBody](#list-teams-internal-server-error-body)

###### Inlined models

**<span id="list-teams-internal-server-error-body"></span> ListTeamsInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



**<span id="list-teams-not-found-body"></span> ListTeamsNotFoundBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



### <span id="search-cities"></span> search cities (*searchCities*)

```
GET /api/api/search/city
```

Поиск городов по параметру запроса

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#search-cities-200) | OK |  |  | [schema](#search-cities-200-schema) |
| [500](#search-cities-500) | Internal Server Error |  |  | [schema](#search-cities-500-schema) |

#### Responses


##### <span id="search-cities-200"></span> 200
Status: OK

###### <span id="search-cities-200-schema"></span> Schema
   
  

[][CitySearchResult](#city-search-result)

##### <span id="search-cities-500"></span> 500
Status: Internal Server Error

###### <span id="search-cities-500-schema"></span> Schema
   
  

[SearchCitiesInternalServerErrorBody](#search-cities-internal-server-error-body)

###### Inlined models

**<span id="search-cities-internal-server-error-body"></span> SearchCitiesInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



### <span id="search-universities"></span> search universities (*searchUniversities*)

```
GET /api/api/search/university
```

Поиск университета по городу

#### Parameters

| Name | Source | Type | Go type | Separator | Required | Default | Description |
|------|--------|------|---------|-----------| :------: |---------|-------------|
| city | `query` | string | `string` |  |  |  |  |

#### All responses
| Code | Status | Description | Has headers | Schema |
|------|--------|-------------|:-----------:|--------|
| [200](#search-universities-200) | OK |  |  | [schema](#search-universities-200-schema) |
| [400](#search-universities-400) | Bad Request |  |  | [schema](#search-universities-400-schema) |
| [404](#search-universities-404) | Not Found |  |  | [schema](#search-universities-404-schema) |
| [500](#search-universities-500) | Internal Server Error |  |  | [schema](#search-universities-500-schema) |

#### Responses


##### <span id="search-universities-200"></span> 200
Status: OK

###### <span id="search-universities-200-schema"></span> Schema
   
  

[][University](#university)

##### <span id="search-universities-400"></span> 400
Status: Bad Request

###### <span id="search-universities-400-schema"></span> Schema
   
  

[SearchUniversitiesBadRequestBody](#search-universities-bad-request-body)

##### <span id="search-universities-404"></span> 404
Status: Not Found

###### <span id="search-universities-404-schema"></span> Schema
   
  

[SearchUniversitiesNotFoundBody](#search-universities-not-found-body)

##### <span id="search-universities-500"></span> 500
Status: Internal Server Error

###### <span id="search-universities-500-schema"></span> Schema
   
  

[SearchUniversitiesInternalServerErrorBody](#search-universities-internal-server-error-body)

###### Inlined models

**<span id="search-universities-bad-request-body"></span> SearchUniversitiesBadRequestBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



**<span id="search-universities-internal-server-error-body"></span> SearchUniversitiesInternalServerErrorBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



**<span id="search-universities-not-found-body"></span> SearchUniversitiesNotFoundBody**


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Error | string| `string` |  | |  |  |



## Models

### <span id="city-search-result"></span> CitySearchResult


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | int64 (formatted integer)| `int64` |  | |  |  |
| Name | string| `string` |  | |  |  |



### <span id="create-team-request"></span> CreateTeamRequest


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| City | string| `string` | ✓ | | City name |  |
| ConsentPDCapitan | boolean| `bool` | ✓ | | Consent of captain |  |
| ConsentPDParticipant | boolean| `bool` | ✓ | | Consent of participants |  |
| ConsentRules | boolean| `bool` | ✓ | | Consent of rules |  |
| Name | string| `string` | ✓ | | Team name |  |
| Participants | [][Participant](#participant)| `[]*Participant` | ✓ | | Participants list |  |
| UniversityID | int64 (formatted integer)| `int64` | ✓ | | University ID |  |



### <span id="participant"></span> Participant


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| Course | int64 (formatted integer)| `int64` |  | |  |  |
| CreatedAt | string| `string` |  | |  |  |
| Email | string| `string` |  | |  |  |
| FirstName | string| `string` |  | |  |  |
| ID | int64 (formatted integer)| `int64` |  | |  |  |
| IsCaptain | boolean| `bool` |  | |  |  |
| LastName | string| `string` |  | |  |  |
| MiddleName | string| `string` |  | |  |  |
| Phone | string| `string` |  | |  |  |
| ShirtSize | string| `string` |  | |  |  |
| TeamID | int64 (formatted integer)| `int64` |  | |  |  |
| Telegram | string| `string` |  | |  |  |



### <span id="team"></span> Team


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| CityID | int64 (formatted integer)| `int64` |  | |  |  |
| CityName | string| `string` |  | |  |  |
| CreatedAt | string| `string` |  | |  |  |
| ID | int64 (formatted integer)| `int64` |  | |  |  |
| Name | string| `string` |  | |  |  |
| Participants | [][Participant](#participant)| `[]*Participant` |  | |  |  |
| UniversityID | int64 (formatted integer)| `int64` |  | |  |  |
| UniversityName | string| `string` |  | |  |  |



### <span id="university"></span> University


  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| ID | int64 (formatted integer)| `int64` |  | |  |  |
| Name | string| `string` |  | |  |  |


