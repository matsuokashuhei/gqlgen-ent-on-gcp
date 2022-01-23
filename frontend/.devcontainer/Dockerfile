FROM node:16-buster

# ENV workdir=/usr/local/src
# WORKDIR $workdir
# COPY package.json package-lock.json $workdir
# RUN npm install

ENV workdir=/app/frontend
WORKDIR $workdir
COPY frontend/package.json frontend/package-lock.json $workdir
RUN npm install

# RUN npm run build
# CMD ["npm", "run", "start"]

# RUN useradd -m nodejs && \
#     mkdir -p $workdir && \
#     chown nodejs -R $workdir
# USER nodejs
# WORKDIR $workdir

# COPY --chown=nodejs frontend/package.json frontend/package-lock.json $workdir
# RUN npm install
# COPY --chown=nodejs . /app
# RUN npm run build

# CMD ["npm", "run", "start"]
