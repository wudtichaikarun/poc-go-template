package adapter

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/wudtichaikarun/poc-go-template/common"
)

type FiberRouter struct {
	*fiber.App
}

type FiberContext struct {
	*fiber.Ctx
}

func NewFiberContext(c *fiber.Ctx) *FiberContext {
	return &FiberContext{Ctx: c}
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New()

	r.Use(cors.New())
	r.Use(logger.New())

	return &FiberRouter{r}
}

func (c *FiberContext) Bind(v interface{}) {
	c.Ctx.BodyParser(v)
}

func (c *FiberContext) JSON(s int, v interface{}) error {
	return c.Ctx.Status(s).JSON(v)
}

func (c *FiberContext) TransactionID() string {
	return string(c.Ctx.Request().Header.Peek("TransactionID"))
}

func (c *FiberContext) Audience() string {
	return c.Ctx.Get("aud")
}

func (c *FiberContext) GetToken() string {
	s := c.Ctx.Get("Authorization")
	tokenString := strings.TrimPrefix(s, "Bearer ")
	return tokenString
}

func (c *FiberContext) GetParam(key string) (int, error) {
	param := c.Params(key)
	v, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func (r *FiberRouter) POST(path string, handler func(common.Ctx)) {
	r.App.Post(path, func(c *fiber.Ctx) error {
		// handler(NewFiberContext(c))
		return nil
	})
}
