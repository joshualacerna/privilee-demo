# **Hello Privilee** (DevOps Challenge)


### **GitHub Workflow (CICD)**

1) The first part of the challenge is to build a GitHub actions workflow to build the application and run the tests.
2) Update the workflow to deploy the application to EKS with the following jobs structure: build → test → deploy



### **Deployment**

Utilise Infrastructure as Code (IaC) tools such as CloudFormation or Terraform to automate the creation of the infrastructure needed to deploy this application in Bahrain. (Choose one of these options: EC2, ECS, EKS, Lambda). Assume any missing information necessary for making a decision and document your choices in a README file.



### **Context**

Considering the Hello Privilee application is a backend application that will provide a list of API endpoints for a mobile application.
- The application is deployable in different countries
- The application should handle the peak usage happening during weekend evenings (3x)
- The application should be protected against brute force attacks
- The application should be deployed without downtime (choose any deployment strategy)


### **Misc**

How to run tests in the application?
```sh
// Run the tests
go test
```

How to run the application?
```sh
// Run the app
go mod tidy && go run hello-privilee
```


### **Delivery**
Upload your solution to GitHub, Bitbucket or any Git repository that is publicly accessible and email the recruiter you are dealing with providing a link to the repository.
