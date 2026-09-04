## Learned User Preferences

- Fix distance, reusability, and coverlint by refactor (coverlint-style DTO type aliases and interfaces in stable packages); do not add `//nolint` or edit `.golangci.yml` to pass.
- Keep README aligned with the coverlint sibling linter README (badges, CLI vs plugin, custom golangci-lint).
- Register and document the golangci-lint plugin as `github.com/gostafa/vulner/plugin`, not the module root.

## Learned Workspace Facts

- Module path is `github.com/gostafa/vulner`; the plugin package is `github.com/gostafa/vulner/plugin`.
- Coverlint requires 100% coverage including `cmd/vulner`; use the coverlint testable-main pattern (injectable `run`/`exit`, tests that exercise `main` without `os.Exit`).
- Distance max is 0.30 and reusability min is 0.80; represent DTOs as `type T = struct{...}` aliases so they are not scored as named types.
- Local CI is `task taskotter:ci` / `task taskotter:ci:fix` via Taskotter custom golangci-lint (coverlint, distance, reusability, and vulner).
