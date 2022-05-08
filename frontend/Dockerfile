FROM node:11.13.0

# create destination directory
RUN mkdir -p /frontend
WORKDIR /frontend

# copy the app, note .dockerignore
COPY . /frontend/
RUN npm install

# build necessary, even if no static files are needed,
# since it builds the server as well
#RUN npm run build

# expose 3000 on container
EXPOSE 3000

# set app serving to permissive / assigned
ENV NUXT_HOST=0.0.0.0
# set app port
ENV NUXT_PORT=3000
