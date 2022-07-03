FROM python:3.8-slim-buster


RUN mkdir /cnnmodel
COPY . /cnnmodel
WORKDIR /cnnmodel


RUN apt-get install ffmpeg libsm6 libxext6  -y
RUN apt-get update && apt-get install -y python3-opencv
RUN pip3 install opencv-python
# Expose the port the app will run on

EXPOSE 5001

# Add VOLUMEs to allow backup of config, logs and databases
#HEALTHCHECK CMD ["wget", "-q", "0.0.0.0:8080"]

# Set the default command to run when starting the container
