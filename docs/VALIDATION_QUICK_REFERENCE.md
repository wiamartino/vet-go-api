# Quick Reference Guide: Error Handling & Validation

## Creating and Using Errors

### Error Creation

```go
// In service layer
return utils.NewNotFoundError("Client")
return utils.NewConflictError("Email already registered")
return utils.NewValidationError("Invalid password")
return utils.NewUnauthorizedError("Invalid credentials")
return utils.NewForbiddenError("Access denied")
return utils.NewBadRequestError("Invalid format")
return utils.NewInternalError("Database error")
```

### In Controllers

```go
// Check error type and respond
if err != nil {
    if utils.IsAppError(err) {
        utils.RespondWithAppError(c, err.(*utils.AppError))
    } else {
        utils.RespondWithAppError(c, utils.AsAppError(err))
    }
    return
}
```

---

## Validation Checklist

### 1. Required Fields
```go
validator.ValidateRequired("field_name", value)
```

### 2. String Length
```go
// Either/or:
validator.ValidateMinLength("name", name, 2)
validator.ValidateMaxLength("name", name, 100)
validator.ValidateLengthRange("name", name, 2, 100) // Both
```

### 3. Email
```go
validator.ValidateEmail("email", email)
```

### 4. Phone
```go
validator.ValidatePhone("phone", phone) // 10-15 digits
```

### 5. Date (ISO 8601)
```go
validator.ValidateDate("birth_date", "2023-12-15T00:00:00Z")
```

### 6. Numeric IDs
```go
validator.ValidateNumericID("client_id", client.ClientID) // Must be > 0
```

### 7. Numeric Ranges
```go
validator.ValidateMinValue("age", age, 0)
validator.ValidateMaxValue("price", price, 99999.99)
```

### 8. Enum/List Validation
```go
validator.ValidateInSlice("status", status, 
    []string{"active", "inactive", "pending"})
```

---

## Controller Pattern Template

### Copy-Paste Template for New Controllers

```go
func (ctrl *SomeController) CreateResource(c *gin.Context) {
    // 1. Parse request
    var resource domain.Resource
    if err := c.ShouldBindJSON(&resource); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, 
            "Invalid request format: "+err.Error())
        return
    }

    // 2. Validate input
    validator := utils.NewValidator()
    validator.ValidateRequired("field1", resource.Field1)
    validator.ValidateLengthRange("field2", resource.Field2, 2, 100)
    validator.ValidateEmail("email", resource.Email)
    validator.ValidatePhone("phone", resource.Phone)
    // ... add all validations

    if !validator.IsValid() {
        utils.RespondWithValidationError(c, validator.Errors)
        return
    }

    // 3. Call service
    if err := ctrl.service.Create(&resource); err != nil {
        if utils.IsAppError(err) {
            utils.RespondWithAppError(c, err.(*utils.AppError))
        } else {
            utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
        }
        return
    }

    // 4. Return 201 Created
    utils.RespondWithCreated(c, resource)
}

func (ctrl *SomeController) UpdateResource(c *gin.Context) {
    // Parse ID
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        utils.RespondWithAppError(c, 
            utils.NewBadRequestError("Invalid ID format"))
        return
    }

    // Parse request
    var resource domain.Resource
    if err := c.ShouldBindJSON(&resource); err != nil {
        utils.RespondWithError(c, http.StatusBadRequest, 
            "Invalid request format: "+err.Error())
        return
    }

    // Check exists
    if _, err := ctrl.service.GetByID(uint(id)); err != nil {
        utils.RespondWithAppError(c, utils.NewNotFoundError("Resource"))
        return
    }

    // Validate
    validator := utils.NewValidator()
    // ... add validations

    if !validator.IsValid() {
        utils.RespondWithValidationError(c, validator.Errors)
        return
    }

    // Update
    resource.ID = uint(id)
    if err := ctrl.service.Update(&resource); err != nil {
        utils.RespondWithAppError(c, utils.AsAppError(err))
        return
    }

    utils.RespondWithSuccess(c, http.StatusOK, resource)
}

func (ctrl *SomeController) DeleteResource(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        utils.RespondWithAppError(c, 
            utils.NewBadRequestError("Invalid ID format"))
        return
    }

    // Check exists
    if _, err := ctrl.service.GetByID(uint(id)); err != nil {
        utils.RespondWithAppError(c, utils.NewNotFoundError("Resource"))
        return
    }

    if err := ctrl.service.Delete(uint(id)); err != nil {
        utils.RespondWithAppError(c, utils.AsAppError(err))
        return
    }

    // Return 204 No Content
    utils.RespondWithNoContent(c)
}
```

---

## Response Helper Quick Reference

### Success Responses

