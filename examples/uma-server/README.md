# uma-server

This is a simple Gin server (https://gin-gonic.com) that implements the UMA protocol using the Lightspark SDK.

## Usage

Configuration parameters (API keys, etc.) and information on how to set them can be found in `config.go`.

You can run this server by running `go run main.go` from the `examples/uma-server` directory. You can also build
a binary by running `go build` from the same directory. You can then run the binary by running `./uma-server`.

By default, this server will run on port 8080, but you can set the PORT environment variable to change that. You can
make a request to the API through curl to make sure the server is working properly (replace bob with the username you
have configured). Here's a full example:

First, we'll start two instances of the server, one on port 8080 and one on port 8081 (in separate terminals):

Terminal 1:
```bash
# First set up config variables. You can also save these in a file or export them to your environment.
$ export LIGHTSPARK_API_TOKEN_CLIENT_ID=<client_id>
$ export LIGHTSPARK_API_TOKEN_CLIENT_SECRET=<client_secret>
# etc... See config.go for the full list of config variables.

# Now start the server on port 8080
$ PORT=8080 go run .
```

Terminal 2:
```bash
# First set up the variables as above. If you want to be able to actually send payments, use a different account.
$ export LIGHTSPARK_API_TOKEN_CLIENT_ID=<client_id_2>
$ export LIGHTSPARK_API_TOKEN_CLIENT_SECRET=<client_secret_2>
# etc... See config.go for the full list of config variables.

# Now start the server on port 8081
$ PORT=8081 go run .
```

Now, you can test the full uma flow like:

```bash
# First, call to vasp1 to lookup Bob at vasp2. This will return currency conversion info, etc. It will also contain a 
# callback ID that you'll need for the next call
$ curl -X GET http://localhost:8080/api/umalookup/\$bob@localhost:8081 -u diane:pa55word

# Now, call to vasp1 to get a payment request from vasp2. Replace the last path component here with the callbackUuid
# from the previous call. This will return an invoice and another callback ID that you'll need for the next call.
$ curl -X GET "http://localhost:8080/api/umapayreq/52ca86cd-62ed-4110-9774-4e07b9aa1f0e?amount=100&currencyCode=USD" -u diane:pa55word

# Now, call to vasp1 to send the payment. Replace the last path component here with the callbackUuid from the payreq
# call. This will return a payment ID that you can use to check the status of the payment.
curl -X POST http://localhost:8080/api/sendpayment/e26cbee9-f09d-4ada-a731-965cbd043d50 -u diane:pa55word
```

### Running with Docker

You can also run this server with Docker. First we need to build the image. From the root `go-sdk` directory, run:

```bash
$ docker build -t uma-server -f examples/uma-server/Dockerfile .
```

Next, we need to set up the config variables. You can do this by creating a file called `local.env` in the root `go-sdk`
directory. This file should contain the following:

```bash
LIGHTSPARK_API_TOKEN_CLIENT_ID=<your lightspark API token client ID from https://app.lightspark.com/api-config>
LIGHTSPARK_API_TOKEN_CLIENT_SECRET=<your lightspark API token client secret from https://app.lightspark.com/api-config>
LIGHTSPARK_UMA_NODE_ID=<your lightspark node ID. ex: LightsparkNodeWithOSKLND:018b24d0-1c45-f96b-0000-1ed0328b72cc>
LIGHTSPARK_UMA_RECEIVER_USER=<receiver UMA>

# These can be generated as described at https://docs.uma.me/uma-standard/keys-authentication-encryption
LIGHTSPARK_UMA_ENCRYPTION_CERT_CHAIN=<pem-encoded x509 certificate chain containing encryption pubkey>
LIGHTSPARK_UMA_ENCRYPTION_PUBKEY=<hex-encoded encryption pubkey>
LIGHTSPARK_UMA_ENCRYPTION_PRIVKEY=<hex-encoded encryption privkey>
LIGHTSPARK_UMA_SIGNING_CERT_CHAIN=<pem-encoded x509 certificate chain containing signing pubkey>
LIGHTSPARK_UMA_SIGNING_PUBKEY=<hex-encoded signing pubkey>
LIGHTSPARK_UMA_SIGNING_PRIVKEY=<hex-encoded signing privkey>

# If you are using an OSK node:
LIGHTSPARK_UMA_OSK_NODE_SIGNING_KEY_PASSWORD=<password for the signing key>

# If you are using a remote signing node:
LIGHTSPARK_UMA_REMOTE_SIGNING_NODE_MASTER_SEED=<hex-encoded master seed>

# Optional: A custom VASP domain in case you're hosting this at a fixed hostname.
LIGHTSPARK_UMA_VASP_DOMAIN=<your custom VASP domain. ex: vasp1.example.com>
```

Then, run the image:

```bash
$ docker run --env-file local.env -p 8081:8081 uma-server
```