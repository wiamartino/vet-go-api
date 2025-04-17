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
