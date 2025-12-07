-- Insert Users
INSERT INTO users (email, password) VALUES
('user1@example.com', 'password1'),
('user2@example.com', 'password2'),
('user3@example.com', 'password3'),
('user4@example.com', 'password4'),
('user5@example.com', 'password5'),
('admin@vetclinic.com', 'admin123'),
('doctor@vetclinic.com', 'doctor123'),
('reception@vetclinic.com', 'reception123'),
('tech@vetclinic.com', 'tech123'),
('manager@vetclinic.com', 'manager123');

-- Insert Clients
INSERT INTO clients (first_name, last_name, address, phone, email) VALUES
('John', 'Doe', '123 Main St', '123-456-7890', 'john.doe@example.com'),
('Jane', 'Smith', '456 Elm St', '987-654-3210', 'jane.smith@example.com'),
('Alice', 'Johnson', '789 Oak St', '555-555-5555', 'alice.johnson@example.com'),
('Bob', 'Brown', '321 Pine St', '222-333-4444', 'bob.brown@example.com'),
('Carol', 'Davis', '654 Maple St', '666-777-8888', 'carol.davis@example.com'),
('David', 'Miller', '789 Spruce Ave', '111-222-3333', 'david.miller@example.com'),
('Emma', 'Wilson', '456 Birch Blvd', '444-555-6666', 'emma.wilson@example.com'),
('Frank', 'Moore', '123 Cedar Lane', '777-888-9999', 'frank.moore@example.com'),
('Grace', 'Taylor', '987 Redwood Rd', '012-345-6789', 'grace.taylor@example.com'),
('Henry', 'Anderson', '654 Walnut Dr', '987-654-3210', 'henry.anderson@example.com');

-- Insert Veterinarians
INSERT INTO veterinarians (first_name, last_name, specialty, phone, email) VALUES
('Dr. Emily', 'Brown', 'Surgery', '111-222-3333', 'emily.brown@example.com'),
('Dr. Michael', 'Green', 'Dentistry', '444-555-6666', 'michael.green@example.com'),
('Dr. Sarah', 'White', 'Dermatology', '777-888-9999', 'sarah.white@example.com'),
('Dr. John', 'Black', 'Cardiology', '999-000-1111', 'john.black@example.com'),
('Dr. Lisa', 'Blue', 'Neurology', '222-444-6666', 'lisa.blue@example.com'),
('Dr. Thomas', 'Gray', 'Orthopedics', '333-444-5555', 'thomas.gray@example.com'),
('Dr. Jessica', 'Gold', 'Ophthalmology', '666-777-8888', 'jessica.gold@example.com'),
('Dr. William', 'Silver', 'Internal Medicine', '111-333-5555', 'william.silver@example.com'),
('Dr. Olivia', 'Rose', 'Oncology', '222-555-8888', 'olivia.rose@example.com'),
('Dr. James', 'Stone', 'Emergency Medicine', '444-666-9999', 'james.stone@example.com');



-- Insert Pets
INSERT INTO pets (name, species, breed, date_of_birth, client_id) VALUES
('Buddy', 'Dog', 'Golden Retriever', '2020-01-01', 1),
('Mittens', 'Cat', 'Siamese', '2019-05-15', 2),
('Charlie', 'Dog', 'Labrador', '2018-08-20', 3),
('Max', 'Dog', 'Beagle', '2021-03-10', 4),
('Bella', 'Cat', 'Persian', '2020-07-22', 5),
('Luna', 'Cat', 'Maine Coon', '2019-11-05', 6),
('Rex', 'Dog', 'German Shepherd', '2021-02-14', 7),
('Oliver', 'Cat', 'Ragdoll', '2020-09-30', 8),
('Lucy', 'Dog', 'Poodle', '2018-12-25', 9),
('Daisy', 'Dog', 'Shih Tzu', '2022-01-15', 10),
('Rocky', 'Dog', 'Boxer', '2021-06-22', 1),
('Chloe', 'Cat', 'Bengal', '2019-08-11', 2),
('Cooper', 'Dog', 'Husky', '2020-04-19', 3),
('Simba', 'Cat', 'Sphynx', '2021-10-31', 4),
('Lily', 'Dog', 'Dachshund', '2019-12-03', 5);

