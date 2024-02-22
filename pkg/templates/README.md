# Templates

## Config

The only required file is `config.yaml`. 
It should contain a list of template files (relative to templates directory) to be rendered
and their destination paths (relative to server directory).
Optionally, you can specify a `reloadable` flag to indicate the file can be updated while the server is running.

It should also contain a list of default values files (relative to templates directory) to be used when rendering the templates.
If multiple files are specified, they are merged through text concatenation so they shouldn't contain any conflicting keys.

Defaults can be overridden by the user through mechanisms described in the main README.

Example `config.yaml`:

```yaml
templates:
  - src: templates/serversettings.con.tpl
    dest: mods/pr/settings/serversettings.con
  - src: templates/realityconfig_admin.py.tpl
    dest: mods/pr/python/game/realityconfig_admin.py
    reloadable: true

defaults:
  - defaults.yaml
```
