# pyre-strict

import os
import secrets

DATABASE_URI: str = "sqlite+aiosqlite:///" + os.path.join(
    os.getcwd(), "instance", "nwc.sqlite"
)
SECRET_KEY: str = secrets.token_hex(32)

# This should match NWC_JWT_PUBKEY in your main UMA server env config.
UMA_VASP_JWT_PUBKEY = "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEEVs/o5+uQbTjL3chynL4wXgUg2R9\nq9UU8I5mEovUf86QZ7kOBIjJwqnzD1omageEHWwHdBO6B+dFabmdT9POxg==\n-----END PUBLIC KEY-----"
FRONTEND_BUILD_PATH = "../static"
UMA_VASP_LOGIN_URL = "http://localhost:8000/login"
UMA_VASP_TOKEN_EXCHANGE_URL = "http://vasp.local/umanwc/token"
VASP_UMA_API_BASE_URL = "http://vasp.local/umanwc"
VASP_NAME = "Go Demo VASP"

# Replace with your own constant private key via `openssl rand -hex 32` if you want.
NOSTR_PRIVKEY: str = secrets.token_hex(32)
RELAY = "wss://relay.getalby.com/v1"
NWC_APP_ROOT_URL = "http://localhost:8080"

VASP_SUPPORTED_COMMANDS = [
    "pay_invoice",
    "make_invoice",
    "lookup_invoice",
    "get_balance",
    "get_budget",
    "get_info",
    "list_transactions",
    "pay_keysend",
    "lookup_user",
    "fetch_quote",
    "execute_quote",
    "pay_to_address",
]
