package blog


var AllBlogs = []*BlogExcerpt{
	{
		Title: "My test markdown file",
		Description: "This is a test Markdown file. Below are some examples of Markdown syntax",
		Url:     "/blog/hello-world",
		Image: "",
		Date:    "2014-01-14",
		Tags:    []string{"test","markdown"},
	},
	{
		Title: "How to update kubernetes cluster",
		Description: `Kubernetes has been updated! Here's how you can upgrade your existing version`,
		Url:         "/blog/updatek8sVersion",
		Date:        "2019-05-19",
		Tags:        []string{"howto","kubernetes"},	
	},

}
