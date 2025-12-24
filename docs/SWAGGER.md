# Swagger API Documentation

This project includes comprehensive API documentation using Swagger (OpenAPI 3.0).

## Accessing Swagger UI

Once the application is running, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

## Using the API with Authentication

Most endpoints require authentication. Follow these steps:

1. **Register a user** (if needed):
   - Use the `POST /auth/register` endpoint
   - Provide name, email, and password

2. **Login**:
   - Use the `POST /auth/login` endpoint
   - Provide email and password
   - Copy the token from the response

3. **Authorize in Swagger**:
   - Click the "Authorize" button (lock icon) at the top right
   - Enter: `Bearer YOUR_TOKEN_HERE`
   - Click "Authorize" then "Close"

4. **Make authenticated requests**:
   - All subsequent requests will include the authorization header automatically

## Regenerating Documentation

After adding or modifying Swagger annotations in the code, regenerate the documentation:

```bash
~/go/bin/swag init
```

Or if swag is in your PATH:

```bash
swag init
```

## Adding Swagger Annotations

### Basic Endpoint Structure

```go
// FunctionName does something
// @Summary Short description
// @Description Longer description of what this endpoint does
// @Tags CategoryName
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param paramName paramType dataType required "description"
// @Success 200 {object} ResponseType "Success message"
// @Failure 400 {object} map[string]string "Error message"
// @Router /endpoint/path [method]
func (ctrl *Controller) FunctionName(c *gin.Context) {
    // Implementation
}
```

### Common Annotations

- `@Summary`: Brief one-line description
- `@Description`: Detailed multi-line description
- `@Tags`: Group endpoints by category
- `@Accept`: Input content type (json, xml, etc.)
- `@Produce`: Output content type
- `@Param`: Define parameters
  - Path params: `@Param id path int true "User ID"`
  - Query params: `@Param page query int false "Page number"`
  - Body params: `@Param user body domain.User true "User object"`
- `@Success`: Success response format
- `@Failure`: Error response format
- `@Security`: Authentication requirement
- `@Router`: Endpoint path and HTTP method

### Parameter Types

- `path`: URL path parameter (e.g., `/users/{id}`)
- `query`: Query string parameter (e.g., `?page=1`)
- `body`: Request body (JSON)
- `header`: HTTP header
- `formData`: Form data

### Examples

#### Public Endpoint (No Auth)
```go
// Login authenticates a user
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Failure 401 {object} ErrorResponse
// @Router /auth/login [post]
func (ctrl *AuthController) Login(c *gin.Context) {
```

#### Protected Endpoint
```go
// GetClient retrieves a client by ID
// @Summary Get client by ID
// @Description Get detailed information about a specific client
// @Tags Clients
// @Security BearerAuth
// @Produce json
// @Param id path int true "Client ID"
// @Success 200 {object} domain.Client
// @Failure 404 {object} ErrorResponse
// @Router /clients/{id} [get]
func (ctrl *ClientController) GetClient(c *gin.Context) {
```

## Available Endpoints

The Swagger UI provides an interactive interface to:
- View all available endpoints
- See request/response schemas
- Test endpoints directly from the browser
- Download API specifications (JSON/YAML)

## API Groups

- **Authentication**: User registration, login, token refresh
- **Clients**: Client management
- **Pets**: Pet registration and management
- **Appointments**: Appointment scheduling
- **Veterinarians**: Veterinarian management
- **Medical Records**: Medical history and records
- **Treatments**: Treatment management
- **Medications**: Medication records
- **Vaccinations**: Vaccination tracking
- **Surgeries**: Surgery records
- **Invoices**: Billing and invoices
- **Allergies**: Allergy information

## Additional Resources

- [Swagger Official Documentation](https://swagger.io/docs/)
- [swaggo Documentation](https://github.com/swaggo/swag)
- [OpenAPI 3.0 Specification](https://swagger.io/specification/)
