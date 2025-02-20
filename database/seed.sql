-- Insert Users
INSERT INTO users (email, password) VALUES
('user1@example.com', 'password1'),
('user2@example.com', 'password2'),
('user3@example.com', 'password3'),
('user4@example.com', 'password4'),
('user5@example.com', 'password5');

-- Insert Clients
INSERT INTO clients (first_name, last_name, address, phone, email) VALUES
('John', 'Doe', '123 Main St', '123-456-7890', 'john.doe@example.com'),
('Jane', 'Smith', '456 Elm St', '987-654-3210', 'jane.smith@example.com'),
('Alice', 'Johnson', '789 Oak St', '555-555-5555', 'alice.johnson@example.com'),
('Bob', 'Brown', '321 Pine St', '222-333-4444', 'bob.brown@example.com'),
('Carol', 'Davis', '654 Maple St', '666-777-8888', 'carol.davis@example.com');

-- Insert Pets
INSERT INTO pets (name, species, breed, date_of_birth, client_id) VALUES
('Buddy', 'Dog', 'Golden Retriever', '2020-01-01', 1),
('Mittens', 'Cat', 'Siamese', '2019-05-15', 2),
('Charlie', 'Dog', 'Labrador', '2018-08-20', 3),
('Max', 'Dog', 'Beagle', '2021-03-10', 4),
('Bella', 'Cat', 'Persian', '2020-07-22', 5);

-- Insert Veterinarians
INSERT INTO veterinarians (first_name, last_name, specialty, phone, email) VALUES
('Dr. Emily', 'Brown', 'Surgery', '111-222-3333', 'emily.brown@example.com'),
('Dr. Michael', 'Green', 'Dentistry', '444-555-6666', 'michael.green@example.com'),
('Dr. Sarah', 'White', 'Dermatology', '777-888-9999', 'sarah.white@example.com'),
('Dr. John', 'Black', 'Cardiology', '999-000-1111', 'john.black@example.com'),
('Dr. Lisa', 'Blue', 'Neurology', '222-444-6666', 'lisa.blue@example.com');

-- Insert Appointments
INSERT INTO appointments (date, time, pet_id, veterinarian_id, reason_for_appointment) VALUES
('2023-10-10', '2023-10-10T10:00:00+00', 1, 1, 'Checkup'),
('2023-10-11', '2023-10-10T11:00:00+00', 2, 2, 'Vaccination'),
('2023-10-12', '2023-10-10T12:00:00+00', 3, 3, 'Surgery'),
('2023-10-13', '2023-10-10T13:00:00+00', 4, 4, 'Heart Checkup'),
('2023-10-14', '2023-10-10T14:00:00+00', 5, 5, 'Neurological Exam');

-- Insert Treatments
INSERT INTO treatments (name, description, cost) VALUES
('Vaccination', 'Routine vaccination', 50.0),
('Dental Cleaning', 'Teeth cleaning', 100.0),
('Surgery', 'Minor surgery', 200.0),
('Heart Checkup', 'Cardiac examination', 150.0),
('Neurological Exam', 'Brain and nerve examination', 180.0);

-- Insert Invoices
INSERT INTO invoices (date, total, client_id, appointment_id) VALUES
('2023-10-10', 150.0, 1, 1),
('2023-10-11', 100.0, 2, 2),
('2023-10-12', 200.0, 3, 3),
('2023-10-13', 150.0, 4, 4),
('2023-10-14', 180.0, 5, 5);

-- Insert Medications
INSERT INTO medications (name, description, price) VALUES
('Antibiotic', '2 pills per day', 20.0),
('Painkiller', '1 pill per day', 15.0),
('Vitamin', '1 pill per day', 10.0),
('Antidepressant', '1 pill per day', 25.0),
('Antifungal', '2 pills per day', 30.0);
