package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	//（可选）添加两个内置处理程序（handlers）
	//可以从任何与 http 相关的恐慌（http-relative panics） 中恢复
	//并将请求记录到终端。
	app.Use(recover.New())
	app.Use(logger.New())

	//  "/"  服务于一个基于根路由的控制器。
	mvc.New(app).Handle(new(ExampleController))

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	// http://localhost:8080/custom_path
	app.Run(iris.Addr(":8080"))
}

// ExampleController 服务于 "/", "/ping" and "/hello" 路由。
type ExampleController struct{}

// Get serves
// Method:   GET
// Resource: http://localhost:8080
func (c *ExampleController) Get() mvc.Result {
	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>Welcome</h1>",
	}
}

// GetPing serves
// Method:   GET
// Resource: http://localhost:8080/ping
func (c *ExampleController) GetPing() string {
	return "pong"
}

// GetHello serves
// Method:   GET
// Resource: http://localhost:8080/hello
func (c *ExampleController) GetHello() interface{} {
	return map[string]string{"message": "Hello Iris!"}
}

// BeforeActivation 会被调用一次，在控制器适应主应用程序之前
// 并且当然也是在服务运行之前
// 在版本 9 之后，你还可以为特定控制器的方法添加自定义路由。
// 在这个控制器，你可以注册自定义方法的处理程序。
// 使用带有 `ca.Router` 的标准路由
// 在不适用 mvc 的情况下，你可以做任何你可以做的事情
// 并将添加的依赖项绑定到
// 一个控制器的字段或者方法函数的输入参数中。
func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
	anyMiddlewareHere := func(ctx iris.Context) {
		ctx.Application().Logger().Warnf("Inside /custom_path")
		ctx.Next()
	}

	b.Handle(
		"GET",
		"/custom_path",
		"CustomHandlerWithoutFollowingTheNamingGuide",
		anyMiddlewareHere,
	)

	// 或者甚至可以添加基于这个控制器路由的全局中间件，
	// 比如在这个例子里面的根路由 "/" :
	// b.Router().Use(myMiddleware)
}

// CustomHandlerWithoutFollowingTheNamingGuide serves
// Method:   GET
// Resource: http://localhost:8080/custom_path
func (c *ExampleController) CustomHandlerWithoutFollowingTheNamingGuide() string {
	return "hello from the custom handler without following the naming guide"
}

// GetUserBy serves
// Method:   GET
// Resource: http://localhost:8080/user/{username:string}
// By 是一个保留的“关键字”来告诉框架你要
// 在函数的输入参数中绑定路径参数，它也是
// 有助于在同一个控制器中拥有 "Get“ 和 "GetBy"。
//
// func (c *ExampleController) GetUserBy(username string) mvc.Result {
//     return mvc.View{
//         Name: "user/username.html",
//         Data: username,
//     }
// }

/*
在一个控制器中可以使用多个 HTTP 方法，工厂会
为每条路由注册正确的 HTTP 方法，你可以按需取消注释：

func (c *ExampleController) Post() {}
func (c *ExampleController) Put() {}
func (c *ExampleController) Delete() {}
func (c *ExampleController) Connect() {}
func (c *ExampleController) Head() {}
func (c *ExampleController) Patch() {}
func (c *ExampleController) Options() {}
func (c *ExampleController) Trace() {}
*/

/*
func (c *ExampleController) All() {}
//        或者
func (c *ExampleController) Any() {}



func (c *ExampleController) BeforeActivation(b mvc.BeforeActivation) {
    // 1 -> HTTP 方法
    // 2 -> 路由路径
    // 3 -> 控制器的方法名称应该是该路由的处理程序（handlers）。
    b.Handle("GET", "/mypath/{param}", "DoIt", optionalMiddlewareHere...)
}

//激活后，所有依赖项都被设置 - 因此只能访问它们
//但仍可以添加自定义控制器或者简单的标准处理程序（handlers）。
func (c *ExampleController) AfterActivation(a mvc.AfterActivation) {}

*/