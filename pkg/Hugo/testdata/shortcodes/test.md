# Shortcodes

> Shortcodes are simple snippets inside your content files calling built-in or
> custom templates.

{{% mdshortcode %}}Stuff to `process` in the *center*.{{% /mdshortcode %}}

Shortcodes are non-standard markup that appears within normal Markdown.

{{< highlt go >}} A bunch of code here {{< /highlt >}}

Shortcodes are non-standard markup that appears within normal Markdown.

{{<  myshortcode This is some <b>HTML</b>,
and a new line with a "quoted string". >}}

Shortcodes are non-standard markup that appears within normal Markdown.

{{<  myshortcode This is some <b>HTML</b>,. >}}

Shortcodes are non-standard markup that appears within normal Markdown.

{{< myshortcode src="/media/spf13.jpg" title="Steve Francia" >}}

Shortcodes are non-standard markup that appears within normal Markdown.

{{< highlight html >}}
<section id="main">
  <div>
   <h1 id="title">{{ .Title }}</h1>
    {{ range .Pages }}
        {{ .Render "summary"}}
    {{ end }}
  </div>
</section>
{{< /highlight >}}

Shortcodes are non-standard markup that appears within normal Markdown.

{{< instagram BWNjjyYFxVx hidecaption >}}

Shortcodes are non-standard markup that appears within normal Markdown.

[Who]({{< relref "about.md#who" >}})

hidecaption

{{< youtube id="w7Ft2ymGmfc" autoplay="true" >}}

Shortcodes are non-standard markup that appears within normal Markdown.

[Contact us]({{< relref "contact/index.md" >}}) if you have any questions.

Shortcodes are non-standard markup that appears within normal Markdown.
