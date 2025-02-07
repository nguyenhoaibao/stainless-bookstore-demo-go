# bookstore2

Methods:

- <code title="post /echo">client.<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go#Bookstore2Service.Echo">Echo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go#EchoParams">EchoParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go#User">User</a>

Methods:

- <code title="get /users/{username}">client.Users.<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go#UserGetParams">UserGetParams</a>) (<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/{username}">client.Users.<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go">bookstore2</a>.<a href="https://pkg.go.dev/github.com/nguyenhoaibao/stainless-bookstore-demo-go#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