-- Insert Allergies
INSERT INTO allergies (pet_id, allergen, allergy_type, severity, reaction, diagnosed_date, diagnosed_by, notes, is_active, created_at, updated_at) VALUES
(1, 'Chicken', 'food', 'moderate', 'Skin rashes, itching', '2023-05-10', 1, 'Avoid all chicken-based foods and treats', TRUE, NOW(), NOW()),
(2, 'Penicillin', 'medication', 'severe', 'Anaphylaxis', '2023-03-15', 2, 'NEVER administer penicillin or related antibiotics. Use alternatives.', TRUE, NOW(), NOW()),
(3, 'Pollen (grass)', 'environment', 'mild', 'Sneezing, watery eyes', '2023-04-20', 3, 'Seasonal allergies worse in spring. Antihistamines as needed.', TRUE, NOW(), NOW()),
(4, 'Bee stings', 'insect', 'severe', 'Swelling, difficulty breathing', '2023-06-01', 4, 'Keep EpiPen available. Previous anaphylactic reaction.', TRUE, NOW(), NOW()),
(5, 'Dairy products', 'food', 'mild', 'Gastrointestinal upset, diarrhea', '2023-07-12', 1, 'Avoid milk, cheese, and dairy-based treats', TRUE, NOW(), NOW()),
(6, 'Flea bites', 'insect', 'moderate', 'Severe itching, hair loss, hot spots', '2023-08-05', 2, 'Flea allergy dermatitis. Strict flea control essential.', TRUE, NOW(), NOW()),
(7, 'Beef', 'food', 'moderate', 'Chronic ear infections, skin inflammation', '2023-02-28', 3, 'Switched to fish-based diet with improvement', TRUE, NOW(), NOW()),
(8, 'Dust mites', 'environment', 'mild', 'Mild skin irritation', '2023-09-10', 4, 'Keep environment clean, wash bedding frequently', TRUE, NOW(), NOW()),
(9, 'Sulfonamides', 'medication', 'severe', 'Vomiting, seizures', '2023-01-20', 1, 'Life-threatening reaction. Avoid all sulfa drugs.', TRUE, NOW(), NOW()),
(10, 'Corn', 'food', 'mild', 'Mild itching', '2023-05-30', 2, 'Use corn-free diet. Symptoms resolved.', TRUE, NOW(), NOW()),
(3, 'Peanut butter', 'food', 'moderate', 'Facial swelling, hives', '2023-10-12', 3, 'Developed allergy recently. No peanut products.', TRUE, NOW(), NOW()),
(7, 'Mold spores', 'environment', 'mild', 'Respiratory irritation', '2023-09-15', 3, 'Keep environment dry, use dehumidifier', TRUE, NOW(), NOW());

-- Insert Appointments
INSERT INTO appointments (date, time, pet_id, veterinarian_id, reason_for_appointment) VALUES
('2023-10-10', '2023-10-10T10:00:00+00', 1, 1, 'Checkup'),
('2023-10-11', '2023-10-10T11:00:00+00', 2, 2, 'Vaccination'),
('2023-10-12', '2023-10-10T12:00:00+00', 3, 3, 'Surgery'),
('2023-10-13', '2023-10-10T13:00:00+00', 4, 4, 'Heart Checkup'),
('2023-10-14', '2023-10-10T14:00:00+00', 5, 5, 'Neurological Exam'),
('2023-10-15', '2023-10-15T09:30:00+00', 6, 6, 'Joint Pain'),
('2023-10-16', '2023-10-16T13:15:00+00', 7, 7, 'Eye Infection'),
('2023-10-17', '2023-10-17T11:45:00+00', 8, 8, 'Stomach Issues'),
('2023-10-18', '2023-10-18T15:00:00+00', 9, 9, 'Cancer Screening'),
('2023-10-19', '2023-10-19T10:30:00+00', 10, 10, 'Emergency Visit'),
('2023-10-20', '2023-10-20T14:30:00+00', 11, 1, 'Annual Checkup'),
('2023-10-21', '2023-10-21T09:00:00+00', 12, 2, 'Dental Cleaning'),
('2023-10-22', '2023-10-22T16:15:00+00', 13, 3, 'Skin Rash'),
('2023-10-23', '2023-10-23T11:00:00+00', 14, 4, 'Heart Murmur'),
('2023-10-24', '2023-10-24T10:45:00+00', 15, 5, 'Seizure Evaluation');

