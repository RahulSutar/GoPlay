# The Book Shelf
Write a simple Book Shelf Web Application which should expose APIs for below operations.


## APIs

 1. Create/Add book with following information
    a. Book name
    b. Author
    c. Book Type (Historical, Fairytale, Horror etc...)
    d. Book Description
    e. Number of Copies - Number of copies should be calclated at runtime and should not be passed as a part of Create/Add API call. Adding book by same name should just increase number of copies by 1.
    f. Availability Status - Availability Status should be "unavailable" of number of copies are 0.
2. Add books in bulk from CSV file
    a. Each CSV can have "n" number of records.
    b. Each record represents a book with above mentioned information except number of copies and availability status.
    c. Put CSVs in a folder for the app to consume.
    d. There can be hundreds of CSV file presend in this folder.
    e. 10 CSV files should be process in parallel.
3. Update book's author, type and description
4. Delete a book based on its name
    a. Should just reduce number of copies.
5. Get book information based on its name
6. Get all books information

## Design quo's

1. Generate logs and write them in a log file.
2. One should be able to track the API request in the log using a unique ID for each API request present in each log line.
3. Proper ORM layer exposing CREATE, UPDATE, GET and DELETE opearations.
4. Unit testing implementation along with functional tests (which uses db and not mock data)
5. Use interfaces where ever applicable.
6. Keep the ORM implementation
    a. Abstract or loosely coupled from the service layer.
    b. Easily switch between databases


## Tools to be used

1. Golang version : 1.11 or later
2. Database : PostgresSQL
3. Loggin framework : Logrus
4. ORM : GORM
5. Dependency injection : dep
6. MUX router
7. net/http module