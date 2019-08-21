#### Auth - Register 
  API register user on mySQL database with token.
  
## Setup
  * Method: `POST`
  * URL: /api/_{version}_/auth/register

**Body Parameters**
  * Required
      * Username `username=[string]`
      * Name `display_name=[string]`
      * Password `password=[string][hash]`
  
## Responses
  * **Code:** `400 BAD REQUEST` <br />
  
    ```json
    {
      "success": false,
      "error":   "Complete all required fields."
    }
    ```
    
  * **Code:** `409 CONFLIT` <br />
  
    ```json
    {
      "success": false,
      "error":   "Username already exists."
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