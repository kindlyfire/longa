# Longa

_("long" in Esperanto)_

Simple, small single-user URL shortner app built with Go + an admin interface in React.

![Screenshot of admin dashboard](./screenshot.png)

## Building from source

Requires: Go 1.17, GCC (for sqlite), NodeJS, Yarn

```sh
# In the root directory of the project
cd frontend
yarn
yarn build
cd ..
go build
```

This should output an executable named `longa.exe` on Windows or `longa` everywhere else.
