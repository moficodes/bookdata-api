# bookdata-api

This is a demo application showing how to build rest api using golang.

## Run Locally

To run, from the root of the repo

```
go run .
```

## Access the app 

The App has a few Endpoints

All api endpoints are prefixed with `/api/v1`

To reach any endpoint use `baseurl:8080/api/v1/{endpoint}`

```text
Get Books by Author
"/books/authors/{author}" 
Optional query parameter for ratingAbove ratingBelow limit and skip

Get Books by BookName
"/books/book-name/{bookName}"
Optional query parameter for ratingAbove ratingBelow limit and skip


Get Book by ISBN
"/book/isbn/{isbn}"

Delete Book by ISBN
"/book/isbn/{isbn}"

Create New Book
"/book"
```

## Deploy app to cloud

This step is completely optional. But if you want to run you go app in the cloud somewhere IBM Cloud has an excellent paas solution using Cloud Foundry.

If you want to follow along

*  [Click Here to Get Started](https://cloud.ibm.com/registration?cm_mmc=Email_Events-_-Developer_Innovation-_-WW_WW-_-advocates:roger-osorio,eherrer\title:buildyourfirstapiwithgolang-newyorkcity-08212019\eventid:5d48e3b2c60e6be4b7305a2e\date:Aug2019\type:workshop\team:global-devadvgrp-newyork\city:newyorkcity\country:unitedstates&cm_mmca1=000019RS&cm_mmca2=10004805&cm_mmca3=M99938765&eventid=5d48e3b2c60e6be4b7305a2e&cvosrc=email.Events.M99938765&cvo_campaign=000019RS)
* Install[ IBM Cloud CLI](https://cloud.ibm.com/docs/cli?topic=cloud-cli-install-ibmcloud-cli#shell_install)

Once you have an IBM Cloud account and the IBM Cloud CLI installed,

From terminal

```text
ibmcloud login
```

```text
ibmcloud target --cf
```

Open the `manifest.yaml` in the root of the cloned repository.

Change the app name to something you like. After the change the file should look something like this

```text
---
applications:
- name: <your-app-name>
  random-route: true
  memory: 256M
  env:
    GOVERSION: go1.12
    GOPACKAGENAME: bookdata-api
  buildpack: https://github.com/cloudfoundry/go-buildpack.git
```

Then from the root of the repo run

```text
ibmcloud cf push
```

Wait a few minutes and voila! It should be running. 

 To find your app url 

```text
ibmcloud cf apps
```

You should see your app as running and also have the URL there.

