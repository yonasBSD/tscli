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

Below is the same coverage table with even-width columns for easier reading.

| **API Area / Action**            | **Status** | **`tscli` Command**                                              |                              |             |
| -------------------------------- | :--------: | ---------------------------------------------------------------- | ---------------------------- | ----------- |
| **Devices**                      |            |                                                                  |                              |             |
| List tailnet devices             |      ✅     | `tscli list devices`                                             |                              |             |
| Get a device                     |      ✅     | `tscli get device --device <device>`                             |                              |             |
| Delete a device                  |      ✅     | `tscli delete device --device <device>`                          |                              |             |
| Expire a device key              |      ✅     | `tscli set expiry --device <device>`                             |                              |             |
| List device routes               |      ✅     | `tscli list routes --device <device>`                            |                              |             |
| Set device routes                |      ✅     | `tscli set routes --device <device> --route <cidr>`              |                              |             |
| Authorize / de-authorize device  |      ✅     | `tscli set authorization --device <device> --approve=<bool>`     |                              |             |
| Set device name                  |      ✅     | `tscli set name --device <device> --name <hostname>`             |                              |             |
| Set device tags                  |      ✅     | `tscli set tags --device <device> --tag tag:<tag>`               |                              |             |
| Rotate device key                |      ❌     | —                                                                |                              |             |
| Set device IPv4 address          |      ✅     | `tscli set ip --device <device> --ip <ip>`                       |                              |             |
| Get posture attributes           |      ✅     | `tscli get posture --device <device>`                            |                              |             |
| Set custom posture attributes    |      ✅     | `tscli set posture --device <device> --key custom:x --value <v>` |                              |             |
| Delete custom posture attributes |      ✅     | `tscli delete posture --device <device> --key custom:x`          |                              |             |
| **Device Invites**               |            |                                                                  |                              |             |
| List device invites              |      ✅     | `tscli list invites device --device <device>`                    |                              |             |
| Create device invite             |      ✅     | `tscli create invite device --device <device> --email <email>`   |                              |             |
| Get a device invite              |      ❌     | —                                                                |                              |             |
| Delete a device invite           |      ✅     | `tscli delete invite device --id <invite-id>`                    |                              |             |
| Resend / accept device invite    |      ❌     | —                                                                |                              |             |
| **User Invites**                 |            |                                                                  |                              |             |
| List user invites                |      ✅     | `tscli list invites user [--state …]`                            |                              |             |
| Create user invite               |      ✅     | `tscli create invite user --email <email> [--role <role>]`       |                              |             |
| Get a user invite                |      ❌     | —                                                                |                              |             |
| Delete a user invite             |      ✅     | `tscli delete invite user --id <invite-id>`                      |                              |             |
| Resend user invite               |      ❌     | —                                                                |                              |             |
| **Logging**                      |            |                                                                  |                              |             |
| List configuration audit logs    |      ✅     | `tscli get logs config --start <t> [--end <t>]`                  |                              |             |
| List network flow logs           |      ✅     | `tscli get logs network --start <t> [--end <t>]`                 |                              |             |
| Log-streaming endpoints          |      ❌     | —                                                                |                              |             |
| **DNS**                          |            |                                                                  |                              |             |
| List DNS nameservers             |      ✅     | `tscli list nameservers`                                         |                              |             |
| Set DNS nameservers              |      ✅     | `tscli set nameservers --nameserver <ip> …`                      |                              |             |
| Get DNS preferences              |      ✅     | `tscli get dns preferences`                                      |                              |             |
| Set DNS preferences              |      ✅     | `tscli set dns preferences --magicdns=<bool>`                    |                              |             |
| List DNS search paths            |      ✅     | `tscli list dns searchpaths`                                     |                              |             |
| Set DNS search paths             |      ✅     | `tscli set dns searchpaths --searchpath <domain> …`              |                              |             |
| Get split-DNS map                |      ✅     | `tscli get dns split`                                            |                              |             |
| Patch split-DNS                  |      ✅     | `tscli set dns split --domain <d>=<ip,ip> …`                     |                              |             |
| Replace split-DNS                |      ✅     | `tscli set dns split --replace --domain <d>=<ip>`                |                              |             |
| **Keys**                         |            |                                                                  |                              |             |
| List tailnet keys                |      ✅     | `tscli list keys`                                                |                              |             |
| Create auth-key / OAuth client   |      ✅     | \`tscli create key --type authkey                                | oauthclient …\`              |             |
| Get key                          |      ✅     | `tscli get key --key <id>`                                       |                              |             |
| Delete / revoke key              |      ❌     | —                                                                |                              |             |
| **Policy File**                  |            |                                                                  |                              |             |
| Get policy file                  |      ✅     | `tscli get policy [--json]`                                      |                              |             |
| Set policy file                  |      ✅     | `tscli set policy --file <acl.hujson>`                           |                              |             |
| Preview rule matches             |      ✅     | \`tscli get policy preview --type user                           | ipport --value X \[--current | --file F]\` |
| Validate / test policy           |      ❌     | —                                                                |                              |             |
| **Posture Integrations**         |            |                                                                  |                              |             |
| List integrations                |      ✅     | `tscli list posture-integrations`                                |                              |             |
| Create integration               |      ✅     | `tscli create posture-integration --provider <p> …`              |                              |             |
| Get integration                  |      ✅     | `tscli get posture-integration --id <id>`                        |                              |             |
| Update integration               |      ✅     | `tscli set posture-integration --id <id> …`                      |                              |             |
| Delete integration               |      ❌     | —                                                                |                              |             |
| **Users**                        |            |                                                                  |                              |             |
| List users                       |      ✅     | `tscli list users [--type …] [--role …]`                         |                              |             |
| Get a user                       |      ✅     | `tscli get user --user <id>`                                     |                              |             |
| Update user role                 |      ✅     | `tscli set user-role --user <id> --role <role>`                  |                              |             |
| Approve / suspend / restore user |      ✅     | \`tscli set user-access --user <id> --approve                    | --suspend                    | --restore\` |
| Delete a user                    |      ✅     | `tscli delete user --user <id>`                                  |                              |             |
| **Contacts**                     |            |                                                                  |                              |             |
| Get contacts                     |      ✅     | `tscli get contacts`                                             |                              |             |
| Update contact                   |      ✅     | `tscli set contacts --contact <id> --email <e@x>`                |                              |             |
| Resend verification              |      ❌     | —                                                                |                              |             |
| **Webhooks**                     |            |                                                                  |                              |             |
| List webhooks                    |      ✅     | `tscli list webhooks`                                            |                              |             |
| Create webhook                   |      ✅     | `tscli create webhook --url <endpoint> …`                        |                              |             |
| Get webhook                      |      ✅     | `tscli get webhook --webhook <id>`                               |                              |             |
| Update webhook                   |      ✅     | `tscli set webhook --webhook <id> …`                             |                              |             |
| Delete webhook                   |      ✅     | `tscli delete webhook --webhook <id>`                            |                              |             |
| Test / rotate webhook            |      ✅     | `tscli rotate webhook --webhook <id>`                            |                              |             |
| **Tailnet Settings**             |            |                                                                  |                              |             |
| Get tailnet settings             |      ✅     | `tscli get settings`                                             |                              |             |
| Update tailnet settings          |      ✅     | `tscli set settings --devices-approval …`                        |                              |             |

> **Legend** – ✅ implemented ❌ not yet implemented


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
