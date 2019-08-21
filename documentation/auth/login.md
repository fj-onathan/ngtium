#### Auth - Login 
  API make user logged with session.
  
## Setup
  * Method: `POST`
  * URL: /api/_{version}_/auth/login

**Body Parameters**
  * Required
      * Username `username=[string]`
      * Password `password=[string][hash]`
  
## Responses
  * **Code:** `400 BAD REQUEST` <br />
  
    ```json
    {
      "success": false,
      "error":   "Complete all required fields."
    }
    ```
    
  * **Code:** `404 NOT FOUND` <br />
  
    ```json
    {
      "success": false,
      "error":   "User not found."
    }
    ```
    
  * **Code:** `401 UNAUTHORIZED` <br />
  
    ```json
    {
      "success": false,
      "error":   "Requested data is invalid."
    }
    ```
    
  * **Code:** `200 SUCCESS` <br />
    
    ```json
    {
      "success": true,
      "token":   "token",
      "user":    "user.info"
    }
    ```