# register user
curl -d '{"name": "user", "email": "test@example.com", "password": "testpass"}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/auth/register | json

# login user
curl -d '{"email": "test@example.com", "password": "testpass"}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/auth/login | json

# register OTP
curl -d '{"user_id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c"}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/auth/otp/generate | json

# verify OTP
curl -d '{"user_id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c", "token": "254966"}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/auth/otp/verify | json

# validate OTP
curl -d '{"user_id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c", "token": "236673"}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/auth/otp/validate | json

# disable OTP
curl -d '{"user_id": "fbe82175-9884-4cee-a4ee-dd2c80a0b03c"}' -H "Content-Type: application/json" -X POST http://localhost:8000/api/auth/otp/disable | json