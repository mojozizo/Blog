# Blog Website
Basic Blog website created using Go, PostgreSQL and React

## Things to DO :
- Add functionality to User Blog Creation/Updation
- Add permission logic for user access
- Validation and Error Handling
-  - Ensure that user inputs are validated before processing
-  - Improve error responses with proper HTTP status codes and clear error messages.
-  Authentication and Authorization
-  - Implement user login, signup, and JWT (JSON Web Token) authentication.
-  - Ensure that only authorized users can perform certain actions, like blog post creation or deletion.
-  Pagination for Data Fetching
-  Deployment
-  - AWS/HEroku and Containerization using Docker

## Things Already Done:
- Create routes for Blog Creation and User Creation
- Create models for Blog and User


## APIs
### User Management
http://localhost:3000/user/create - User Creation
http://localhost:3000/user/fetch/{user_id} - Fetch User
http://localhost:3000/user/update/{user_id} - Update User
http://localhost:3000/user/delete/{user_id} - Delete User

