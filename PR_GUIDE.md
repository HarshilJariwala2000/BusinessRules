# Pull Request and Code Review Guide

## 1. Branching Strategy
We use a standard branching strategy for this repository. 
Please branch out from `main` or your designated environment branch.

- **Feature branches:** `feature/<your-username>/<short-description>` (e.g. `feature/johnd/add-formula-modulo`)
- **Bug fixes:** `bugfix/<issue-id-or-description>` (e.g. `bugfix/evaluator-division-by-zero`)
- **Hotfixes:** `hotfix/<description>`

## 2. Commit Messages (Conventional Commits)
Write clean, readable commit messages. Use the Conventional Commits format:
- `feat: added formula compiler`
- `fix: resolved DAG topological sorting panic`
- `docs: updated style guide`
- `refactor: extracted handler to service`

## 3. Pull Request Requirements
Before submitting a PR, ensure you have checked all relevant items:

### Pre-submission Checklist
- [ ] Code is formatted with `gofmt`.
- [ ] Code builds without errors (`go build ./...`).
- [ ] You have manually tested your changes.
- [ ] All errors are explicitly handled.
- [ ] Naming conventions align with the `STYLE_GUIDE.md`.

### Pull Request Template
Please use the following template structure when adding your PR description in GitHub / GitLab:

```md
## Context
Provide an overview of the problem being solved or the feature being implemented.

## Changes Made
- Added feature X to calculation engine.
- Refactored DAG node traversals.
- Updated database schema for CategoryAssignments.

## How to Test
1. Make a POST request to `/v1/formula/create` with payload X.
2. Confirm the 201 response.
3. Validate Topological sort execution correctness in DB.

## Additional Notes
(Optional) Note any trade-offs, technical debt, or things reviewers should pay special attention to.
```

## 4. Code Review Expectations
- **Reviewers:** Focus on architectural compliance (Separation of layers), error propagation, DB transaction handling, and security.
- **Authors:** Re-request reviews once a requested change line item has been fully addressed. Resolve conversation threads once addressed.