```go
// 200 OK (for GET, PUT)
utils.RespondWithSuccess(c, http.StatusOK, data)

// 201 Created (for POST)
utils.RespondWithCreated(c, data)

// 204 No Content (for DELETE)
utils.RespondWithNoContent(c)
```

### Error Responses

```go
// Generic error with custom message
utils.RespondWithError(c, http.StatusBadRequest, "message")

// AppError with auto status code
utils.RespondWithAppError(c, appErr)

// Validation errors with field details
utils.RespondWithValidationError(c, validator.Errors)
```

---

## HTTP Status Code Reference

| Code | Use Case | Example |
|------|----------|---------|
| **200** | Success (GET, PUT) | Get/update succeeded |
| **201** | Created (POST) | Resource created |
| **204** | No Content (DELETE) | Deleted successfully |
| **400** | Bad Request | Invalid JSON, validation error |
| **401** | Unauthorized | Missing/invalid token |
| **403** | Forbidden | User lacks permission |
| **404** | Not Found | Resource doesn't exist |
| **409** | Conflict | Duplicate email, business rule violation |
| **500** | Internal Error | Server/database error |

---

## Testing Error Responses

### Test Pattern for Error Cases

```go
t.Run("CreateResource - should return validation error", func(t *testing.T) {
    // Setup
    mockRepo := new(mocks.MockRepository)
    service := application.NewService(mockRepo)
    controller := controllers.NewController(service)
    
    router := testhelpers.SetupTestRouter()
    router.POST("/resources", controller.Create)
    
    // Invalid data
    body := []byte(`{"name": "", "email": "invalid"}`)
    req, _ := http.NewRequest("POST", "/resources", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    // Assert
    assert.Equal(t, http.StatusBadRequest, w.Code)
    
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "error", response["status"])
    assert.Equal(t, "VALIDATION_ERROR", response["code"])
})
```

---

## Common Validation Patterns

### Email Validation
```go
// Always validate emails
validator.ValidateEmail("email", client.Email)
```

### Phone Validation
```go
// Phone must be 10-15 digits (with optional formatting)
validator.ValidatePhone("phone", client.Phone)
```

### Name Fields
```go
// Names: 2-100 characters
validator.ValidateLengthRange("first_name", client.FirstName, 2, 100)
validator.ValidateLengthRange("last_name", client.LastName, 2, 100)
```

### Address Validation
```go
// Addresses: 5-255 characters
validator.ValidateLengthRange("address", client.Address, 5, 255)
```

### Password Validation
```go
// Passwords: 8+ characters minimum
validator.ValidateMinLength("password", password, 8)
```

### Numeric IDs
```go
// All foreign keys must be > 0
validator.ValidateNumericID("pet_id", pet.PetID)
validator.ValidateNumericID("client_id", pet.ClientID)
```

### Status/Enum Validation
```go
// Only allow specific values
validator.ValidateInSlice("status", status, 
    []string{"pending", "completed", "cancelled"})
```

---

## Migration Checklist

### For Each Remaining Controller:

- [ ] Add input validation in Create method
- [ ] Add input validation in Update method
- [ ] Add ID format validation in single-resource endpoints
- [ ] Check if resource exists before update/delete
- [ ] Use `RespondWithCreated` for POST (201)
- [ ] Use `RespondWithNoContent` for DELETE (204)
- [ ] Use `RespondWithAppError` for error responses
- [ ] Remove generic 500 errors for validation failures
- [ ] Add validation tags to domain model
- [ ] Test with invalid data to verify error responses

---

## Domain Model Binding Tags Reference

```go
type SomeModel struct {
    // String fields
    Name    string `json:"name" binding:"required,min=2,max=100"`
    Email   string `json:"email" binding:"required,email"`
    Phone   string `json:"phone" binding:"required,min=10,max=20"`
    
    // Numeric fields
    Age     int    `json:"age" binding:"required,min=0,max=150"`
    Price   float64 `json:"price" binding:"required,min=0"`
    
    // Foreign keys
    ClientID uint `json:"client_id" binding:"required"`
    
    // Enum (use in validator, not binding)
    Status string `json:"status"` // Validate with ValidateInSlice
}
```

---

## Example: Complete Controller with All Features

See `controllers/clients.go` for a full working example with:
✅ Input parsing  
✅ Comprehensive validation  
✅ Error handling  
✅ Proper HTTP status codes  
✅ Clear error messages  

---

**Quick Links:**
- [Error System](/users/walteriamartino/Documents/vet-go/utils/errors.go)
- [Validator System](/users/walteriamartino/Documents/vet-go/utils/validator.go)
- [Response Helpers](/users/walteriamartino/Documents/vet-go/utils/response.go)
- [Client Controller Example](/users/walteriamartino/Documents/vet-go/controllers/clients.go)
