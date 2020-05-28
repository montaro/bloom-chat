# Bloom Chat Protocol

# Bloom Chat Protocol

# Features

- ???
- ???
- ???

# Definitions

Client

User

Session

Room

...

# Protocol

The bloom-chat protocol is built on top of WebSockets as a transport (for push/pull mixed-model) interaction between server and client.

Generally, The server is expecting requests messages from clients encoded as JSON strings including a locally generated request ID(preferably UUID).

The server replies to the clients’ requests with JSON responses including the correlated request ID.

## Client Request

```json
{
    "Request_id": "UUID",
     "op": "op_name",
     "Data":{
           "key":"value",
     }
}
```

## Server Response

```json
{
    "Request_id": "UUID",
     "Data":{
          "key":"value",
     }
}
```

## Handshake

Server is expecting the client to send the first message which includes the client supported protocol version.

```json
{
    "request_id": "e2992f2f-1ccc-44d9-ac55-d486e52248d0",
    "op": "INITIALIZE",
    "data": {
      "protocolversion": 1.0,
			"email": "ahmed@example.com",
			"username": "montaro"
    }
 }
```

The server will reply with the latest supported protocol version in the server in either one of these responses:

### Response Welcome! with client ID

```json
{
		"request_id":"e2992f2f-1ccc-44d9-ac55-d486e52248d0",
		"data":{
			"client_id":"2a0ad008-4442-444a-8a34-6659883cd566"
	}
}
```

### Response Error if unsupported protocol version

```json
{
	"request_id":"e2992f2f-1ccc-44d9-ac55-d486e52248d0",
	"data":{
		"error":"protocol handshake error: unsupported protocol version, supported version=1.0"
	}
}
```

### Response Error if the connection was already initialized

```json
{
	"request_id":"e2992f2f-1ccc-44d9-ac55-d486e52248d0",
	"data":{
		"error":"unexpected op: INITIALIZE"
	}
}
```

## Room Operations Requests

## Create room

```json
{
    "request_id": "e2992f2f-1ccc-44d9-ac55-d486e52248d0",
    "op": "CREATE_ROOM",
    "data": {
        "topic": "kewl stuff"
    }
}
```

### Response

```json
{
	"roomId":1
}
```

### Set room topic

```json
{
    "request_id": "e2992f2f-1ccc-44d9-ac55-d486e52248d0",
    "op": "SET_ROOM_TOPIC",
    "data": {
        "roomId": 1,
         "topic": "kewl stuff"
    }
}
```

### Response

```json
{
	"request_id":"e2992f2f-1ccc-44d9-ac55-d486e52248d0",
	"data":{
	    "done":true
	}
}
```

### Join room

```json
{
    "request_id": "e2992f2f-1ccc-44d9-ac55-d486e52248d0",
    "op": "JOIN_ROOM",
    "data": {
        "roomId": 1
    }
}
```

### Response

```json
{
	"request_id":"e2992f2f-1ccc-44d9-ac55-d486e52248d0",
	"data":{
	    "done":true
	}
}
```

### List Rooms

```json
{
	"request_id": "e2992f2f-1ccc-44d9-ac55-d486e52248d0",
  "op": "LIST_ROOMS",
	  "data": {
	}
}
```

### Response

```json
{
	"request_id":"e2992f2f-1ccc-44d9-ac55-d486e52248d0",
	"data":{
		"rooms":[
			{"id":9,"topic":"kewl stuff"},
			{"id":10,"topic":"good guys"},
			{"id":8,"topic":"shitty chat"}
		]
	}
}
```

### Retrieve Room Messages //TODO

```json
{
	"request_id": "e2992f2f-1ccc-44d9-ac55-d486e52248d0",
  "op": "Ret_ROOM_MSGS",
	  "data": {
		//TODO
	}
}
```

## User Operations Requests

### Set user name

```json
{
    "request_id": "e2992f2f-1ccc-44d9-ac55-d486e52248d0",
    "op": "SET_USER_NAME",
    "data": {
        "name": "kewl name"
    }
}
```

### Response

```json
{
    "request_id":"e2992f2f-1ccc-44d9-ac55-d486e52248d0",
    "data":{
	"done":true
	}
}
```

## Message Operations

### Message Structure

Text Message

```json
{
    "request_id": "bb6fcb9c-7c13-40e7-a4fa-4db73b60aaf6",
    "op": "SEND_MSG",
    "data": {
        "room_id": 1,
        "message": {
            "id": "234324",
            "kind": "TEXT",
            "content": "I'm Ahmed @ihab 😂",
            "formatted_content": "I'm _Ahmed_ [userID12](@ihab) 😂",
            "timestamp": 13213213,
            "status": "SEEN", // optional only required in case of response
            "sender": {
                "id": "213kj213123",
                "name": "Ahmed",
                "status": "ONLINE",
                "client": "VIA_WEB"
            },
            "reply_to": {
                "id": "223"
            },
            "permissions": [
                "CAN_DELETE"
            ]
        }
    }
}
```

Photo Message

```json
{
    "request_id": "bb6fcb9c-7c13-40e7-a4fa-4db73b60aaf6",
    "op": "SEND_MSG_PHOTO",
    "data": {
        "room_id": 1,
        "message": {
            "id": "234324",
            "kind": "PHOTO",
            "sizes": [
                {
                    "thumbnail": {
                        "url": "",
                        "width": "",
                        "height": ""
                    }
                },
                {
                    "large": {
                        "url": "",
                        "width": "",
                        "height": ""
                    }
                }
            ],
            "content": "I'm Ahmed @ihab 😂",
            "formatted_content": "I'm _Ahmed_ [userID12](@ihab) 😂",
            "timestamp": 13213213,
            "status": "SEEN", // optional only required in case of response
            "sender": {
                "id": "213kj213123",
                "name": "Ahmed",
                "status": "ONLINE",
                "client": "VIA_WEB"
            },
            "reply_to": {
                "id": "223"
            },
            "permissions": [
                "CAN_DELETE",
                "CAN_DOWNLOAD"
            ]
        }
    }
}
```

Animated Photo

```json
{
    "request_id": "bb6fcb9c-7c13-40e7-a4fa-4db73b60aaf6",
    "op": "SEND_MSG_PHOTO_ANIM",
    "data": {
        "room_id": 1,
        "message": {
            "id": "234324",
            "kind": "ANIMATED_PHOTO",
            "sizes": [
                {
                    "thumbnail": {
                        "url": "",
                        "width": "",
                        "height": ""
                    }
                },
                {
                    "large": {
                        "url": "",
                        "width": "",
                        "height": ""
                    }
                },
                {
                    "video": {
                        "url": "",
                        "width": "",
                        "height": ""
                    }
                },
            ],
            "content": "I'm Ahmed @ihab 😂",
            "formatted_content": "I'm _Ahmed_ [userID12](@ihab) 😂",
            "timestamp": 13213213,
            "status": "SEEN", // optional only required in case of response
            "sender": {
                "id": "213kj213123",
                "name": "Ahmed",
                "status": "ONLINE",
                "client": "VIA_WEB"
            },
            "reply_to": {
                "id": "223"
            },
            "permissions": [
                "CAN_DELETE",
                "CAN_DOWNLOAD"
            ]
        }
    }
}
```