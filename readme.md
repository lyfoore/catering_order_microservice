# Description
Microservice for processing CRUD operations with PostgreSQL DB.

# Deployment
`docker-compose up --build`

# API
## Rest:
`:8080/orders/{id}` - where {id} is `id` in the database (only for GET, PATCH, DELETE).
### Standard request & response:
#### GET:
Request: Without JSON, just call `:8080/orders/{id}`
Response:
```json
{
  "ID": 1,
  "UserID": 13,
  "Items": [13,15,7],
  "Status": "processing"
}
```
#### POST:
Request to `:8080/orders`:
```json
{
  "user_id" : 13,
  "items" : [13,15,7]
}
```
Response:
```json
{
  "id": 3,
  "user_id": 13,
  "items": [13,15,7],
  "status": "created"
}
```
#### PATCH:
Request:
(You do not have to submit all fields if you only want to edit one field)
```json
{
  "status" : "cooked",
  "items" : [13,15,4]
}
```
Response:
```json
{
  "id": 2,
  "user_id": 13,
  "items": [13,15,4],
  "status": "cooked"
}
```
#### DELETE:
Request: Without JSON, just call `:8080/orders/{id}`
Response:
```json
{
  "message": "order deleted"
}
```

## gRPC
gRPC Server runs at `9090` port
```protobuf
syntax = "proto3";  
  
service OrderService {  
  rpc CreateOrder(CreateOrderRequest) returns (OrderResponse);  
  rpc GetOrder(GetOrderRequest) returns (OrderResponse);  
  rpc UpdateOrder(UpdateOrderRequest) returns (OrderResponse);  
  rpc DeleteOrder(DeleteOrderRequest) returns (DeleteOrderResponse);  
}  
  
message Order {  
  uint64 id = 1;  
  uint64 user_id = 2;  
  repeated int64 items = 3;  
  string status = 4;  
}  
  
message CreateOrderRequest {  
  uint64 user_id = 1;  
  repeated int64 items = 2;  
}  
  
message GetOrderRequest {  
  uint64 id = 1;  
}  
  
message UpdateOrderRequest {  
  uint64 id = 1;  
  uint64 user_id = 2;  
  repeated int64 items = 3;  
  string status = 4;  
}  
  
message DeleteOrderRequest {  
  uint64 id = 1;  
}  
  
message OrderResponse {  
  uint64 id = 1;  
  uint64 user_id = 2;  
  repeated int64 items = 3;  
  string status = 4;  
}  
  
message DeleteOrderResponse {  
  bool deleted = 1;  
}
```
