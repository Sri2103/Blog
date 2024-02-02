package blog


var AllBlogs = []*BlogExcerpt{
	{
		Title: "My test markdown file",
		Description: "This is a test Markdown file. Below are some examples of Markdown syntax",
		Url:     "/blog/hello-world",
		Image: "",
		Date:    "2022-01-29",
		Tags:    []string{"test","markdown"},
	},
	{
		Title: "How to update kubernetes cluster",
		Description: `Kubernetes has been updated! Here's how you can upgrade your existing version`,
		Url:         "/blog/updatek8sVersion",
		Date:        "2022-01-25",
		Tags:        []string{"howto","kubernetes"},	
	},
	{
		Title: "How to debug the crashed apiserver",
		Description: `When Kubernetes crashes, it becomes hard to understand what went wrong. This guide will help you in debugging the issue and fix it if possible.`,
		Url: "/blog/debuggingCluster",
		Date: "2024-02-02",
		Tags: []string{"cka","kubernetes"},
	},

}
