
 ```
   ____ _               _      _____                 _ _   ____                     __ 
  / ___| |__   ___  ___| | __ | ____|_ __ ___   __ _(_) | / ___| _ __   ___   ___  / _|
 | |   | '_ \ / _ \/ __| |/ / |  _| | '_ ` _ \ / _` | | | \___ \| '_ \ / _ \ / _ \| |_ 
 | |___| | | |  __/ (__|   <  | |___| | | | | | (_| | | |  ___) | |_) | (_) | (_) |  _|
  \____|_| |_|\___|\___|_|\_\ |_____|_| |_| |_|\__,_|_|_| |____/| .__/ \___/ \___/|_|  
                                                                |_|                    
 ```

The objective of this project is to automate email spoof vulnerability checking.
The app checks the DNS txt records of the host entered to make the decision.
[OWASP top ten - Misconfiguration](https://owasp.org/www-project-top-ten/2017/A6_2017-Security_Misconfiguration)

## Disclaimer

```
I am not responsible for the misuse of the project.
I want to make the internet safer, so I posted it here to help 
other developers find this vulnerability and fix it.

USE WITH RESPONSABILITY!
```

## Dependencies

- [Go lang](https://golang.org/dl/)


## 1 - Run project

```
  $ go run src/main.go
```


## 1.2 - Choose option

```
############ ENTER OPTION ############
[1] - Check vulnerability an host
[2] - Make send test email
[0] - Exit
########################################
```

## 2 - Option [1]:

```
Enter to host:
example: example.com
```

result:
```diff
+ The host be safe.
or
! The host IS VULNERABLE!
```

## 2.1 - Option [2]:

#### Fake email
```
Enter the fake email you would like to test.

Enter to fake email to test:
example: admin@example.com
```
#### From email
```
Enter the email that will be used to send.
Here you need a valid email.

Warning: Use your private and not public email: @gmail, @yahoo, @hotmail, @live and others.
These big public servers won't let you run the spoof process.

Enter to from email:
example: senderemail@mydomain.com
```
#### Password
```
Enter the password from the email you entered in the previous step.

Enter to from password:
example: myP@ssw0rd
```
#### SMTP Server
```
Enter your SMTP or mail sending server with the server.

Enter to from smtp server:
example: smtp.mydomain.com
```
#### Port
```
Enter to from port:
example: 465
```
#### Receive email
```
Enter the email you want to receive the message.

Enter to email:
example: my-email-receiver@gmail.com
```

result:
```diff
+ Email Sent Successfully!.
or
! Error
```

## I found a vulnerable host, now what?
- Report!
- To fix the fault, just change the spf registry: `~all` or `?all` to `-all`

## Roadmap
- Use docker

## Author
- Harrisson Ricardo Biaggio
