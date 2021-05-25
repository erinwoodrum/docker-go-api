# Dockerized Golang API for testing out Spinnaker in Google Cloud. 

- Change this app to only use static data and not connect to db. 
- Make sure all unit and integrated tests pass.
- Setup env w/ Google Cloud Shell
- Download THIS sample app, create git repo and upload to Google Cloud Source Repo. 
- Deploy Spinnaker to Kubernetes Engine using Helm.
- Build Docker image
- Create triggers to create Docker images when app changes. 
- Configure Spinnaker pipeline to reliably and continuously deploy this app to Kubernetes Engine
- Deploy a code change and see the deployment to production. 