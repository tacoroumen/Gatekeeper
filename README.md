# Go REST API for User Management

This is a simple Go application that provides a REST API for user management. It includes functionality for adding, modifying, deleting, and retrieving user information, as well as handling user reservations based on license plate.

## Prerequisites

- Go installed on your machine
- MySQL server running
- `config.json` file in the `config` directory with the necessary database connection details

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/your-repo.git
   ```

2. Navigate to the project directory:

   ```bash
   cd your-repo
   ```

3. Install dependencies:

   ```bash
   go get -d ./...
   ```

4. Build the application:

   ```bash
   go build
   ```

5. Run the application:

   ```bash
   ./your-repo
   ```

## API Endpoints

### 1. License Plate Information

- **Endpoint**: `/licenseplate`
- **Method**: `GET`
- **Parameters**:
  - `licenseplate`: License plate of the vehicle
- **Response**:
  - Returns user information associated with the provided license plate.

### 2. Reservation Management

- **Endpoint**: `/reservering`
- **Methods**:
  - `GET`: Method not allowed
  - `POST`: Create a new reservation
- **Parameters**:
  - `checkin`: Check-in date
  - `checkout`: Check-out date
  - `housenumber`: House number
  - `email`: User email
  - `password`: User password
- **Response**:
  - Returns a status message indicating the success or failure of the reservation creation.

### 3. User Management

#### 3.1 Add User

- **Endpoint**: `/user/add`
- **Methods**:
  - `GET`: Method not allowed
  - `POST`: Create a new user
- **Parameters**:
  - Various user details including `firstname`, `lastname`, `email`, `password`, `phonenumber`, `postalcode`, `housenumber`, `street`, `town`, `country`, `birthdate`, `licenseplate`
- **Response**:
  - Returns a status message indicating the success or failure of the user creation.

#### 3.2 Modify User

- **Endpoint**: `/user/modify`
- **Methods**:
  - `GET`: Method not allowed
  - `POST`: Modify user details
- **Parameters**:
  - Various user details including `firstname`, `lastname`, `birthdate`, `town`, `email`, `oldpassword`, `newpassword`, `phonenumber`, `licenseplate`
- **Response**:
  - Returns a status message indicating the success or failure of the user modification.

#### 3.3 Delete User

- **Endpoint**: `/user/delete`
- **Methods**:
  - `GET`: Method not allowed
  - `POST`: Delete a user
- **Parameters**:
  - `email`: User email
  - `password`: User password
- **Response**:
  - Returns a status message indicating the success or failure of the user deletion.

#### 3.4 Get User Information

- **Endpoint**: `/user/get`
- **Method**: `GET`
- **Parameters**:
  - `email`: User email
  - `password`: User password
- **Response**:
  - Returns user information associated with the provided email and password.

### 4. User Login

- **Endpoint**: `/login`
- **Method**: `GET`
- **Parameters**:
  - `email`: User email
  - `password`: User password
- **Response**:
  - Returns a status message indicating the success or failure of the login attempt.

## Configuration

Ensure that the `config.json` file in the `config` directory contains the correct database connection details.

```json
{
  "username": "your_username",
  "password": "your_password",
  "ip": "your_database_ip",
  "port": "your_database_port",
  "database": "your_database_name"
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.