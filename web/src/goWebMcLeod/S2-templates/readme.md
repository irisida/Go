![](/assets/gologo.png)

# Web development with Go - Section 2 - Templates and rendering templates

Template parsing with....

- [Html](/web/src/goWebMcLeod/S2-templates/01-parsingHtml)
- [Html and writing to file](/web/src/goWebMcLeod/S2-templates/02-parsing-writingHtml)
- [Multiple templates](/web/src/goWebMcLeod/S2-templates/03-parsing-multiple-templates)
- [Globs](/web/src/goWebMcLeod/S2-templates/04-parsing-globs)
- [Data interpolation](/web/src/goWebMcLeod/S2-templates/05-with-data)
- [Variables](/web/src/goWebMcLeod/S2-templates/06-with-variables)
- [Data structures - slice](/web/src/goWebMcLeod/S2-templates/07-with-slice)
- [Data structures - slice and variables](/web/src/goWebMcLeod/S2-templates/08-with-slice-and_variables)
- [Data structures - map](/web/src/goWebMcLeod/S2-templates/09-with-map)
- [Data structures - struct](/web/src/goWebMcLeod/S2-templates/10-with-struct)
- [Data structures - nested data](/web/src/goWebMcLeod/S2-templates/11-nested-data)
- [Functions](/web/src/goWebMcLeod/S2-templates/12-with-functions)
- [Time and time formatting](/web/src/goWebMcLeod/S2-templates/13-time-formatting)
- [Pipelines](/web/src/goWebMcLeod/S2-templates/14-pipelines)
- [Global functions](/web/src/goWebMcLeod/S2-templates/15-global-functions)
- [Composition](/web/src/goWebMcLeod/S2-templates/16-with-composition)
- [Methods](/web/src/goWebMcLeod/S2-templates/16-with-methods)

#### key takeaways

1. we essentially get one chance to pass data to a template. The `template.ExecuteTemplate(location, filename, data-value)` call is what facilitates the passing of the data. Such a limitation on face value usually triggers a fairly negative response but its worth noting that you can pass in aggregated data types, structs, maps etc etc.

2. We can use the `{{.}}` to render the value(s) passed.

3. We can use variables in a template in a way that is familiar to some JS and PHP type template engines. We declare the variables `{{ $variableName = . }}` then render the `{{ $variableName }}`.

4. When passing in data structures and nesting a functional piping is implicit.

5. FuncMaps form part of the template construction as static compilation is required to work correctly.

![](/web/src/goWebMcLeod/S2-templates/assets/201-tpl.png)

![](/web/src/goWebMcLeod/S2-templates/assets/202-more-tpl.png)

![](/web/src/goWebMcLeod/S2-templates/assets/203-multiple-templates.png)

![](/web/src/goWebMcLeod/S2-templates/assets/204-globs.png)

![](/web/src/goWebMcLeod/S2-templates/assets/205-with-data.png)

![](/web/src/goWebMcLeod/S2-templates/assets/209-maps.png)

![](/web/src/goWebMcLeod/S2-templates/assets/214-pipelines.png)

![](/web/src/goWebMcLeod/S2-templates/assets/215-functions.png)

![](/web/src/goWebMcLeod/S2-templates/assets/217-methods.png)
