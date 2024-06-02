-- Create the database
CREATE DATABASE book_collections;

-- Switch to the new database
USE book_collections;

-- Create the books table
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    isbn VARCHAR(13) NOT NULL UNIQUE,
    published_date DATE NOT NULL
);

-- Insert dummy data into the books table
INSERT INTO books (title, author, isbn, published_date) VALUES
('The Great Gatsby', 'F. Scott Fitzgerald', '9780743273565', '1925-04-10'),
('To Kill a Mockingbird', 'Harper Lee', '9780061120084', '1960-07-11'),
('1984', 'George Orwell', '9780451524935', '1949-06-08'),
('Pride and Prejudice', 'Jane Austen', '9780141439518', '1813-01-28'),
('Moby-Dick', 'Herman Melville', '9781503280786', '1851-11-14');
