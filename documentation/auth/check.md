#### Auth - Check 
  API check if user has session active and if successfully logged.
  
## Setup
  * Method: `GET`
  * URL: /api/_{version}_/auth/check
  
## Responses
  * **Code:** `404 NOT FOUND` <br />
  
    ```json
    {
      "success": false,
      "error":   "User not found."
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
    
