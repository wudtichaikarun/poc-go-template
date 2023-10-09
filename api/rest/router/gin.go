package router

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	*gin.Engine
}

type GinRouterGroup struct {
	*gin.RouterGroup
}

type GinContext struct {
	// composition gin.Context
	// สามารถเรียก method gin ได้หมดเลย
	*gin.Context
}

func NewGinContext(c *gin.Context) *GinContext {
	return &GinContext{Context: c}
}

func NewGinRouter() *GinRouter {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
		"TransactionID",
	}
	r.Use(cors.New(config))

	// sign := os.Getenv("SIGN")
	// r.Use(auth.Protect([]byte(sign)))

	return &GinRouter{r}
}

func (c *GinContext) Bind(v interface{}) error {
	// c.Context คล้ายๆการเรียก supper
	// c.ShouldBindJSON(v) ก็ได้เพราะ composition gin.Context ไปแล้วแต่แบบนี้ save กว่า
	return c.Context.ShouldBindJSON(v)
}

func (c *GinContext) JSON(s int, v interface{}) error {
	c.Context.JSON(s, v)
	return nil
}

func (c *GinContext) TransactionID() string {
	return c.Request.Header.Get("TransactionID")
}

func (c *GinContext) Audience() string {
	if aud, ok := c.Get("aud"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}

func (c *GinContext) GetToken() string {
	if aud, ok := c.Get("Authorization"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}

func (c *GinContext) GetParam(key string) (int, error) {
	param := c.Param(key)

	v, err := strconv.Atoi(param)
	if err != nil {
		return 0, err

	}

	return v, nil
}

/*
ส่วนนี้ไม่ค่อยจำเป็น แค่ช่วยให้ตอนเขียน

  - from
    protected.POST("/todos",
    func(c *gin.Context){
    todoHandler.NewTask(router.NewGinContext(c))
    })

  - to
    protected.POST("/todos",
    router.ConvertToGinHandler(todoHandler.NewTask)
    )
*/
func ConvertToGinHandler(handler func(Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&GinContext{Context: c})
		// handler(NewGinContext(c))
	}
}

// router r.POST()
func (r *GinRouter) POST(path string, handler func(Context)) {
	r.Engine.POST(path, ConvertToGinHandler(handler))
}

func (r *GinRouter) GET(path string, handler func(Context)) {
	r.Engine.GET(path, ConvertToGinHandler(handler))
}

func (r *GinRouter) DELETE(path string, handler func(Context)) {
	r.Engine.DELETE(path, ConvertToGinHandler(handler))
}

// router group r.Group().POST
func (rg *GinRouterGroup) POST(path string, handler func(Context)) {
	rg.RouterGroup.POST(path, ConvertToGinHandler(handler))
}

func (rg *GinRouterGroup) GET(path string, handler func(Context)) {
	rg.RouterGroup.GET(path, ConvertToGinHandler(handler))
}

func (rg *GinRouterGroup) DELETE(path string, handler func(Context)) {
	rg.RouterGroup.DELETE(path, ConvertToGinHandler(handler))
}
