package blog


var AllBlogs = []*BlogExcerpt{
	{
		Title: "My test markdown file",
		Description: "This is a test Markdown file. Below are some examples of Markdown syntax",
		Url:     "/blog/hello-world",
		Image: "",
		Date:    "January 15, 2014",
		Tags:    []string{"test","markdown"},
	},
	{
		Title: "How to update kubernetes cluster",
		Description: `Kubernetes has been updated! Here's how you can upgrade your existing version`,
		Url:         "/blog/updatek8sVersion",
		Date:        "May 31, 2019",
		Tags:        []string{"howto","kubernetes"},	
	},

}