-- Insert Treatments
INSERT INTO treatments (name, description, cost) VALUES
('Vaccination', 'Routine vaccination', 50.0),
('Dental Cleaning', 'Teeth cleaning', 100.0),
('Surgery', 'Minor surgery', 200.0),
('Heart Checkup', 'Cardiac examination', 150.0),
('Neurological Exam', 'Brain and nerve examination', 180.0),
('Orthopedic Treatment', 'Joint and bone treatment', 220.0),
('Ophthalmology Exam', 'Eye examination and treatment', 120.0),
('Gastrointestinal Treatment', 'Treatment for digestive issues', 160.0),
('Cancer Screening', 'Screening for various cancers', 250.0),
('Emergency Treatment', 'Urgent care and stabilization', 300.0),
('Annual Wellness Exam', 'Complete yearly health check', 90.0),
('X-ray', 'Diagnostic imaging', 140.0),
('Ultrasound', 'Soft tissue examination', 175.0),
('Blood Work', 'Complete blood analysis', 95.0),
('Microchipping', 'Pet identification implant', 45.0);

-- Insert Invoices
INSERT INTO invoices (date, total, client_id, appointment_id) VALUES
('2023-10-10', 150.0, 1, 1),
('2023-10-11', 100.0, 2, 2),
('2023-10-12', 200.0, 3, 3),
('2023-10-13', 150.0, 4, 4),
('2023-10-14', 180.0, 5, 5),
('2023-10-15', 220.0, 6, 6),
('2023-10-16', 120.0, 7, 7),
('2023-10-17', 160.0, 8, 8),
('2023-10-18', 250.0, 9, 9),
('2023-10-19', 300.0, 10, 10),
('2023-10-20', 90.0, 1, 11),
('2023-10-21', 100.0, 2, 12),
('2023-10-22', 175.0, 3, 13),
('2023-10-23', 150.0, 4, 14),
('2023-10-24', 180.0, 5, 15);

-- Insert Medications
INSERT INTO medications (name, description, price) VALUES
('Antibiotic', '2 pills per day', 20.0),
('Painkiller', '1 pill per day', 15.0),
('Vitamin', '1 pill per day', 10.0),
('Antidepressant', '1 pill per day', 25.0),
('Antifungal', '2 pills per day', 30.0),
('Heartworm Prevention', 'Monthly chewable', 35.0),
('Flea & Tick Treatment', 'Monthly application', 40.0),
('Anti-inflammatory', 'Once daily with food', 22.0),
('Allergy Medicine', 'As needed for allergic reactions', 18.0),
('Insulin', 'Twice daily injection', 65.0),
('Ear Drops', '3 drops twice daily', 15.0),
('Eye Ointment', 'Apply thin layer twice daily', 18.0),
('Joint Supplement', 'Once daily with food', 25.0),
('Probiotic', 'Once daily with food', 30.0),
('Dewormer', 'Single dose treatment', 12.0);

