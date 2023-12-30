# Themes

amigo supports custom themes that can be enabled for your community members to choose from. To add a custom theme, place it in the `/themes` folder alongside your amigo binary.

## Creating your own theme

You can easily craft your own themes with some knowledge of HTML, CSS, and golang [html/template](https://pkg.go.dev/html/template).

### Anatomy of a Theme
To start with, add a new folder for your theme with the basic structure of:

```sh
.
`-- themes
    `-- mytheme
        |-- assets
        |   `-- style.css
        |-- config.json
        `-- html
            `-- index.html
```

Inside the theme folder, the `config.json` file holds metadat about the theme. Minimally, it should have the name and whether to publish the theme publicly:
```json
{
    "name": "My Theme", // A display name for your theme
    "baseDir": "default23", // The fallback theme to use for missing templates
    "publish": true // Controls availability of the theme
}
```

### Assets
For any CSS, images, icons, or other things your theme requires, you can place them in the `assets` folder which is made available at through the `/assets/<theme-name>/*` path.

Your templates, which go in `html`, can reference assets through this path.

### Templates

Reference `default23` for the required templates in a theme. That said, a key feature of themes is that they use an "override" pattern which allows you to keep your themes minimal/leightweight and focus on only CSS styling.

When a theme is missing a template in the `html` directory, a fallback template is picked from `baseThemeDir` instead.

You only need to include a template if you wish to override the base theme. In general, the following minimal files are needed for a theme to "stand alone"- though you can also just borrow the templates from another theme:
1. A `layout.html` template
2. A `header.html` template
3. A `footer.html` template

Header and Footer are purposely split out to allow for inclusion of custom styling or javascript ontop of existing themes without needing to edit them.