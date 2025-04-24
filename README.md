# Amartha Test  
# Loan Engine - REST API

Loan management system with a state machine (Proposed → Approved → Invested → Disbursed), including investor funding, disbursement handling, and automated profit calculation.

---

## Business Flow (State Machine)

### 1. `Proposed`
- Initial status when the loan is created
- Information required:
  - `borrower_id`, `principal_amount`, `rate`, `roi`

### 2. `Approved`
- Approved by a field validator
- Additional required data:
  - Picture proof of visit, employee ID, approval date

### 3. `Invested`
- Status when 100% of the loan principal is funded
- Supports multiple investors
- Automatic email to all investors with agreement letter link when fully funded

### 4. `Disbursed`
- Final state after borrower signs agreement
- Required information:
  - Signed agreement file, disbursing employee ID, disbursement date
- Auto calculations:
  - Total amount to be paid by borrower
  - Return to investors based on ROI

---

## Tech Stack

- **Backend**: Go
- **Database**: PostgreSQL (UUID Primary Keys)
- **ORM**: GORM (Go)
- **Auth**: JWT with Role-Based Middleware
- **Architecture**: Monolithic Service
- **API Documentation**: Swagger (OpenAPI)

---

## Database Schema

- `loans`  
- `loans_approvals`  
- `loans_investments`  
- `loan_disbursements`  
- `loan_repayments_summary`  
- `loan_investor_returns`  

ERD diagram available at `/docs/erd.png` (optional)

---

## Authentication

All endpoints are protected by JWT tokens. Available roles:

- `employee`: approve loans
- `investor`: make investments
- `borrower`: create loan requests

---

## API Endpoints (Contoh)

### Create Loan
POST /v1/loan/create
curl --location 'http://localhost:8000/v1/loan/create' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9' \
--form 'loan="{\"principal_amount\":1000000,\"rate\":10,\"roi\":15}";type=application/json' \
--form 'agreement_letter_link=@"/home/afif/Gambar/Screenshots/Screenshot from 2025-04-23 06-13-27.png"'

### Loan Approved
POST /v1/loan/approved
curl --location 'http://localhost:8000/v1/loan/approved' \
--header 'Authorization: ••••••' \
--form 'loan_approved="{\"loan_id\" : \"f6ccd486-5481-44a8-9643-50b6ca1c4542\"}"' \
--form 'proof_picture_url=@"/home/afif/Gambar/Screenshots/Screenshot from 2025-04-23 06-12-12.png"'

### Loan Investment
POST /v1/loan/invested
curl --location 'http://localhost:8000/v1/loan/invested' \
--header 'Authorization: ••••••' \
--form 'loan_investment="{\"loan_id\" : \"f6ccd486-5481-44a8-9643-50b6ca1c4542\", \"amount\":700000}"'

### Loan Disbursed
curl --location 'http://localhost:8000/v1/loan/disbursed' \
--header 'Authorization: ••••••' \
--form 'loan_disbursement="{\"loan_id\" : \"f6ccd486-5481-44a8-9643-50b6ca1c4542\"}"' \
--form 'signed_agreement_url=@"/home/afif/Gambar/Screenshots/Screenshot from 2025-04-23 06-07-52.png"'




## Screenshoot Open API Swagger