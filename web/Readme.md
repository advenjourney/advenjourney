# AdvenJourney Website
Based on VuePress, find detailed guides here: https://vuepress.vuejs.org/

## Run
### Local development

```
yarn docs:dev
```
Browse to http://localhost:8080

## Folder structure
### Website Contents

- docs

### Special folders
- docs/.vuepress: Used to store global configuration, components, static resources, etc.
- docs/.vuepress/components: The Vue components in this directory will be automatically registered as global components.
- docs/.vuepress/theme: Used to store local theme.
- docs/.vuepress/styles: Stores style related files.
- docs/.vuepress/styles/index.styl: Automatically applied global style files, generated at the ending of the CSS file, have a higher priority than the default style.
- docs/.vuepress/styles/palette.styl: The palette is used to override the default color constants and to set the color constants of Stylus.
- docs/.vuepress/public: Static resource directory.
- docs/.vuepress/templates: Store HTML template files.
- docs/.vuepress/templates/dev.html: HTML template file for development environment.
- docs/.vuepress/templates/ssr.html: Vue SSR based HTML template file in the built time.
- docs/.vuepress/config.js: Entry file of configuration, can also be yml or toml.
- docs/.vuepress/enhanceApp.js: App level enhancement.