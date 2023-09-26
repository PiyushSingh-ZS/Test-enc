FROM ubuntu:latest

# installing curl
RUN apt-get update && \
    apt-get install -y curl && \
    apt-get -y install nano


# Install helm
RUN curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
RUN helm version

# Install Python
RUN apt-get install python3 -y
RUN python3 --version

# Install the Google Cloud SDK
RUN apt-get install apt-transport-https ca-certificates gnupg curl sudo -y
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg  add -
RUN apt-get update -y
RUN apt-get install google-cloud-cli -y

RUN apt-get install kubectl

RUN apt-get install google-cloud-sdk-gke-gcloud-auth-plugin

CMD ["/bin/bash"]