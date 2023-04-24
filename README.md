![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)
# leggo

leggo is a tool that i built for myself to quickly scaffold new projects. it makes a directory in my `~/Projects` folder with the program name and sets up `git` and other stuff i usually want there.

## build

since this is specific to my workflow, there some hard coded values in there that i encourage you to change to suit yourself, unless you want to credit me with your work :)

make the changes you want and then compile it like this:
```bash
git clone https://github.com/donuts-are-good/leggo.git
cd leggo
go build
```

## usage

to use `leggo` run the program like this: 

```bash
leggo my-new-project
```

this will create a new directory named `my-new-project` in your `~/Projects` folder and set up these files:

- `main.go`: a "hello, world" program to get started with
- `readme.md`: a markdown file with the name of the project
- `license.md`: a file crediting me for your work :)
- `.gitignore`: stuff i commonly forget to ignore in git
- `go.mod`: idk some go stuff
- a `git` repository with an initial commit for these files

and then, when all this has happened, it opens vscode in that directory.

**reminder**: this tool is specific to my workflow, feel free to change what you need to change, recompile it, and then use it for yourself. 

## license

2023 mit license. if you don't know what it means, don't worry about it. see `license.md` for more.