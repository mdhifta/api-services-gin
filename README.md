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