-- Insert Medical Records
INSERT INTO medical_records (pet_id, veterinarian_id, appointment_id, visit_date, diagnosis, symptoms, notes, weight, temperature, heart_rate, created_at, updated_at) VALUES
(1, 1, 1, '2023-10-10 09:00:00', 'Ear infection', 'Scratching ears, head shaking', 'Prescribed antibiotics and ear drops. Follow-up in 2 weeks.', 32.5, 38.5, 90, NOW(), NOW()),
(2, 2, 2, '2023-10-11 10:30:00', 'Dental cleaning needed', 'Bad breath, tartar buildup', 'Scheduled dental cleaning procedure. Pre-op bloodwork required.', 4.2, 38.2, 140, NOW(), NOW()),
(3, 3, 3, '2023-10-12 14:00:00', 'Skin allergies', 'Itching, red skin patches', 'Allergy testing recommended. Started on antihistamines.', 28.0, 38.6, 85, NOW(), NOW()),
(4, 4, 4, '2023-10-13 11:15:00', 'Heart murmur detected', 'No symptoms, routine checkup', 'Grade 2/6 heart murmur. Recommend cardiac ultrasound.', 12.3, 38.3, 120, NOW(), NOW()),
(5, 1, 5, '2023-10-14 15:30:00', 'Upper respiratory infection', 'Sneezing, nasal discharge', 'Prescribed antibiotics. Keep isolated from other pets.', 5.1, 39.0, 150, NOW(), NOW()),
(6, 2, 6, '2023-10-15 09:45:00', 'Healthy checkup', 'None', 'Annual wellness exam. All vitals normal. Vaccines updated.', 8.5, 38.4, 130, NOW(), NOW()),
(7, 3, 7, '2023-10-16 13:20:00', 'Hip dysplasia', 'Limping, difficulty standing', 'X-rays confirm hip dysplasia. Started on joint supplements and pain management.', 35.0, 38.7, 95, NOW(), NOW()),
(8, 4, 8, '2023-10-17 10:00:00', 'Conjunctivitis', 'Red, watery eyes', 'Prescribed eye ointment. Apply twice daily for 7 days.', 6.8, 38.3, 145, NOW(), NOW()),
(9, 1, 9, '2023-10-18 16:00:00', 'Gastrointestinal upset', 'Vomiting, diarrhea', 'Prescribed bland diet and probiotics. Monitor hydration.', 8.0, 38.9, 110, NOW(), NOW()),
(10, 2, 10, '2023-10-19 11:30:00', 'Obesity consultation', 'Overweight', 'Weight loss plan created. Reduce food portions and increase exercise.', 15.5, 38.5, 100, NOW(), NOW());

-- Insert Vaccinations
INSERT INTO vaccinations (pet_id, veterinarian_id, vaccine_name, manufacturer, batch_number, date_administered, date_scheduled, next_due_date, status, notes, side_effects, created_at, updated_at) VALUES
(1, 1, 'Rabies', 'Zoetis', 'RAB123456', '2023-10-10 09:30:00', NULL, '2024-10-10', 'completed', 'Annual rabies vaccination', 'None reported', NOW(), NOW()),
(1, 1, 'DHPP (Distemper, Hepatitis, Parvovirus, Parainfluenza)', 'Merck', 'DHPP789012', '2023-10-10 09:35:00', NULL, '2024-10-10', 'completed', 'Core vaccine for dogs', 'None reported', NOW(), NOW()),
(2, 2, 'FVRCP (Feline Viral Rhinotracheitis, Calicivirus, Panleukopenia)', 'Boehringer Ingelheim', 'FVRCP345678', '2023-10-11 10:45:00', NULL, '2024-10-11', 'completed', 'Core vaccine for cats', 'None reported', NOW(), NOW()),
(2, 2, 'Rabies', 'Zoetis', 'RAB234567', '2023-10-11 10:50:00', NULL, '2024-10-11', 'completed', 'Annual rabies vaccination', 'None reported', NOW(), NOW()),
(3, 3, 'Leptospirosis', 'Zoetis', 'LEPTO901234', '2023-10-12 14:15:00', NULL, '2024-10-12', 'completed', 'Protects against bacterial infection', 'Mild lethargy for 24 hours', NOW(), NOW()),
(4, 4, 'Bordetella (Kennel Cough)', 'Merck', 'BORD567890', NULL, '2023-11-13 11:00:00', '2024-11-13', 'scheduled', 'Scheduled for boarding', '', NOW(), NOW()),
(5, NULL, 'Rabies', 'Zoetis', NULL, NULL, '2023-11-20 14:00:00', '2024-11-20', 'scheduled', 'Upcoming annual rabies', '', NOW(), NOW()),
(6, 2, 'FeLV (Feline Leukemia)', 'Boehringer Ingelheim', 'FELV123456', '2023-10-15 10:00:00', NULL, '2024-10-15', 'completed', 'Recommended for outdoor cats', 'None reported', NOW(), NOW()),
(7, 3, 'Lyme Disease', 'Merck', 'LYME789012', '2023-10-16 13:30:00', NULL, '2024-10-16', 'completed', 'High tick area protection', 'None reported', NOW(), NOW()),
(8, NULL, 'FVRCP', 'Boehringer Ingelheim', NULL, NULL, '2023-12-01 10:00:00', '2024-12-01', 'scheduled', 'Booster due', '', NOW(), NOW()),
(9, NULL, 'Rabies', 'Zoetis', NULL, NULL, '2023-10-15 16:00:00', '2024-10-15', 'overdue', 'Overdue for annual rabies', '', NOW(), NOW()),
(10, 2, 'Canine Influenza', 'Zoetis', 'CINFL345678', '2023-10-19 11:45:00', NULL, '2024-10-19', 'completed', 'H3N2 and H3N8 strains', 'None reported', NOW(), NOW());

