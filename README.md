# **tscli**

`tscli` is a fast, single-binary CLI for the [Tailscale HTTP API](https://tailscale.com/api).
From your terminal you can manage devices, users, auth keys, webhooks, posture integrations, tailnet-wide settings, and even hit raw endpoints when the SDK hasn’t caught up yet.

## ✨ Highlights

| Area                     | What you can do                                                                                             |
| ------------------------ | ----------------------------------------------------------------------------------------------------------- |
| **Devices**              | List, get, (de)authorize, rename, force IPv4, enable subnet routes, expire, set / delete posture attributes |
| **Keys**                 | List & get existing keys; create **auth-keys** _or_ **OAuth clients** (with full scope/tag validation)      |
| **Users**                | List (filter by type / role), get, suspend / restore / approve, manage invites                              |
| **Tailnet settings**     | Get & patch booleans + key-expiry with a single command (`tscli set settings …`)                            |
| **Policy file (ACL)**    | Fetch as raw HUJSON **or** canonical JSON                                                                   |
| **Webhooks**             | List, get, delete, **create** (generic / Slack) with subscription & provider validation                     |
| **Posture integrations** | List, get, create, patch existing integrations                                                              |
| **Invites**              | List / delete device- or user-invites                                                                       |
| **Contacts**             | Get & update contact emails                                                                                 |
| **Debug switch**         | `--debug` or `TSCLI_DEBUG=1` prints full HTTP requests / responses to stderr                                |
| **Config precedence**    | _flags_ → _env_ → `~/.tscli/.tscli.yaml` (or local `./.tscli.yaml`)                                         |

## 🔧 Install

### 🔧 Installation

#### macOS / Linux (Homebrew)

```bash
brew tap jaxxstorm/tap
brew install tscli          # upgrades via ‘brew upgrade’
```

#### Windows (Scoop)

```powershell
scoop bucket add jaxxstorm https://github.com/jaxxstorm/scoop-bucket.git
scoop install tscli
```

#### Nix

```bash
nix shell github:jaxxstorm/tscli
```

#### Manual download

Pre-built archives for **macOS, Linux, Windows (x86-64 / arm64)** are published on every release:

```bash
# example for Linux amd64
curl -sSfL \
  https://github.com/jaxxstorm/tscli/releases/latest/download/tscli_$(uname -s)_$(uname -m).tar.gz \
  | sudo tar -xz -C /usr/local/bin tscli
```

#### Go install (always builds from HEAD)

```bash
go install github.com/jaxxstorm/tscli@latest
```

After any method, confirm:

```bash
tscli --version
```

## ⚙️ Configuration

| Option            | Flag / Env var                          | YAML key  | Default |
| ----------------- | --------------------------------------- | --------- | ------- |
| Tailscale API key | `--api-key`, `-k` / `TAILSCALE_API_KEY` | `api-key` | —       |
| Tailnet name      | `--tailnet`, `-n` / `TAILSCALE_TAILNET` | `tailnet` | —       |

```yaml
# ~/.tscli/.tscli.yaml
api-key: tskey-abc123…
tailnet: example.com
format: pretty # other options are: human, json or yaml
```

## 🚀 Usage

```text
tscli <noun> <verb> [flags]
```

### Global flags

```
-k, --api-key string   Tailscale API key
-n, --tailnet string   Tailnet (default "-")
-d, --debug            Dump raw HTTP traffic to stderr
```

## 📜 Coverage

## 📜 Coverage

| API Area / Action                                    | Status | Command |
| ---------------------------------------------------- | ------ | ------- |
| **Devices**                                          |        |         |
| List tailnet devices                                 | :white_check_mark:        |         |
| Get a device                                         | :white_check_mark:     |         |
| Delete a device                                      | :white_check_mark:        |         |
| Expire a device's key                                | :white_check_mark:       |         |
| List device routes                                   | :white_check_mark:        |         |
| Set device routes                                    | :white_check_mark:       |         |
| Authorize device                                     | :white_check_mark:        |         |
| Set device name                                      |        |         |
| Set device tags                                      | :white_check_mark:        |         |
| Update device key                                    |        |         |
| Set device IPv4 address                              | :white_check_mark:       |         |
| Get device posture attributes                        | :white_check_mark:       |         |
| Set custom device posture attributes                 | :white_check_mark:       |         |
| Delete custom device posture attributes              | :white_check_mark:        |         |
| **Device Invites**                                   |        |         |
| List device invites                                  | :white_check_mark:       |         |
| Create device invites                                | :white_check_mark:      |         |
| Get a device invite                                  |        |         |
| Delete a device invite                               | :white_check_mark:        |         |
| Resend a device invite                               |        |         |
| Accept a device invite                               |        |         |
| **User Invites**                                      |        |         |
| List user invites                                    | :white_check_mark:        |         |
| Create user invites                                  | :white_check_mark:       |         |
| Get a user invite                                    |        |         |
| Delete a user invite                                 | :white_check_mark:       |         |
| Resend a user invite                                 |        |         |
| **Logging**                                          |        |         |
| List configuration audit logs                        | :white_check_mark:      |         |
| List network flow logs                               | :white_check_mark:      |         |
| Get log streaming status                             |        |         |
| Get log streaming configuration                      |        |         |
| Set log streaming configuration                      |        |         |
| Disable log streaming                                |        |         |
| Create or get AWS external id                        |        |         |
| Validate external ID integration with IAM role trust policy |        |         |
| **DNS**                                              |        |         |
| List DNS nameservers                                 | :white_check_mark:        |         |
| Set DNS nameservers                                  | :white_check_mark:        |         |
| Get DNS preferences                                  | :white_check_mark:      |         |
| Set DNS preferences                                  | :white_check_mark:       |         |
| List DNS search paths                                |        |         |
| Set DNS search paths                                 | :white_check_mark:       |         |
| Get split DNS                                        | :white_check_mark:      |         |
| Update split DNS                                     | :white_check_mark:       |         |
| Set split DNS                                        | :white_check_mark:       |         |
| **Keys**                                             |        |         |
| List tailnet keys                                    | :white_check_mark:       |         |
| Create an auth key or OAuth client                   | :white_check_mark:      |         |
| Get key                                              | :white_check_mark:      |         |
| Delete key                                           |        |         |
| **Policy File**                                      |        |         |
| Get policy file                                      | :white_check_mark:      |         |
| Set policy file                                      |        |         |
| Preview rule matches                                 |        |         |
| Validate and test policy file                        |        |         |
| **Device Posture**                                   |        |         |
| List all posture integrations                        | :white_check_mark:        |         |
| Create a posture integration                         | :white_check_mark:       |         |
| Get a posture integration                            |        |         |
| Update a posture integration                         | :white_check_mark:       |         |
| Delete a posture integration                         |        |         |
| **Users**                                            |        |         |
| List users                                           | :white_check_mark:        |         |
| Get a user                                           |        |         |
| Update user role                                     | :white_check_mark:       |         |
| Approve a user                                       | :white_check_mark:       |         |
| Suspend a user                                       | :white_check_mark:       |         |
| Restore a user                                       |        |         |
| Delete a user                                        | :white_check_mark:        |         |
| **Contacts**                                         |        |         |
| Get contacts                                         | :white_check_mark:      |         |
| Update contact                                       | :white_check_mark:       |         |
| Resend verification email                            |        |         |
| **Webhooks**                                         |        |         |
| List webhooks                                        | :white_check_mark:       |         |
| Create a webhook                                     | :white_check_mark:       |         |
| Get webhook                                          | :white_check_mark:      |         |
| Update webhook                                       |        |         |
| Delete webhook                                       | :white_check_mark:        |         |
| Test a webhook                                       |        |         |
| Rotate webhook secret                                | :white_check_mark:       |         |
| **TailnetSettings**                                  |        |         |
| Get tailnet settings                                 | :white_check_mark:      |         |
| Update tailnet settings                              | :white_check_mark:       |         |

### Quick examples

```bash
# Approve a waiting device
tscli device authorize --device node-abc123 --approve

# Rotate an auth-key that expires in 30 days
tscli create key --description "CI" --expiry 720h | jq .key

# Create Slack webhook for device deletions
tscli create webhook \
  --url https://hooks.slack.com/services/T000/B000/XXXXX \
  --provider slack \
  --subscription nodeDeleted
```

## 🛠 Development

```bash
git clone https://github.com/jaxxstorm/tscli
cd tscli
TAILSCALE_API_KEY=tskey-… TAILSCALE_TAILNET=example.com go run ./cmd/tscli list devices
```

Tests & lint:

```bash
go test ./...
```

## 📄 License

MIT — see [`LICENSE`](./LICENSE).
