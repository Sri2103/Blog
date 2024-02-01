package blog

import (
	"encoding/json"
	"sort"
	"strings"
	"time"
)

type BlogExcerpt struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Image       string `json:"image,omitempty"`
	Date        string `json:"date,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Slug         string   `json:"slug,omitempty"` // Used for generating canonical URLs.
}

//string to Date conversion method
func stringToDate(dateString string) (time.Time) {
	// Assuming the date string is in the format "2006-01-02"
	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		panic(err)
	}
	return date
}


func NewBlogExcerpt(title, description, url, image string) *BlogExcerpt {
	return &BlogExcerpt{
		Title:       title,
		Description: description,
		Url:         url,
		Image:       image,
	}
}



func (b *BlogExcerpt) SetDatetime(dt time.Time) *BlogExcerpt {
	b.Date = dt.Format("2006-01-02T15:04:05Z") // ISO 8601 format as per RFC 3339
	return b
}

func (b *BlogExcerpt) SetDate(dt string) *BlogExcerpt {
	if len(dt) == 0 {
		return b
	}
	t, err := time.ParseInLocation("2006-01-02T15:04:05Z", dt, time.UTC)
	if err != nil {
		panic(err)
	}
	b.Date = t.Local().Format("2006-01-02")
	return b
}

func (b *BlogExcerpt)SetSlug(s string)*BlogExcerpt{
	b.Slug=strings.TrimSpace(s)
	return b
}
// A Post is a full blog post that includes all the information from an Excerpt plus  additional fields.
type Post struct {
	*BlogExcerpt

	Content   string `json:"content"`
	Author    string `json:"author,omitempty"`
	AuthorUrl string `json:"author_url,omitempty"`
	AuthorImg string `json:"author_img,omitempty"`
}

func NewPost(title, description, url, image, content string) *Post {
	return &Post{
		BlogExcerpt: NewBlogExcerpt(title, description,url,image),
		Content:     content,
	}
}

func (p *Post) SetAuthor(author, authorUrl, authorImg string) *Post {
	p.Author = author
	p.AuthorUrl = authorUrl
	p.AuthorImg = authorImg
	return p
}

func (p *Post) SetDatetime(dt  time.Time) *Post {
	p.SetDate(dt.Format("2006-01-02T15:04:05Z")) // use UTC for datetime as it's displayed in local timezone on frontend
	return p
}

func (p *Post) ToJSON() []byte {
	b,  _ := json.MarshalIndent(p, "", " ")
	return b
}

// Blogs is a slice of Blog pointers.
type Blogs []*BlogExcerpt



func (blogs Blogs) FilterByTag(tag string) Blogs {
	var filtered Blogs
	for _, blog := range blogs {
		if blog.Tags != nil {
			for _, t := range blog.Tags{
				if tag == t {
					filtered = append(filtered, blog)
					break; // Skip to the next blog if we've found
			}
			}
	}
	
}
	return filterDuplicates(filtered)
}

func filterDuplicates(input Blogs) Blogs {
	var output Blogs
	for _, blog := range input {
		found := false
		for _, existing := range output {
			if blog.Slug == existing.Slug {
				found=true
				break
			}
		}
		if !found {
			output = append(output, blog)
		}
	}
	return output
}
// SortBlogsByDate takes a slice of Blogs and sorts them by date in descending order. It returns a new sorted slice without modifying the original one.

// Sorting function for use with sort pkg.
func SortBlogSliceByDate(blogSlice []*BlogExcerpt) {
	sort.Sort(ByDate(blogSlice))
}

func SortBlogSliceByLatestDate(blogSlice []*BlogExcerpt){
	sort.Sort(sort.Reverse(ByDate(blogSlice)))
}

type ByDate []*BlogExcerpt

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return stringToDate(a[i].Date).Before(stringToDate(a[j].Date))}



