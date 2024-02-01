package blog

import (
    "testing"
    "time"
)

func TestStringToDate(t *testing.T) {
    validDate := "2006-01-02"
    date := stringToDate(validDate)
    expected := time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC)
    if date != expected {
        t.Errorf("Expected %v, got %v", expected, date)
    }

    invalidDate := "abc"
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("stringToDate should have panicked for invalid date")
        }
    }()
    _ = stringToDate(invalidDate)
}

func TestNewBlogExcerpt(t *testing.T) {
    title := "Test Title"
    desc := "Test Description" 
    url := "http://test.com"
    image := "test.jpg"

    excerpt := NewBlogExcerpt(title, desc, url, image)

    if excerpt.Title != title {
        t.Errorf("Expected title %q, got %q", title, excerpt.Title)
    }

    if excerpt.Description != desc {
        t.Errorf("Expected description %q, got %q", desc, excerpt.Description)
    }

    if excerpt.Url != url {
        t.Errorf("Expected url %q, got %q", url, excerpt.Url)
    }

    if excerpt.Image != image {
        t.Errorf("Expected image %q, got %q", image, excerpt.Image)
    }
}


func TestFilterDuplicates(t *testing.T) {
	title := "Test Title"
    desc := "Test Description" 
    url := "http://test.com"
    image := "test.jpg"


  t.Run("filters list with no duplicates", func(t *testing.T) {
	post1 := NewBlogExcerpt(title,desc,url,image)
	_ = post1.SetSlug("post-1")
	post2 := NewBlogExcerpt(title,desc,url,image)
	_ = post2.SetSlug("post-2")
    input := Blogs{
      post1, 
      post2,
    }

    expected := Blogs{
      post1,
	  post2,
    }

    output := filterDuplicates(input)

    if len(output) != len(expected) {
      t.Errorf("Expected %d items, got %d", len(expected), len(output)) 
    }

    for i := range output {
      if output[i] != expected[i] {
        t.Errorf("Expected item %v, got %v", expected[i], output[i])
      }
    }
  })

  t.Run("filters list with duplicates", func(t *testing.T) {
	post1 := NewBlogExcerpt(title,desc,url,image)
	_ = post1.SetSlug("post-1")
	post2 := NewBlogExcerpt(title,desc,url,image)
	_ = post2.SetSlug("post-2")

	post3 := NewBlogExcerpt(title,desc,url,image)
	_ = post3.SetSlug("post-1")

    input := Blogs{
       post1,
	   post2,
	   post3,
    }

    expected := Blogs{
		post1,
		post2,
    }

    output := filterDuplicates(input)

    // Assert expected length
    if len(output) != len(expected) {
      t.Errorf("Expected %d items, got %d", len(expected), len(output))
    }

    // Assert no duplicates
    seen := make(map[string]bool)
    for _, item := range output {
      if seen[item.Slug] {
        t.Errorf("Duplicate found: %v", item.Slug)
      }
      seen[item.Slug] = true
    }
  })

  t.Run("filters empty list", func(t *testing.T) {
    input := Blogs{}
    output := filterDuplicates(input)

    if len(output) != 0 {
      t.Errorf("Expected empty list, got %v", output)
    }
  })

}


func TestFilterByTag(t *testing.T) {

	t.Run("filters blogs by tag", func(t *testing.T) {
		blogs := Blogs{
			{Title: "Post 1", Tags: []string{"go"}},
			{Title: "Post 2", Tags: []string{"python"}},
		}
		expected := Blogs{
			{Title: "Post 1", Tags: []string{"go"}},
		}

		result := blogs.FilterByTag("go")

		if len(result) != len(expected) {
			t.Errorf("Expected %d blogs, got %d", len(expected), len(result))
		}

		if result[0].Title != expected[0].Title {
			t.Errorf("Expected blog title %q, got %q", expected[0].Title, result[0].Title)
		}
	})

	t.Run("returns empty list if no matches", func(t *testing.T) {
		blogs := Blogs{
			{Title: "Post 1", Tags: []string{"go"}},
		}
		// expected := Blogs{}

		result := blogs.FilterByTag("python")

		if len(result) != 0 {
			t.Errorf("Expected empty list, got %v", result)
		}
	})
  
  }


  func TestSortBlogSliceByDate(t *testing.T) {

	t.Run("sorts slice in ascending date order", func(t *testing.T) {
	  slice := []*BlogExcerpt{
		{Date: "2020-01-01"},
		{Date: "2019-01-01"}, 
	  }
	  expected := []*BlogExcerpt{
		{Date: "2019-01-01"},
		{Date: "2020-01-01"},
	  }
  
	  SortBlogSliceByDate(slice)
  
	  if !equalSlices(slice, expected) {
		t.Errorf("Expected %v, got %v", expected, slice) 
	  }
	})
  
	t.Run("works on empty slice", func(t *testing.T) {
	  var slice []*BlogExcerpt
	  SortBlogSliceByDate(slice)
	  // no need to assert anything, should not panic
	})
  
	t.Run("works on slice with one element", func(t *testing.T) {
	  slice := []*BlogExcerpt{{Date: "2019-01-01"}}
	  SortBlogSliceByDate(slice)
	  // no need to assert anything, should not panic
	})
  
  }
  
  func equalSlices(a, b []*BlogExcerpt) bool {
	if len(a) != len(b) {
	  return false
	}
	for i := range a {
	  if a[i].Date != b[i].Date {
		return false
	  }
	}
	return true
  }

