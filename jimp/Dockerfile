FROM node:18-alpine

WORKDIR /app

COPY index.js . 
COPY package*.json . 

RUN npm ci --omit dev

CMD ["index.js"]
