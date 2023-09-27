package main

func main() {
	csharpBlog := &Blog{
		programmingLanguage: "C#",
	}
	javaBlog := &Blog{
		programmingLanguage: "Java",
	}
	goLangBlog := &Blog{
		programmingLanguage: "Go",
	}

	developers := make([]*Developer, 0)
	developers = append(developers, &Developer{
		Name:         "Nurda",
		MainLanguage: "Go",
	})
	developers = append(developers, &Developer{
		Name:         "Slark",
		MainLanguage: "Java",
	})
	developers = append(developers, &Developer{
		Name:         "Pudge",
		MainLanguage: "C#",
	})
	developers = append(developers, &Developer{
		Name:         "Invoker",
		MainLanguage: "Go",
	})

	for _, developer := range developers {
		if developer.MainLanguage == csharpBlog.programmingLanguage {
			csharpBlog.register(developer)
		}
		if developer.MainLanguage == javaBlog.programmingLanguage {
			javaBlog.register(developer)
		}
		if developer.MainLanguage == goLangBlog.programmingLanguage {
			goLangBlog.register(developer)
		}
	}

	csharpBlog.writeNewPost("What's new in C# 11.0?")
	javaBlog.writeNewPost("Boilerplate code in Java")
	goLangBlog.writeNewPost("Backend development with Go")

}
