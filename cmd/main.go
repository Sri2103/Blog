package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/unrolled/render"
	"github.com/yuin/goldmark"
)

type RenderWrapper struct {
    rnd *render.Render
}

func (r *RenderWrapper) Render(w io.Writer, name string, data interface{},c echo.Context) error {
    return r.rnd.HTML(w, 0, name, data)
}

func main() {
	Router := echo.New()
	r := &RenderWrapper{render.New(render.Options{
		Directory:  "../templates",
		Extensions: []string{".html"},
		IsDevelopment: true,
	})}
	Router.Renderer = r

	Router.GET("/",func(c echo.Context) error {
		//  read file
		md,err := os.ReadFile("../posts/hello-world.md")
		if err!= nil {
			return c.String(500,err.Error())
		}
		var buf bytes.Buffer
        err = goldmark.Convert(md, &buf);
		if err!= nil {
			return c.String(500,err.Error())
		}
		
		fmt.Println(buf.String())
		//  map of string and interface as data
		data := make(map[string]interface{})

		data["blog"] = template.HTML(buf.Bytes())
		
		err = c.Render(200,"base.page",data)

		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}

		return nil
	})

	Router.Logger.Fatal(Router.Start(":4500"))

}