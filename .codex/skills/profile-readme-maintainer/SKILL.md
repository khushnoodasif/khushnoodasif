---
name: profile-readme-maintainer
description: Use when editing, reviewing, or regenerating the khushnoodasif GitHub profile README or its update helper, including profile copy, badges, toolbox sections, DevOps and SRE positioning, Markdown or HTML formatting, and the Go updater under update.
---

# Profile README Maintainer

## Purpose

Use this skill for the `khushnoodasif` profile repository. The primary artifact is the public GitHub profile `README.md`; the `update/` Go helper can modify it and must be treated carefully.

## README Style

- Preserve the opening Arabic greeting and centered identity block unless the user explicitly asks to change them.
- Keep the profile calm, professional, and concise.
- Keep the positioning centered on DevOps, platform engineering, production operations, infrastructure automation, observability, security, and safe AI-assisted operations.
- Prefer concrete work areas and engineering principles over hype.
- Keep Markdown and inline HTML simple enough for GitHub profile rendering.

## Content Rules

- No secrets, private employer details, private homelab hostnames, private URLs, or unverifiable production claims.
- Keep claims current and supportable.
- Keep the toolbox grouped by domain and avoid long uncurated badge lists.
- Keep contact links current when the user asks for contact changes.

## Badge And Link Rules

- Keep badges consistent with the existing shields.io style.
- Link only to public destinations.
- Do not add generated metrics, activity graphs, or external widgets unless the user asks and the data source is public-safe.

## Go Updater Rules

- Read `update/main.go` before running or changing the updater.
- Do not run `go run .` casually; the updater writes to `../README.md` and can replace README content.
- If changing the updater, prefer making it testable against temporary files before using it on the real README.

## Validation

Always check whitespace from the repo root:

```bash
git diff --check
```

For updater changes:

```bash
cd update
go fmt ./...
go test ./...
```

Only when intentionally testing README generation:

```bash
cd update
go run .
```

When network access is available and a change touches public links, verify changed links with a header request or browser check. If link validation cannot run, say so.
