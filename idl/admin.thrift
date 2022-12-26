namespace go admin

struct Request {
        1: string message
}

struct Response {
        1: string message
}

struct AddRequest {
  1: i64 first
  2: i64 second
}

struct AddResponse {
  1: i64 sum
}

service Hello {
    Response echo(1: Request req)
    AddResponse add(1: AddRequest req)
}