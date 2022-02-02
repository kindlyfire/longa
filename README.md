# Longa

_("long" in Esperanto)_

Simple, small single-user URL shortner app built with Go + an admin interface in React.

![Screenshot of admin dashboard](./screenshot.png)

## To-Do

- [ ] Back-end
  - [x] Config from toml files or env variables
  - [x] SQLite Database
  - [x] Authentication
  - [x] List, create, update and delete links
    - [ ] Add pagination and search to list endpoint
  - [x] Functional redirects
- [ ] Front-end
  - [x] Authentication
  - [x] List links
    - [ ] Move pagination to server-side
    - [ ] Search
  - [x] Create link
  - [x] Delete link
  - [ ] Update link
  - [x] Mobile friendly

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
