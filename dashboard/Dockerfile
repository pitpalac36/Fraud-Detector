FROM node:16

WORKDIR /dashboard

ENV PATH /dashboard/node_modules/.bin:$PATH
ENV REACT_APP_RESULTS_URL ws://localhost:8084/results
ENV PRODUCTION 1

COPY package.json ./
COPY package-lock.json ./
RUN npm install --silent

COPY . ./

EXPOSE 3000

CMD ["npm", "start"]