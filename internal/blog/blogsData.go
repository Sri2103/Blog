package blog


var AllBlogs = []*BlogExcerpt{
	{
		Title:       "Go in Action",
		Description: "Effective Go: Embrace Concurrency",
		Url:         "/goAction",
		Image:       "",
		Date:        "",
		Tags:        []string{"golang"},
	},

	{
		Title:       "The Hitchhiker's Guide to the Galaxy",
		Description: "Chapter 1: Outer Space is Not Just for Living There Anymore.",
		Url:         "/guideToGalaxy",
		Image:       "",
		Date:        "October 20, 2013",
		Tags:        []string{"books", "space"},
	},
}
