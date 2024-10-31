# gRPC service "Auth"

HOST = ```localhost:44044```\
PROTO_File = https://github.com/DeMarDeXis/AuthProto

## Methods
1. GetToken\
_Message:_
```
{
    "user_id": "8"
}
```

_Response:_
```
{
"token": <token>
}
```