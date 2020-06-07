# Bloom Chat Protocol V2.0

# Features

- Client can join or start a room
- Set display name
- Set room topic
- Send a message to room

# Definitions

### Client

Any program that uses the services offered by the Bloom chat server (Web, CLI, etc..)

### User

Any human that uses the Bloom chat service via a client

### Session

Established communication between a client and Bloom chat server

### Room

Stream of message channel that are sent by and delivered to joined members

# JSON Protocol

The bloom-chat protocol is built on top of **WebSockets** as a transport (for push/pull mixed-model) interaction between server and client.

Generally, The server is expecting requests messages from clients encoded as **JSON** strings including a locally generated request ID.

The server replies to the clients’ requests with JSON responses including the correlated request ID.

## Client Request

```json
{
  "request_id": "RequestID",
  "op": "op_name",
  "data": {
    "key": "value"
  }
}
```

## Server Response

```json
{
  "request_id": "RequestID",
  "data": {
    "key": "value"
  }
}
```

## Authentication

### TODO

## Handshake

### Initialize

Server is expecting the client to send the first message which optionally includes the session ID and the display name

```json
{
  "request_id": "e2992f2f",
  "op": "INITIALIZE",
  "data": {
    "session_id": "e2992f2f",
    "display_name": "montaro"
  }
}
```

The server will reply with the same session ID if it still exists in the server data, otherwise create a new session and reply with it's ID

### Response Welcome! with session ID

```json
{
  "request_id": "e2992f2f",
  "data": {
    "session_id": "2a0ad008"
  }
}
```

## Room Operations

### Join room

Join or create a new room, if it doesn't exist in the server data

```json
{
  "request_id": "e2992f2f",
  "op": "JOIN_ROOM",
  "data": {
    "roomId": "e2992f2f"
    }
}
```

### Response

### Set room topic

```json
{
  "request_id": "e2992f2f",
  "op": "SET_ROOM_TOPIC",
  "data": {
    "roomId": "e2992f2f",
    "topic": "kewl stuff"
  }
}
```

### Response

```json
{
  "request_id": "e2992f2f",
  "data": {
    "roomId": "e2992f2f",
    "topic": "kewl stuff"
  }
}
```

### Get Room Metadata

```json
{
  "request_id": "e2992f2f",
  "op": "GET_ROOM",
  "data": {
     "roomId": "e2992f2f"
    }
}
```

### Response

```json
{
  "request_id": "e2992f2f",
  "data": {
    "roomId": "e2992f2f",
    "topic": "Kewl Stuff!",
    "members": [
      {
        "id": "rtyu567",
        "username": "rafaello",
        "visual": {
          "display-name": "Ahmed ElRefaey",
          "color": "RGB",
          "photo": "http://www.example.com/photos/avatar1.png"
        },
        "presence": {
          "status": "Away",
          "last_active": "TS"
        }
      },
      {
        "id": "rtyu568",
        "username": "soli",
        "visual": {
          "display-name": "Ahmed Soliman",
          "color": "RGB",
          "photo": "http://www.example.com/photos/avatar2.png"
        },
        "presence": {
          "status": "Online",
          "last_active": "TS"
        }
      }
    ]
  }
}
```

### Fetch Room Messages

Fetch messages before the specified `before_seq` up to a limit

If the server messages are less than the limit or equals zero, this means the server doesn't have any more messages, the client should stop asking for older messages.

```json
{
  "request_id": "e2992f2f",
  "op": "FETCH_ROOM_MSGS",
  "data": {
    "roomId": "e2992f2f",
    "before_seq": 123,
    "limit": 10
  }
}
```

### Response

```json
{
  "request_id": "e2992f2f",
  "data": {
    "roomId": "e2992f2f",
    "messages": [
      {
        "seq_nr": 122,
        "body": {
          "kind": "TEXT",
          "content": "I'm Ahmed @ihab 😂",
          "formatted_content": "I'm _Ahmed_ [userID12](@ihab) 😂"
        },
        "sender": "e2992f2f",
        "client_time": "TS",
        "created_at": "TS"
      },
      {
        "seq_nr": 121,
        "body": {
          "kind": "TEXT",
          "content": "I'm Soli",
          "formatted_content": "I'm Soli"
        },
        "sender": "e2992f2d",
        "client_time": "TS",
        "created_at": "TS"
      }
    ]
  }
}
```

### Subscribe

Subscribe to the messages stream of a specific room. The client should specify in the `FetchQuery` either the `Since` or the `Limit` value.

If `Since` is specified, this means that the client is interested in messages newer than a `SeqNr` and from then on.

If `Limit` is specified, this means that the client is interested only in the latest `X` messages and from now on.

```json
{
  "request_id": "e2992f2f",
  "op": "SUBSCRIBE",
  "data": {
    "roomId": "e2992f2f",
    "fetch_query": {
      "since": 123,
      "limit": 10
    }
  }
}
```

### Response

If the `since` and `limit` are both specified or not of them is specified, client will receive a Bad Request response

If since or limit is specified, client will receive the queried messages and the newer messages will be pushed to the client

```json
{
  "request_id": "e2992f2f",
  "data": {
    "roomId": "e2992f2f",
    "messages": [
      {
        "seq_nr": 122,
        "body": {
          "kind": "TEXT",
          "content": "I'm Ahmed @ihab 😂",
          "formatted_content": "I'm _Ahmed_ [userID12](@ihab) 😂"
        },
        "sender": "e2992f2f",
        "client_time": "TS",
        "created_at": "TS"
      },
      {
        "seq_nr": 121,
        "body": {
          "kind": "TEXT",
          "content": "I'm Soli",
          "formatted_content": "I'm Soli"
        },
        "sender": "e2992f2d",
        "client_time": "TS",
        "created_at": "TS"
      }
    ]
  }
}
```

## User Operations Requests

### Set user name

```json
{
  "request_id": "e2992f2f",
  "op": "SET_USER_DISPLAY_NAME",
  "data": {
    "name": "Mr. Kewl"
    }
}
```

### Response

```json
{
  "request_id": "e2992f2f",
  "data": {
    "name": "Mr. Kewl"
  }
}
```

## Message Operations

### Send Message

**Text Message**

```json
{
  "request_id": "bb6fcb9c",
  "op": "SEND_MSG_TXT",
  "data": {
    "room_id": "bb6fcb9c",
    "message": {
      "seq_nr": 1234,
      "body": {
        "kind": "TEXT",
        "content": "I'm Ahmed @ihab 😂",
        "formatted_content": "I'm _Ahmed_ [userID12](@ihab) 😂"
      },
      "client_time": "TS"
    }
  }
}
```

**Photo Message (TODO)**

```json
{
  "request_id": "bb6fcb9c",
  "op": "SEND_MSG_PHOTO",
  "data": {
    "room_id": "bb6fcb9c",
    "message": {
      "seq_nr": 1234,
      "body": {
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
        "formatted_content": "I'm _Ahmed_ [userID12](@ihab) 😂"
      },
      "client_time": "TS"
    }
  }
}
```

**Animated Photo (TODO)**

```json
{
  "request_id": "bb6fcb9c",
  "op": "SEND_MSG_PHOTO",
  "data": {
    "room_id": "bb6fcb9c",
    "message": {
      "seq_nr": 1234,
      "body": {
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
            },
            "video": {
              "url": "",
              "width": "",
              "height": ""
            }
          }
        ],
        "content": "I'm Ahmed @ihab 😂",
        "formatted_content": "I'm _Ahmed_ [userID12](@ihab) 😂"
      },
      "client_time": "TS"
    }
  }
}
```