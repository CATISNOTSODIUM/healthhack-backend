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
      - [Retrieve history](#retrieve-history)
      - [Create history](#create-history)
    - [Voice analysis endpoint `/internal/voice-analysis`](#voice-analysis-endpoint-internalvoice-analysis)

This document provides detailed information about the available API endpoints, including request methods, parameters, responses, and error codes.

At this stage, **I have not set up any forms of authentication** (since I'm quite busy lately). Hence, this documentation is subjected to change.
## HTTP codes
* **StatusBadRequest (400)** The request was malformed. The request does not follow the schema guidelined by this documentation.
* **StatusInternalServerError (500)** The request is unsuccessfully processed due to internal server error.
## Endpoints

### User endpoints `/api/users`
#### Get user
* **Method**: GET
* **URL** `/api/users/get<id:USER_ID>` 
* **Description** Get basic information for a user with ID `USER_ID`
* **Example Response** 
```json
{
  "username": "username",
  "age": 10,
  "medical_record": "A single child with personal trauma"
}
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
  "id": USER_UUID,
  "username": "username",
  "age": 10,
  "medical_record": "A single child with personal trauma"
}
```
#### Update user
* **Method**: POST
* **URL** `/api/users/update` 
* **Description** Update user information based on user id.
* **Example body** Note that the `id` field is required to specify the target. Any other fields provided in the body will be updated accordingly.
```json
{
  "id": USER_UUID,
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
* **Method**: POST
* **URL** `/api/users/delete` 
* **Description** Delete user information based on user id.
* **Example body** 
```json
{
  "id": USER_UUID
}
```
* **Example Response** 
```json
"Successfully delete user id: USER_ID"
```
### History endpoint `/api/history`
#### Retrieve history
To be updated
#### Create history
To be updated

### Voice analysis endpoint `/internal/voice-analysis`
**Add new voice analysis**
* **Method**: POST
* **URL** `/internal/voice-analysis/create`
* **Description** Add voice analysis data from specified history ID.
*  **Example body** 
```JSON
{
  "history_id": HISTORY_UUID, 
  // insert data
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
    // Additional segments...
  ],
  "pause_segments": [
    {
      "start_time": 10.2,
      "end_time": 15.5,
      "duration": 5.3
    }
    // Additional segments...
  ]
}
```