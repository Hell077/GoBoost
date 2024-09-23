package utils

func Attribute() string {
	return `
* text=auto

*.go text eol=lf
*.md text eol=lf
*.yaml text eol=lf

*.png binary
*.jpg binary
*.gif binary
*.zip binary
`
}
