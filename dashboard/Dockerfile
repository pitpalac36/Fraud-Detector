FROM node:16

WORKDIR /dashboard

ENV PATH /dashboard/node_modules/.bin:$PATH

COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent

COPY . ./

EXPOSE 3000

CMD ["npm", "start"]