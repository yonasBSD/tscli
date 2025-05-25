# **tscli**

`tscli` is a fast, single-binary CLI for the [Tailscale HTTP API](https://tailscale.com/api).
From your terminal you can manage devices, users, auth keys, webhooks, posture integrations, tailnet-wide settings, and even hit raw endpoints when the SDK hasn’t caught up yet.

---

## ✨ Highlights

| Area                     | What you can do                                                                                             |
| ------------------------ | ----------------------------------------------------------------------------------------------------------- |
| **Devices**              | List, get, (de)authorize, rename, force IPv4, enable subnet routes, expire, set / delete posture attributes |
| **Keys**                 | List & get existing keys; create **auth-keys** *or* **OAuth clients** (with full scope/tag validation)      |
| **Users**                | List (filter by type / role), get, suspend / restore / approve, manage invites                              |
| **Tailnet settings**     | Get & patch booleans + key-expiry with a single command (`tscli set settings …`)                            |
| **Policy file (ACL)**    | Fetch as raw HUJSON **or** canonical JSON                                                                   |
| **Webhooks**             | List, get, delete, **create** (generic / Slack) with subscription & provider validation                     |
| **Posture integrations** | List, get, create, patch existing integrations                                                              |
| **Invites**              | List / delete device- or user-invites                                                                       |
| **Contacts**             | Get & update contact emails                                                                                 |
| **Debug switch**         | `--debug` or `TSCLI_DEBUG=1` prints full HTTP requests / responses to stderr                                |
| **Config precedence**    | *flags* → *env* → `~/.tscli/.tscli.yaml` (or local `./.tscli.yaml`)                                         |

---

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


---

## ⚙️ Configuration

| Option            | Flag / Env var                          | YAML key  | Default |
| ----------------- | --------------------------------------- | --------- | ------- |
| Tailscale API key | `--api-key`, `-k` / `TAILSCALE_API_KEY` | `api-key` | —       |
| Tailnet name      | `--tailnet`, `-n` / `TAILSCALE_TAILNET` | `tailnet` | `-`     |

```yaml
# ~/.tscli/.tscli.yaml
api-key: tskey-abc123…
tailnet: example.com
```

---

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

---

## 📜 Coverage

| API area / action                |     Status    | `tscli` command                                             |
| -------------------------------- | :-----------: | ----------------------------------------------------------- |
| **Devices**                      |               |                                                             |
| list devices                     |  **complete** | `device list`                                               |
| get device                       |  **complete** | `device get --device <id>`                                  |
| authorize / de-authorize device  |  **complete** | `device authorize --device <id> [--approve=<bool>]`         |
| expire device key                |  **complete** | `device expire --device <id>`                               |
| set device name                  |  **complete** | `set name --device <id> --name <host>`                      |
| set device tags                  |  **complete** | `set tags --device <id> --tag <tag>`                        |
| set device IPv4                  |  **complete** | `set ip --device <id> --ip <addr>`                          |
| list subnet routes               |  **complete** | `list routes --device <id>`                                 |
| set subnet routes                |  **complete** | `set routes --device <id> --route <cidr> …`                 |
| **delete device**                |  **complete** | `delete device --device <id>`                               |
| **Device-posture attributes**    |               |                                                             |
| get attributes                   |  **complete** | `get posture --device <id>`                                 |
| set attribute                    |  **complete** | `set attribute --device <id> --key custom:x --value 42`     |
| delete attribute                 |  **complete** | `delete attribute --device <id> --key custom:x`             |
| **Posture integrations**         |               |                                                             |
| list integrations                |  **complete** | `list posture-integrations`                                 |
| get integration                  |  **complete** | `get posture-integration --id <id>`                         |
| create integration               |  **complete** | `create posture-integration --provider …`                   |
| update integration               |  **complete** | `set posture-integration --id <id> …`                       |
| delete integration               |  *incomplete* | —                                                           |
| **Auth / OAuth keys**            |               |                                                             |
| list keys                        |  **complete** | `list keys`                                                 |
| get key                          |  **complete** | `get key --key <id>`                                        |
| create auth-key                  |  **complete** | `create key --type authkey …`                               |
| create OAuth client              |  **complete** | `create key --type oauthclient …`                           |
| delete / revoke key              |  *incomplete* | —                                                           |
| **Users**                        |               |                                                             |
| list users                       |  **complete** | `list users [--type …] [--role …]`                          |
| get user                         |  **complete** | `get user --user <id>`                                      |
| approve / suspend / restore user |  **complete** | `set user-access --user <id> --approve/--suspend/--restore` |
| delete user                      |  **complete** | `delete user --user <id>`                                   |
| **Invites**                      |               |                                                             |
| list user invites                |  **complete** | `list invites user [--state …]`                             |
| list device invites              |  **complete** | `list invites device --device <id>`                         |
| delete invite                    |  *incomplete* | —                                                           |
| get invite                       |  *incomplete* | —                                                           |
| **Contacts**                     |               |                                                             |
| update contact                   |  **complete** | `set contacts --contact <id> --email <e@x>`                 |
| list / create / delete contacts  |  *incomplete* | —                                                           |
| **DNS**                          |               |                                                             |
| list nameservers                 |  **complete** | `list nameservers`                                          |
| set nameservers                  |  **complete** | `set nameservers --nameserver <ip> …`                       |
| advanced DNS settings            |  *incomplete* | —                                                           |
| **Policy file (ACL)**            |               |                                                             |
| get policy file                  |  **complete** | `get policy [--json]`                                       |
| set policy file                  |  **complete** | `set policy --file <acl.hujson>`                            |
| preview rule matches             |  **complete** | `get policy-preview --type … --value … [--file]`            |
| policy history / tests           |  *incomplete* | —                                                           |
| **Tailnet settings**             |               |                                                             |
| get settings                     |  **complete** | `get settings`                                              |
| update settings                  |  **complete** | `set settings --devices-approval …`                         |
| **Webhooks**                     |               |                                                             |
| list webhooks                    |  **complete** | `list webhooks`                                             |
| get webhook                      |  **complete** | `get webhook --webhook <id>`                                |
| create webhook                   |  **complete** | `create webhook --url <endpoint> --subscription …`          |
| update webhook                   |  **complete** | `set webhook --webhook <id> …`                              |
| delete webhook                   |  **complete** | `delete webhook --webhook <id>`                             |
| rotate webhook secret            |  *incomplete* | —                                                           |


---

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

---

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

---

## 📄 License

MIT — see `LICENSE`.
