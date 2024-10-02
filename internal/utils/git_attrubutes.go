package utils

func Attribute() string {
	return `
* text=auto

*.png binary
*.jpg binary
*.jpeg binary
*.gif binary
*.exe binary
*.dll binary
*.so binary


*.go text eol=lf


*.mod text eol=lf
*.sum text eol=lf
*.toml text eol=lf
*.json text eol=lf
*.yaml text eol=lf
*.yml text eol=lf

.gitattributes text eol=lf
.gitignore text eol=lf
*.sh text eol=lf

Makefile text eol=lf
`
}
