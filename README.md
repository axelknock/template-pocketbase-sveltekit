# template-pocketbase-sveltekit

GitHub template for a **PocketBase (Go)** backend + **SvelteKit (TypeScript)** frontend.

## Structure

- `pocketbase/` – Go PocketBase app with Go-based migrations registered in `main.go`
- `web/` – SvelteKit frontend scaffold
- `Taskfile.yml` – common development tasks

## Quick start

```bash
cd pocketbase && go run . serve
```

```bash
cd web && pnpm install && pnpm dev
```

Or from repo root with `task`:

```bash
task web-install
task dev
```

## Migrations

This template uses PocketBase Go migrations per:
https://pocketbase.io/docs/go-migrations/

Create a migration:

```bash
task migration-create -- add_new_collection
```

Apply migrations:

```bash
task migration-up
```

> After creating your own repository from this template, update the Go module path in `pocketbase/go.mod` and `pocketbase/main.go` import path if needed.
