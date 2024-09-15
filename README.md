# akro
![](https://img.shields.io/github/actions/workflow/status/sweeneyngo/akro/deploy-build.yml)

Create passwords with "slightly coherent" sentences. Ideal for those who have trouble remembering passwords!

<p align="center">
<a href="https://ifuxyl.dev/akro">
<img src="https://i.imgur.com/NmfgmzP.png" width="800"><br>
<sup><strong>ifuxyl.dev/akro</a></strong></sup>
</p>

The application is written in Typescript + [React](https://react.dev/) and built with [Vite](https://vitejs.dev/).
Implemented with the [Markov chain generator](https://en.wikipedia.org/wiki/Markov_chain) with Go. Multiple independent replicas are hosted + provisioned in different distributed locations in U.S.A with [Fly](https://fly.io/).

<!-- See the [full article](https://www.ifuxyl.dev/blog/conway-hashlife) about seagull & HashLife! -->

## Building
Not necessarily in active development, but we welcome any contributions. Feel free to submit an issue or contribute code via PR to the `main` branch.

You need [Go 1.20](https://go.dev/) and [Docker](https://www.docker.com/) to work on the backend and [Node.js v22](https://nodejs.org/en/) for the frontend.

To build the site for development:
```bash
# If you don't have Node v22 or pnpm v9:
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash
nvm install node
npm install -g pnpm

# Install in project root
pnpm install && pnpm run dev
```

You should now access the webpage at `http://localhost:5173/akro/`,
Any changes in `src` will be immediately available through [Vite](https://vitejs.dev/).

To run the server:
```bash
go run .
```
The server listens for web requests on port 8080. 

To build the server via Docker:
```bash
docker build -t akro/v1 .
docker run -it akro/v1
```

### Deployment
```bash
fly deploy
```

## License

<sup>
All code is licensed under the <a href="LICENSE">MIT license</a>.
</sup>
