# Savannah-Informatics-technical-backend-challange

## About The Project

This technical challange was designed to test the skills/knowledge/experience that a Savannah Informatics Backend Developer should possess for the purposes of recruitment. It can also be used as a yardstick for training and improving existing backend developers.

The technical interview was built around a coding assignment that is designed to screen for the following basic skills:
* Experience in developing REST and GraphQL APIs in Python/Go
* Experience with a configuration management tool e.g. Chef, Puppet, Ansible etc
* Experience working with an infrastructure as code tool e.g. Terraform or Pulumi
will be an advantage.
* Experience working with containers and container orchestration tools e.g. k8s
* Experience writing automated tests at all levels - unit, integration and acceptance testing
* Experience with CI/CD (any CI/CD platform)

They were particularly interested in:
- Testing + coverage + CI/CD
- HTTP and APIs e.g. REST
- OAuth2
- Web security e.g. XSS
- Logic / flow of thought
- Setting up a database
- Version control

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Screening Test

They were particularly interested in:
1. Create a simple Python or Go service.
2. Design a simple customers and orders database (keep it simple)
3. Add a REST or GraphQL API to input / upload customers and orders:
    - Customers have simple details e.g., name and code.
    - Orders have simple details e.g., item, amount, and time.
4. Implement authentication and authorization via OpenID Connect
5. When an order is added, send the customer an SMS alerting them (you can use the
Africa’s Talking SMS gateway and sandbox)
6. Write unit tests (with coverage checking) and set up CI + automated CD. You can deploy
to any PAAS/FAAS/IAAS of your choice
7. Write a README for the project and host it on your GitHub

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

Fot the technical screening test I decided to develop the service using Golang as my primary language and PostgreSQL as my database of choice.

## Features

- Create and manage customers
  - Implement authentication and authorizatio
- Create and manage orders
- Record all order changes
- When an order is added, send the customer an SMS alerting them


### Pre-requisites

Requirements for the software and other tools to work, build, test and push 
- [Golang](https://go.dev/) `1.20.3` or newer versions
- [Gin](https://gin-gonic.com/)
- [PostgreSQL](https://www.postgresql.org/)
- [Gorm](gorm.io/)
- [Docker](https://www.docker.com/)
- [JWT](https://jwt.io/)
- [Github-Action](https://docs.github.com/en/actions)
- [Africa's Talking SMS Gateway](https://developers.africastalking.com/docs/sms/sending/bulk)

### Docker and Postgres

- Execute `docker pull postgres:12-alpine` to get the postgres image
- Execute `docker run --name customer_order_service -p 5432:5432 -e POSTGRES_USER=abdulrahman -e POSTGRES_PASSWORD=secret -d postgres:12-alpine` to run the postgres container
- Execute `docker logs customer_order_service` to see the logs
- Execute `docker exec -it customer_order_service psql -U root` to connect to the postgres container and login as `abdulrahman` user

### Installing Setup

The easiest way to get started with the rest api in your local environment is to clone it using git:

```
git clone https://github.com/AbdulrahmanDaud10/Savannah-Informatics-technical-backend-challange.git
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

**Environment Variables**

You need to configure environment variables for the rest api to work as expected.

Copy the content on .env_example to .env file:
```
cat .env_example > .env
``` 

The .env file should look like this:

```
# Database Configurations
DB_HOST = "127.0.0.1"
DB_PORT = "5432"
DB_USER = "abdulrahman"
DB_NAME = "customer_order_service"
DB_PASSWORD = ""

# Port for Application to listen and Server
LISTEN_ADDRESS = "127.0.0.1"
LISTEN_PORT = "3000"

# Africa's Talking Bulk SMS Configurations
BASE_LIVE_ENDPOINT    = " "
BASE_SANDBOX_ENDPOINT = " "
API_KEY               = " "
USERNAME              = " "
```
You can change the values to your own preference.
<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/Savannah-Informatics-technical-backend-challange`)
3. Commit your Changes (`git commit -m 'Add some Amazing Feature'`)
4. Push to the Branch (`git push origin feature/Savannah-Informatics-technical-backend-challange.git`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Acknowledgments
Some of the articles I read to help me solve the technical challang:

  - [README Templates](https://www.readme-templates.com/)
  - [TDD](https://codechill.hashnode.dev/testing-in-go-increasing-efficiency-of-code?ref=twitter-share)
  - [Docker](https://www.bacancytechnology.com/blog/dockerize-golang-application)
  - [Github Actions](https://medium.com/swlh/setting-up-github-actions-for-go-project-ea84f4ed3a40)
  - [Gorm](https://gorm.io/docs/models.html)

## Limitations
Some of the few things that need improvement:
* `JWT` feature should be enhanced:
  ````
  ````
* Africa's Talking SMS:
  ````
  Post "": unsupported protocol scheme ""[GIN] 2023/09/19 - 17:21:05 | 200 |     988.168µs |       127.0.0.1 | POST     "/api/messages/send-bulk-sms"
  ````
* GCP Golang Engine Configuration:
  ```
  ERROR: (gcloud.app.deploy) User [dulisyke3015@gmail.com] does not have permission to access apps instance [silver-treat-399511] (or it may not exist): Read access to project 'silver-treat-399511' was denied: please check billing account associated and retry
  ```