# Pet Medical History Feature

## Overview

Comprehensive medical history tracking system for pets including medical records, vaccinations, surgeries, and allergies.

## Features Implemented

### 1. Medical Records (`/api/v1/medical-records`)
Track detailed visit records for each pet including:
- Diagnosis and symptoms
- Vital signs (weight, temperature, heart rate)
- Veterinarian notes
- Link to appointments

**Endpoints:**
- `GET /api/v1/medical-records` - Get all medical records
- `GET /api/v1/medical-records/:id` - Get specific medical record
- `GET /api/v1/pets/:pet_id/medical-records` - Get all records for a pet
- `POST /api/v1/medical-records` - Create new medical record
- `PUT /api/v1/medical-records/:id` - Update medical record
- `DELETE /api/v1/medical-records/:id` - Delete medical record

**Example Request (Create Medical Record):**
```json
{
  "pet_id": 1,
  "veterinarian_id": 1,
  "appointment_id": 1,
  "visit_date": "2023-12-02T10:00:00Z",
  "diagnosis": "Ear infection",
  "symptoms": "Scratching ears, head shaking",
  "notes": "Prescribed antibiotics",
  "weight": 32.5,
  "temperature": 38.5,
  "heart_rate": 90
}
```

### 2. Vaccinations (`/api/v1/vaccinations`)
Manage vaccination schedules and records:
- Track vaccination history
- Schedule future vaccinations
- Monitor due and overdue vaccines
- Record side effects

**Status Types:**
- `scheduled` - Vaccination scheduled
- `completed` - Vaccination administered
- `overdue` - Past due date
- `cancelled` - Vaccination cancelled

**Endpoints:**
- `GET /api/v1/vaccinations` - Get all vaccinations
- `GET /api/v1/vaccinations/:id` - Get specific vaccination
- `GET /api/v1/pets/:pet_id/vaccinations` - Get pet vaccinations
- `GET /api/v1/vaccinations/due?days=30` - Get vaccinations due in next N days
- `GET /api/v1/vaccinations/overdue` - Get overdue vaccinations
- `POST /api/v1/vaccinations` - Schedule new vaccination
- `PUT /api/v1/vaccinations/:id` - Update vaccination
- `PATCH /api/v1/vaccinations/:id/complete` - Mark vaccination as completed
- `DELETE /api/v1/vaccinations/:id` - Delete vaccination

**Example Request (Schedule Vaccination):**
```json
{
  "pet_id": 1,
  "vaccine_name": "Rabies",
  "manufacturer": "Zoetis",
  "date_scheduled": "2024-01-15T10:00:00Z",
  "next_due_date": "2025-01-15",
  "status": "scheduled",
  "notes": "Annual rabies vaccination"
}
```

**Example Request (Complete Vaccination):**
```json
{
  "veterinarian_id": 1,
  "side_effects": "None reported"
}
```

### 3. Surgeries (`/api/v1/surgeries`)
Track surgical procedures with detailed pre/post-op notes:
- Schedule surgeries
- Track surgery status workflow
- Record anesthesia and complications
- Manage follow-up appointments

**Surgery Types:**
- `routine` - Standard procedures
- `emergency` - Emergency surgeries
- `elective` - Non-urgent procedures

**Surgery Status:**
- `scheduled` - Surgery scheduled
- `in_progress` - Surgery underway
- `completed` - Surgery finished
- `cancelled` - Surgery cancelled

**Endpoints:**
- `GET /api/v1/surgeries` - Get all surgeries
- `GET /api/v1/surgeries/:id` - Get specific surgery
- `GET /api/v1/pets/:pet_id/surgeries` - Get pet surgeries
- `GET /api/v1/surgeries/status?status=scheduled` - Filter by status
- `GET /api/v1/surgeries/scheduled?start_date=2023-12-01&end_date=2023-12-31` - Get surgeries in date range
- `POST /api/v1/surgeries` - Schedule new surgery
- `PUT /api/v1/surgeries/:id` - Update surgery
- `PATCH /api/v1/surgeries/:id/start` - Mark surgery as started
- `PATCH /api/v1/surgeries/:id/complete` - Mark surgery as completed
- `PATCH /api/v1/surgeries/:id/cancel` - Cancel surgery
- `DELETE /api/v1/surgeries/:id` - Delete surgery

**Example Request (Schedule Surgery):**
```json
{
  "pet_id": 1,
  "veterinarian_id": 1,
  "surgery_name": "Dental cleaning",
  "surgery_type": "routine",
  "scheduled_date": "2023-12-15T09:00:00Z",
  "pre_op_notes": "NPO after midnight",
  "anesthesia_used": "Isoflurane",
  "follow_up_required": true,
  "cost": 350.00
}
```

**Example Request (Complete Surgery):**
```json
{
  "post_op_notes": "Surgery successful. Patient recovering well.",
  "complications": "None"
}
```

