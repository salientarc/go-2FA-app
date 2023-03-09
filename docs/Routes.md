# API Routes Examples

## Available Routes

```go
func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
    router := rg.Group("auth")

    router.POST("/register", rc.authController.SignUpUser)
    router.POST("/login", rc.authController.LoginUser)
    router.POST("/otp/generate", rc.authController.GenerateOTP)
    router.POST("/otp/verify", rc.authController.VerifyOTP)
    router.POST("/otp/validate", rc.authController.ValidateOTP)
    router.POST("/otp/disable", rc.authController.DisableOTP)
}
```

## Request Commands

### Register User

```sh
# register user
curl -d '{"name": "user", "email": "test@example.com", "password": "testpass"}'\
     -H "Content-Type: application/json"\
     -X POST http://localhost:8000/api/auth/register | json
```

Sample Response

```json
[
  {
    "status": "success",
    "message": "Registered successfully, please login"
  }
]
```

### Login User

```sh
# login user
curl -d '{"email": "test@example.com", "password": "testpass"}'\
     -H "Content-Type: application/json"\
     -X POST http://localhost:8000/api/auth/login | json
```

Sample Reponse

```json
[
  {
    "status": "success",
    "user": {
      "email": "test@example.com",
      "id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c",
      "name": "user",
      "otp_enabled": false
    }
  }
]
```

### Register OTP

```sh
# register OTP

curl -d '{"user_id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c"}' \
     -H "Content-Type: application/json" \
     -X POST http://localhost:8000/api/auth/otp/generate | json
```

Sample Response

```json
[
  {
    "base32": "YOFNVNJKMLRTNWVHANMVRWR4",
    "otpauth_url": "otpauth://totp/domain.com:admin@domain.com?algorithm=SHA1&digits=6&issuer=domain.com&period=30&secret=YOFNVNJKMLRTNWVHANMVRWR4"
  }
]
```

Copy the `base32` secret key value and add a manual entry in your authenticator

### Verify OTP

```sh
# verify OTP

curl -d '{"user_id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c", "token": "254966"}' \
     -H "Content-Type: application/json" \
     -X POST http://localhost:8000/api/auth/otp/verify | json
```

Sample Response

```json
[
  {
    "otp_verified": true,
    "user": {
      "email": "test@example.com",
      "id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c",
      "name": "user",
      "otp_enabled": true
    }
  }
]
```

### Validate OTP

```sh
# validate OTP

curl -d '{"user_id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c", "token": "236673"}' \
     -H "Content-Type: application/json"
     -X POST http://localhost:8000/api/auth/otp/validate | json
```

Sample Response

```json
[
  {
    "otp_valid": true
  }
]
```

### Disable OTP

```sh
# disable OTP

curl -d '{"user_id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c"}' \
     -H "Content-Type: application/json" \
     -X POST http://localhost:8000/api/auth/otp/disable | json
```

Sample Response

```json
[
  {
    "otp_disabled": true,
    "user": {
      "email": "test@example.com",
      "id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c",
      "name": "user",
      "otp_enabled": false
    }
  }
]
```