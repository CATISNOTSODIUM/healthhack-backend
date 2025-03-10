# API Guide

- [API Guide](#api-guide)
  - [HTTP codes](#http-codes)
  - [Endpoints](#endpoints)
    - [User endpoints `/api/users`](#user-endpoints-apiusers)
      - [Get user](#get-user)
      - [Create user](#create-user)
      - [Update user](#update-user)
      - [Delete user](#delete-user)
    - [History endpoint `/api/history`](#history-endpoint-apihistory)
      - [Create empty history](#create-empty-history)
      - [Retrieve latest histories from userID (Including analyzed data)](#retrieve-latest-histories-from-userid-including-analyzed-data)
    - [Voice analysis endpoint `/internal/voice-analysis`](#voice-analysis-endpoint-internalvoice-analysis)
  - [Add/Update new voice analysis](#addupdate-new-voice-analysis)

This document provides detailed information about the available API endpoints, including request methods, parameters, responses, and error codes.

At this stage, **I have not set up any forms of authentication** (since I'm quite busy lately). Hence, this documentation is subjected to change.
## HTTP codes
* **StatusBadRequest (400)** The request was malformed. The request does not follow the schema guidelined by this documentation.
* **StatusInternalServerError (500)** The request is unsuccessfully processed due to internal server error.
## Endpoints

### User endpoints `/api/users`
#### Get user
* **Method**: GET
* **URL** `/api/users/get` 
* **Description** Get basic information for a user with ID `USER_ID`
* **Example body**
```json
{
  "user_id": "USER_UUID",
}
```
* **Example Response (If the user is founded)** 
```json
{
  "username": "username",
  "age": 10,
  "medical_record": "A single child with personal trauma"
}
```
* **Example Response (not founded)** 
```
record not found
User XXX not founded
```
#### Create user
* **Method**: POST
* **URL** `/api/users/create` 
* **Description** Create user from username and google ID.
* **Example body** Note that **username** and **google_id** fields are required. Other fields are optional.
```json
{
  "username": "username",
  "google_id": "xxxxxxx",
  "age": 10,
  "medical_record": "A single child with personal trauma"
}
```
* **Example Response** 
```json
{
  "id": "USER_UUID",
  "username": "username",
  "age": 10,
  "medical_record": "A single child with personal trauma"
}
```
#### Update user
* **Method**: PUT
* **URL** `/api/users/update` 
* **Description** Update user information based on user id.
* **Example body** Note that the `id` field is required to specify the target. Any other fields provided in the body will be updated accordingly.
```json
{
  "id": "USER_UUID",
  "username": "username",
  "age": 10,
  "medical_record": "A single child with personal trauma"
}
```
* **Example Response** 
```json
{
  "username": "username",
  "age": 10,
  "medical_record": "A single child with personal trauma"
}
```
#### Delete user
* **Method**: PUT
* **URL** `/api/users/delete` 
* **Description** Delete user information based on user id.
* **Example body** 
```json
{
  "id": "USER_UUID"
}
```
* **Example Response** 
```json
"Successfully delete user id: USER_ID"
```
### History endpoint `/api/history`
#### Create empty history
**Add new voice analysis**
* **Method**: POST
* **URL** `/api/history/create`
* **Description** Create new empty history.
*  **Example body** 
```json
{
  "user_id": "USER_UUID"
}
```
*  **Example response** The newly created history UUID.
```json
{
  "id": "HISTORY_UUID"
}
```
#### Retrieve latest histories from userID (Including analyzed data)
* **Method**: GET
* **URL** `/api/history/get`
* **Description** Retrieve user particular histories. You can specify the number of histories you want to retrieve.
* **Example body**
```json
{
  "user_id": "USER_UUID",
  "number_of_histories": 1  
}
```
* **Example response (some fields are truncated)**
```json
[{
  "id":"1f020d38-6f1b-465c-b476-a31ae153b469",
  "user_id":"45d7c3cc-2a10-4de8-bb3d-ec8be81164e3",
  "voice_activity_analysis":{},
  "text_analysis":{},
}]
```
### Voice analysis endpoint `/internal/voice-analysis`
## Add/Update new voice analysis
* **Method**: PUT
* **URL** `/internal/voice-analysis/create`
* **Description** Add voice analysis data from specified history ID.
*  **Example body** 
```JSON
{
  "history_id": "HISTORY_UUID", 
  "total_duration": 120.5,
  "total_speech_duration": 85.2,
  "total_pause_duration": 35.3,
  "num_speech_segments": 12,
  "num_pauses": 11,
  "answer_delay_duration": 1.8,
  "speech_segments": [
    {
      "start_time": 0.5,
      "end_time": 10.2,
      "duration": 9.7
    }
  ],
  "pause_segments": [
    {
      "start_time": 10.2,
      "end_time": 15.5,
      "duration": 5.3
    }
  ]
}
```
* **Example message (If success)**
```
Successfully update history 1f020e78-1f1b-965c-b476-a31ae153b469 with voice activity analysis 06b7e111-0cb9-46cf-97f1-ac7c6bb4f70e
```