### 4. Allergies (`/api/v1/allergies`)
Track pet allergies and sensitivities:
- Record allergen details
- Track severity levels
- Monitor active/inactive allergies
- Important for medication safety

**Allergy Types:**
- `food` - Food allergies
- `medication` - Drug allergies
- `environment` - Environmental allergies
- `insect` - Insect bite allergies
- `other` - Other allergies

**Severity Levels:**
- `mild` - Minor reactions
- `moderate` - Moderate symptoms
- `severe` - Serious reactions
- `fatal` - Life-threatening

**Endpoints:**
- `GET /api/v1/allergies` - Get all allergies
- `GET /api/v1/allergies/:id` - Get specific allergy
- `GET /api/v1/pets/:pet_id/allergies` - Get pet allergies
- `GET /api/v1/pets/:pet_id/allergies/active` - Get active allergies only
- `GET /api/v1/allergies/severity?severity=severe` - Filter by severity
- `POST /api/v1/allergies` - Record new allergy
- `PUT /api/v1/allergies/:id` - Update allergy
- `PATCH /api/v1/allergies/:id/deactivate` - Mark allergy as inactive
- `PATCH /api/v1/allergies/:id/reactivate` - Mark allergy as active
- `DELETE /api/v1/allergies/:id` - Delete allergy

**Example Request (Record Allergy):**
```json
{
  "pet_id": 1,
  "allergen": "Penicillin",
  "allergy_type": "medication",
  "severity": "severe",
  "reaction": "Anaphylaxis",
  "diagnosed_date": "2023-03-15T00:00:00Z",
  "diagnosed_by": 1,
  "notes": "NEVER administer penicillin or related antibiotics",
  "is_active": true
}
```

## Database Schema

### medical_records
- medical_record_id (PK)
- pet_id (FK)
- veterinarian_id (FK)
- appointment_id (FK, optional)
- visit_date
- diagnosis
- symptoms
- notes
- weight
- temperature
- heart_rate
- created_at, updated_at

### vaccinations
- vaccination_id (PK)
- pet_id (FK)
- veterinarian_id (FK, optional)
- vaccine_name
- manufacturer
- batch_number
- date_administered
- date_scheduled
- next_due_date
- status (enum)
- notes
- side_effects
- created_at, updated_at

### surgeries
- surgery_id (PK)
- pet_id (FK)
- veterinarian_id (FK)
- surgery_name
- surgery_type (enum)
- status (enum)
- scheduled_date
- actual_date
- duration
- pre_op_notes
- post_op_notes
- complications
- anesthesia_used
- follow_up_required
- follow_up_date
- cost
- created_at, updated_at

### allergies
- allergy_id (PK)
- pet_id (FK)
- allergen
- allergy_type (enum)
- severity (enum)
- reaction
- diagnosed_date
- diagnosed_by (FK)
- notes
- is_active
- created_at, updated_at

## Architecture

The implementation follows Clean Architecture principles:

1. **Domain Layer** (`domain/`): Entities and repository interfaces
2. **Infrastructure Layer** (`infrastructure/repositories/`): Database implementations
3. **Application Layer** (`application/`): Business logic services
4. **Controllers** (`controllers/`): HTTP request handlers
5. **Routes** (`routes/`): API endpoint definitions

## Sample Data

The system includes seed data with:
- 10 medical records with various diagnoses
- 12 vaccination records (scheduled, completed, overdue)
- 10 surgical procedures (routine, emergency, elective)
- 12 allergy records with different severities

## Testing

Run the application in development mode to automatically seed the database:
```bash
export GO_ENV=development
go run main.go
```

## Authentication

All endpoints require authentication via JWT token:
```
Authorization: Bearer <your-jwt-token>
```

## Use Cases

### 1. Complete Pet Medical History
```bash
# Get all medical information for a pet
GET /api/v1/pets/1/medical-records
GET /api/v1/pets/1/vaccinations
GET /api/v1/pets/1/surgeries
GET /api/v1/pets/1/allergies/active
```

### 2. Vaccination Reminder System
```bash
# Check vaccinations due in next 30 days
GET /api/v1/vaccinations/due?days=30

# Check overdue vaccinations
GET /api/v1/vaccinations/overdue
```

### 3. Surgery Schedule Management
```bash
# View surgeries for the week
GET /api/v1/surgeries/scheduled?start_date=2023-12-04&end_date=2023-12-10

# Filter by status
GET /api/v1/surgeries/status?status=scheduled
```

### 4. Critical Allergy Checks
```bash
# Get all severe/fatal allergies
GET /api/v1/allergies/severity?severity=severe
GET /api/v1/allergies/severity?severity=fatal

# Check active allergies before prescribing medication
GET /api/v1/pets/1/allergies/active
```

## Future Enhancements

- Prescription management
- Medical document uploads (X-rays, lab results)
- Automated vaccination reminders
- Drug interaction warnings
- Treatment plan tracking
- Insurance claim integration