-- Insert Surgeries
INSERT INTO surgeries (pet_id, veterinarian_id, surgery_name, surgery_type, status, scheduled_date, actual_date, duration, pre_op_notes, post_op_notes, complications, anesthesia_used, follow_up_required, follow_up_date, cost, created_at, updated_at) VALUES
(1, 1, 'Ear canal ablation', 'routine', 'completed', '2023-09-15 09:00:00', '2023-09-15 09:00:00', 90, 'Pre-op bloodwork normal. NPO after midnight.', 'Surgery successful. Patient recovering well.', 'None', 'Isoflurane', TRUE, '2023-09-22', 850.00, NOW(), NOW()),
(2, 2, 'Dental extraction', 'routine', 'completed', '2023-09-20 10:00:00', '2023-09-20 10:00:00', 60, 'Multiple decayed teeth identified.', 'Extracted 3 teeth. Prescribed pain medication.', 'Minor bleeding controlled', 'Sevoflurane', TRUE, '2023-09-27', 450.00, NOW(), NOW()),
(3, 1, 'Cruciate ligament repair (TPLO)', 'elective', 'completed', '2023-08-10 08:00:00', '2023-08-10 08:00:00', 150, 'Left hind leg affected. X-rays confirmed.', 'TPLO performed successfully. Strict rest for 8 weeks.', 'None', 'Isoflurane', TRUE, '2023-08-24', 2500.00, NOW(), NOW()),
(7, 3, 'Hip replacement surgery', 'elective', 'scheduled', '2023-11-05 08:00:00', NULL, NULL, 'Right hip severe dysplasia. Pre-op scheduled.', '', '', '', TRUE, NULL, 3500.00, NOW(), NOW()),
(4, 4, 'Emergency gastric torsion', 'emergency', 'completed', '2023-10-01 22:00:00', '2023-10-01 22:30:00', 120, 'Bloat emergency. Immediate surgery required.', 'Stomach derotated and gastropexy performed. Critical first 24 hours.', 'None', 'Isoflurane', TRUE, '2023-10-08', 1800.00, NOW(), NOW()),
(5, 1, 'Spay (ovariohysterectomy)', 'routine', 'completed', '2023-07-15 09:00:00', '2023-07-15 09:00:00', 45, 'Routine spay procedure.', 'Surgery completed without complications. E-collar placed.', 'None', 'Sevoflurane', TRUE, '2023-07-22', 350.00, NOW(), NOW()),
(6, 2, 'Mass removal - lipoma', 'elective', 'scheduled', '2023-11-15 10:00:00', NULL, NULL, 'Large benign lipoma on chest. Not urgent but recommended.', '', '', '', TRUE, NULL, 650.00, NOW(), NOW()),
(8, 1, 'Enucleation (eye removal)', 'routine', 'completed', '2023-09-05 14:00:00', '2023-09-05 14:00:00', 75, 'Chronic glaucoma unresponsive to treatment.', 'Left eye removed. Pain management protocol initiated.', 'None', 'Isoflurane', TRUE, '2023-09-12', 950.00, NOW(), NOW()),
(9, 1, 'Pyometra surgery', 'emergency', 'completed', '2023-08-20 16:00:00', '2023-08-20 16:30:00', 90, 'Emergency uterine infection. Critical condition.', 'Emergency spay performed. IV antibiotics started.', 'Mild post-op infection treated', 'Isoflurane', TRUE, '2023-08-27', 1200.00, NOW(), NOW()),
(10, 2, 'Neuter (castration)', 'routine', 'completed', '2023-06-01 09:00:00', '2023-06-01 09:00:00', 30, 'Routine neuter procedure.', 'Surgery successful. Minimal bleeding. E-collar placed.', 'None', 'Sevoflurane', TRUE, '2023-06-08', 250.00, NOW(), NOW());

