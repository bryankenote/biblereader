# BibleReader

### Install Depenedencies

- [templ](https://templ.guide/quick-start/installation)
- clone [BibleApi](https://github.com/bryankenote/bibleapi) and [BibleUtils](https://github.com/bryankenote/bibleutils) to the same directory and create a go workspace `go work init BibleReader BibleApi BibleUtils`

### Run

```
make run
```

## Development

### View Templates

After making changes to any view template (.templ files), go files can be generated using `make templ`. These are also generated when running with `make run`.
