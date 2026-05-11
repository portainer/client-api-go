---
name: release
description: Cut a new SDK release for client-api-go. Rotates the [Unreleased] section in CHANGELOG.md to a versioned section, bumps VERSION in Makefile and the install examples in README.md, commits, and creates a local git tag. Does NOT push and does NOT regenerate the swagger client (run `make generate-client VERSION=X.Y.Z` separately, beforehand). Invoked as `/release X.Y.Z` or `/release vX.Y.Z`.
---

# /release skill

Cut a new release of `client-api-go`. The version arg should match the upstream Portainer server version this release targets (see CLAUDE.md for why these are coupled).

## Inputs

- One positional arg: the version, with or without a leading `v` (e.g. `2.31.3` or `v2.31.3`).
  - Tag form (with `v`): used for the git tag.
  - Plain form (no `v`): used in `Makefile` and `README.md`.

If no version was provided, ask the user before doing anything else.

## Pre-flight checks

Run these in parallel and abort if any fails:

1. `git status --porcelain` — working tree must be clean. If dirty, stop and tell the user which files are dirty; do not proceed.
2. `git rev-parse --abbrev-ref HEAD` — record the current branch (informational; do not require `main`, but mention it in the summary if it isn't `main`).
3. `git tag --list vX.Y.Z` — the target tag must not already exist. If it does, abort.
4. Read `Makefile` line 1 to capture the previous version (`VERSION := A.B.C`). You will need it to find/replace in `README.md`.
5. Read `CHANGELOG.md` to confirm a `## [Unreleased]` section exists.

If the `## [Unreleased]` section has no bullet points under any of its category headings, warn the user that the release will have an empty changelog and ask whether to proceed.

## Edits

Make these edits, then show the user the resulting diff before committing.

### CHANGELOG.md

- Replace the `## [Unreleased]` heading with a fresh empty Unreleased block followed by a versioned heading dated today. Use today's date in `YYYY-MM-DD` (UTC is fine; do not invent a timezone).
- Drop empty category subheadings (`### Added`, `### Changed`, etc.) from the *new versioned* section — keep only categories that have at least one bullet under them.
- The new `## [Unreleased]` block should re-seed the full set of empty category headings, matching the seed in CHANGELOG.md's initial commit.

Resulting structure:

```markdown
## [Unreleased]

### Added

### Changed

### Deprecated

### Removed

### Fixed

### Security

## [vX.Y.Z] - YYYY-MM-DD

### <only categories that had entries>

- <preserved bullets>
```

### Makefile

- Line 1: `VERSION := <old>` → `VERSION := <new-plain>` (no `v` prefix).

### README.md

- Replace every occurrence of the old plain version (e.g. `2.31.2`) with the new plain version. Also replace `v<old>` (e.g. `v2.31.2`) with `v<new>`. Use a tool that fails if the search string is not found, so a stale README that no longer mentions the old version is caught loudly rather than silently no-op'd.

Do not touch `pkg/`, `swagger.yaml`, or `go.mod` — those are regeneration outputs and out of scope here.

## Confirm and commit

1. Run `git diff` and show the user a short summary (files changed + the new CHANGELOG entry headline). Ask for explicit confirmation before committing.
2. On confirm:
   - `git add CHANGELOG.md Makefile README.md`
   - `git commit -m "version: bump to X.Y.Z"` (matches the existing repo commit style — see `git log --oneline | head` for prior `version:` commits)
   - `git tag vX.Y.Z`
3. Do NOT push and do NOT push the tag. Print the exact push commands the user can run themselves:
   ```
   git push origin <branch>
   git push origin vX.Y.Z
   ```

## Edge cases

- **Empty Unreleased**: ask the user whether they really want an empty release entry. If they confirm, proceed and write a single bullet `- No user-visible changes; release tracks Portainer server vX.Y.Z.` under a new `### Changed` category (or similar) so the entry isn't blank.
- **Old version string not found in README.md**: do not silently skip. Stop, show the user the README, and ask how to handle it (the README may have been hand-edited and drifted from `Makefile`).
- **Working tree dirty**: never auto-stash. Abort and tell the user.
- **Tag already exists**: abort. Do not delete or overwrite tags.
- **User invokes from a detached HEAD or non-`main` branch**: proceed but mention this in the summary so the user knows where the tag landed.

## Out of scope (do these manually before invoking the skill)

- Regenerating the swagger client (`make generate-client VERSION=X.Y.Z`).
- Reviewing the regenerated `pkg/` diff and updating `client/` to match any breaking upstream changes.
- Adding the corresponding bullets to `[Unreleased]` for those upstream changes.
- Pushing commits and tags to the remote.
