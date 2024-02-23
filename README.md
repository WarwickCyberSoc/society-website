# Warwick Cyber Society Website

## Hosting
- This is hosted on github pages
[https://warwickcybersoc.com](https://warwickcybersoc.com)

## Editing
To add new static assets add them to the public folder.
To add new templated pages create a new template in the templates folder with the base file of:
```template
{{defeine "head"}}
# This can be left empty but is needed to prevent errors
{{end}}
{{define "content"}}
# Content goes here
{{end}}
```

The go program follows the following logic:
- Remove build folder if exists
- Create new build folder
- Copy all assets inside public to build folder
- Load the config.yaml file
- Run every template in the templates directory and add those to the build folder with the config context

## Adding new config values
To add new values to a template you need to add the parameter to the Config stuct in `config.go` e.g
```go
type Config struct {
    .... Existing data
    NewData string `yaml:"newData"`
}
```
and then to the config.yaml file
```yaml
... Existing data
newData: new value
```

This can then be referenced in the templates by `.NewData`

## Building
There is a github action set to automatically build and publish the site when pushed to the master branch


### Troublshooting
#### <variable> is an unexported field
Variables in structs need to start with a capital letter to be exported


#### Other
Feel free to contact me on discord (tomsteer) or email (tom@tomsteer.me)

