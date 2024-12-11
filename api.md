# bookstore2

Methods:

- <code title="post /echo">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go#Bookstore2Service.Echo">Echo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go#EchoParams">EchoParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go#User">User</a>

Methods:

- <code title="get /users/{username}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go#UserGetParams">UserGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/{username}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/bookstore-2-go#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
