package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/Sri2103/blog/internal/blog"
	"github.com/labstack/echo/v4"
	"github.com/unrolled/render"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"go.abhg.dev/goldmark/frontmatter"
)

type RenderWrapper struct {
	rnd *render.Render
}

func (r *RenderWrapper) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.rnd.HTML(w, 0, name, data)
}

//go:embed all:templates
var temps embed.FS

//go:embed all:posts
var postings embed.FS

//go:embed all:static
var staticFiles embed.FS

func main() {
	Router := echo.New()
	r := &RenderWrapper{render.New(render.Options{
		Directory:     "templates",
		Extensions:    []string{".html"},
		IsDevelopment: true,
		FileSystem: &render.EmbedFileSystem{
			FS: temps,
		},
	})}
	Router.Renderer = r

	gd := goldmark.New(
		goldmark.WithExtensions(extension.NewTypographer(), extension.GFM, &frontmatter.Extender{}),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),

			html.WithUnsafe(),
			html.WithXHTML(),
		))

	Router.GET("/", func(c echo.Context) error {
		data := make(map[string]interface{})
		fileNames, err := ReadFileNames()
		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}
		data["posts"] = fileNames
		data["blogsExcerpts"] = blog.AllBlogs
		err = c.Render(200, "home", data)

		if err != nil {
			return echo.NewHTTPError(500, err.Error())
		}

		return nil
	})

	Router.GET("/blog/:slug", func(c echo.Context) error {
		title := c.Param("slug")
		// read blog file
		md, err := postings.ReadFile("posts/" + title + ".md")
		if err != nil {
			return c.String(500, err.Error())
		}
		ctx := parser.NewContext()
		var buf bytes.Buffer
		err = gd.Convert(md, &buf, parser.WithContext(ctx))
		if err != nil {
			return c.String(500, err.Error())
		}
		tempData := frontmatter.Get(ctx)

		var tdata struct {
			Title string   `yaml:"title"`
			Tags  []string `yaml:"tags"`
		}
		if err := tempData.Decode(&tdata); err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("Title: ", tdata.Title)
		fmt.Println("Tags: ", tdata.Tags)

		data := make(map[string]interface{})
		data["blog"] = template.HTML(buf.Bytes())
		

		return c.Render(200, "base.page", data)

	})

	Router.StaticFS("/static/", echo.MustSubFS(staticFiles, "static"))

	Router.Logger.Fatal(Router.Start(":4500"))

}

func ReadFileNames() ([]string, error) {
	var filenames []string
	files, err := postings.ReadDir("posts")
	if err != nil {
		return filenames, err
	}
	for _, f := range files {
		filenames = append(filenames, strings.Split(f.Name(), ".")[0])
	}
	return filenames, nil
}
