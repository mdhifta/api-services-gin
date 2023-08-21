<img src="https://github-production-user-asset-6210df.s3.amazonaws.com/55729354/257034907-22a20df7-2173-4d1b-9e0a-5d2e43b1b003.png">

# Clean Api Services Template
this is gin structure template for api or website. Easy to use and easy to understand, I will continue to develop it as long as I learn everything.

# ENV
You have to make .env settings according to your needs.
- SECERET= used for you to do integration with jwt (for API auth)
- KEY_SECERET= used to create a password hash or encrypt and decrypt data

# EMAIL SENDER
first you have to use the email settings in .env and then use the feature in the file called sendmail.go in the utils folder
- Send Email
  utils.SendEmail("your_subject", "your_message", []string{"sendto@gmail.com"})
  Value:
  * subject
  * message
  * send to

    Enter the desired object, then enter the desired message and enter the email you want to go to
    
- Send Email with HTML
  utils.SendEmailHTML("your_subject", "./(path_html)", "https://urls.com", []string{"sendto@gmail.com"})
  Value:
  * subject
  * path html
  * url/link
  * send to
  
  Enter the desired subject, then the html path is the location where your html file is stored (put it in the template folder), the link is a url that you can use to direct a verification or the like, the email you want to send

  # HASH (ENCRYPT-DECRYPT)
  enter the key_secret data in the .env according to your wishes
  
  - Encrypt
    * utils.HashPassword("your_password")
      - insert your_password string
        
  - Validation 
    * utils.CheckPasswordHash("your_password", "password_encrypt")
      - insert your_password string
      - insert password_encrypt is password your encrypt before


  # REQUEST CONSUME API
  enter the api_url (base_url) dan api_key (opsional) data in the .env according to your wishes
  
  - Request GET
    * utils.RequestGET("sub-url")
      - insert your sub url 
      - for example in this code with mangadex api
        
  - Request POST 
    * utils.RequestPOST("sub-url")
      - insert your sub url 
      - you can modify code at utils -> requestAPI.go
      - add data post to example RequestPOST(url string, json string)
      - user with new input data. to example RequestPOST("manga/", `{ "name": "myname" }`)
      - and change your jsonPost to jsonPost := json