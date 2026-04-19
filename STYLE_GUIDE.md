# Go Project Style Guide

## 1. Code Formatting
- **Standard Formatting:** All code MUST be formatted using `gofmt` or `goimports`. No exceptions.
- **Line Length:** Avoid exceptionally long lines. Keep lines well within 120 characters to improve readability.
- **Indentation:** Use tabs (the Go standard) for indentation, not spaces.

## 2. Naming Conventions
- **Packages:** Package names should be short, lowercase, and one-word (e.g., `router`, `service`, `store`). Avoid `_` or mixedCaps.
- **Variables and Functions:** 
  - Use `camelCase` for unexported (private) variables and functions.
  - Use `PascalCase` for exported (public) variables and functions.
- **Acronyms:** Keep acronyms consistently cased. Example: `APIResponse`, `URLParser` (not `ApiResponse` or `UrlParser`).

## 3. Web Framework and Routing (Gin)
- **Validation:** Always use structured validation mapping. Use the standard `validator` tags (`validate:"required"`) on request structs.
- **Responses:** Follow a consistent, structured JSON response format like the `ApiResponse` struct:
  ```go
  type ApiResponse struct{
      Message string  `json:"message"`
      Data    []any   `json:"data"`
  }
  ```
- **Error Handling:** Send standard HTTP error status codes (400 for bad request, 404 for not found, 500 for internal server error). Do not rely on 200 OK for logical errors. Write descriptive user-friendly error messages.

## 4. Architecture Standards
- **Separation of Concerns:** 
  - The API router layer should *only* handle parsing, structural validation, and dispatching.
  - The Service layer handles *all* specific business logic, rule implementation, and orchestrations.
  - The Store layer strictly handles interactions with the underlying datastore (PostgreSQL/GORM). 
- **Dependency Injection:** Database connections or dependencies should ideally be passed down or accessed locally via connection pools, preventing hard coupling.

## 5. Lexer/Parser Rules
- When editing calculation engine components, explicitly follow the current Pratt Parser structures. Abstract syntax structures (AST nodes) must implement standard `Node`, `Statement`, or `Expression` interfaces.

## 6. Comments and Documentation
- Provide a comment on exported package variables or functions describing its function. 
- Use standard `//` formatting.
- Ex:
  ```go
  // Eval interprets a node using the provided environment mapping.
  func Eval(node parser.Node, env *Environment) Object {
  ```
