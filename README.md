# Rate-Limited Notification Service
![technology Go](https://img.shields.io/badge/technology-go-blue.svg)

## The problem
We have a Notification system that sends out email notifications of various types (status update, daily news, project invitations, etc). We need to protect recipients from getting too many emails, either due to system errors or due to abuse, so let’s limit the number of emails sent to them by implementing a rate-limited version of NotificationService.
The system must reject requests that are over the limit.
Some sample notification types and rate limit rules, e.g.:

Status: not more than 2 per minute for each recipient
News: not more than 1 per day for each recipient
Marketing: not more than 3 per hour for each recipient
Etc. these are just samples, the system might have several rate limit rules!

## The solution
Reading all the information, there are two conditions that could be applied to know if an email should be sent:
- The date of sending of the last email
- The count of notifications that has been sent to the user recipent in the interval of time
With this in mind, there are two validations, if the date of last notification is above the range, and if it's not the case validate if the count of notifications is less than the allowed.

Talking about the validations, there a struct called "Notification" which have the validation logic depending of the interval and limit fields, with this in mind there are 3 clases that "extends" this struct and override the interval and limit fields, called statusNotification, newsNotification and marketingNotification, each one overrides the fields as the problem description and examples.

### Architecture (Layers)
The hex arxhiture allows to segregate the aplication in 2 tiers, that can comunicate in one direction, with this in mind, the 2 tiers of this architecture are **infraestructure** and **domain**.

- **Infraestructure**: Has the responsability of interact directly with the outputs and inputs of the system also do the adaptors to the ports defined in the domain layer. Examples of this layer can be expose web services, connect to databases, read task queues, etc.
- **Domain**: Represents all the business logic of the application which is the reason for the business to exist. It try to avoid the [Anemic anti pattern](https://martinfowler.com/bliki/AnemicDomainModel.html) and suport the [Tell dont ask principle](https://martinfowler.com/bliki/TellDontAsk.html). In this layer you can find the following aggregate patterns, domain services, entities, value objects, repositories (port), etc.

To obtain more documentation on this type of architecture, it is recommended to read about [Hexagonal architecture and DDD](https://codely.tv/blog/screencasts/arquitectura-hexagonal-ddd/)

## Run
To run the app just type the following command
```
$ go run main.go
``` 

## Getting started with dependencies
Run the following commands to install cli dependencies:
- ``` $ sudo chmod +x ./scripts/*.sh ```
- ``` $ ./scripts/install_dependencies.sh  ```

## Testing
For testing are used three libraries:
- [Ginkgo](https://onsi.github.io/ginkgo/): for structure, before and after functions 
- [Gomega](https://onsi.github.io/gomega/): for spects structure and implementation
- [Gomock](https://github.com/golang/mock): for mocking thinks like receiver functions or structs

### Getting started
Install dependencies (do not apply if you ran the install_dependencies script):
- ``` $ go get github.com/onsi/ginkgo/v2/ginkgo ```
- ``` $ go get github.com/onsi/gomega/ ```
- ``` $ go install github.com/golang/mock/mockgen@v1.6.0 ```
- ``` $ go install github.com/axw/gocov/gocov@latest ```
- ``` $ go install github.com/AlekSi/gocov-xml@latest ```
- ``` $ go get github.com/matm/gocov-html ```
- Install the cli of ginkgo by running the command: ``` $ go install github.com/onsi/ginkgo/ginkgo  ```

#### Running
To run all tests and view a the coverage report:
```
$ ./scripts/test.sh
```

To test all suites you can use the following command:
```
$ ginkgo -r
```

To run only one suite of test you have to run the following commnd on the file directory: 
```
$ ginkgo 
```

To create a suite of testing run the following command inside of the folder: 
```
$ ginkgo bootstrap 
``` 

# Todo's
- [ ] Add a Dockerfile
- [ ] Create a GitHub action or pipeline to run application test
- [ ] Add a precommit hook to lint application
- [ ] Add traceability
- [ ] Add a logger
- [ ] Create a capability or lambda function or web server