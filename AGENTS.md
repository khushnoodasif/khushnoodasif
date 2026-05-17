# AGENTS.md

Repo-local guidance for the khushnoodasif profile repository.

## Repo-Local Skills

Additional Codex skills live under `.codex/skills/`.

- Use `profile-readme-maintainer` for the public GitHub profile README, badges,
  profile copy, DevOps/SRE positioning, Markdown/HTML formatting, links, and
  the Go updater under `update/`.

## Profile Rules

- Preserve the existing professional DevOps, production platforms,
  infrastructure automation, observability, security, and safe AI operations
  positioning unless the user asks for a different direction.
- Do not add secrets, private employer details, private homelab URLs, or
  unverifiable production claims.
- Treat `update/go run .` as potentially destructive because it writes to the
  profile README.

## Validation

- Run `git diff --check` for README changes.
- For updater changes, run `go fmt ./...` and `go test ./...` from `update/`.
