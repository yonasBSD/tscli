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

## 📜 Command cheat-sheet (most common)

| Command                                                                                                                   | Purpose / Notes                                           |                                                           |                              |                  |
| ------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | ---------------------------- | ---------------- |
| **Devices**                                                                                                               |                                                           |                                                           |                              |                  |
| `tscli device list [--all]`                                                                                               | List devices (add `--all` for connectivity & routes)      |                                                           |                              |                  |
| `tscli device get  --device <id> [--all]`                                                                                 | Fetch one                                                 |                                                           |                              |                  |
| `tscli device authorize   --device <id> [--approve=<bool>]`                                                               | Approve / un-approve                                      |                                                           |                              |                  |
| `tscli device name        --device <id> --name <hostname>`                                                                | Rename                                                    |                                                           |                              |                  |
| `tscli device routes      --device <id> --routes 10.0.0.0/24,…`                                                           | Replace enabled subnet routes                             |                                                           |                              |                  |
| `tscli device ip          --device <id> --ip 100.64.0.42`                                                                 | Force IPv4                                                |                                                           |                              |                  |
| `tscli device expire      --device <id>`                                                                                  | Immediately expire node key                               |                                                           |                              |                  |
| `tscli set attribute      --device <id> --key custom:x --value 1`                                                         | Add / update posture attribute                            |                                                           |                              |                  |
| `tscli delete attribute   --device <id> --key custom:x`                                                                   | Delete posture attribute                                  |                                                           |                              |                  |
| **Keys**                                                                                                                  |                                                           |                                                           |                              |                  |
| `tscli list keys`                                                                                                         | List existing keys                                        |                                                           |                              |                  |
| `tscli get  key  --key <id>`                                                                                              | Show one                                                  |                                                           |                              |                  |
| \`tscli create key \[authkey                                                                                              | oauthclient …]\`                                          | Create auth-key or OAuth client (validates scopes & tags) |                              |                  |
| **Users & invites**                                                                                                       |                                                           |                                                           |                              |                  |
| \`tscli list users \[--type member                                                                                        | shared                                                    | all] \[--role admin                                       | …]\`                         | Filtered listing |
| `tscli get  user --user <id>`                                                                                             | Details                                                   |                                                           |                              |                  |
| \`tscli set user-access  --user <id> --suspend                                                                            | --restore                                                 | --approve\`                                               | Change approval / suspension |                  |
| \`tscli list invites user   \[--state pending                                                                             | accepted                                                  | all]\`                                                    | User invites                 |                  |
| `tscli list invites device --device <id>`                                                                                 | Device invites                                            |                                                           |                              |                  |
| **Policy & settings**                                                                                                     |                                                           |                                                           |                              |                  |
| `tscli get policy [--json]`                                                                                               | Pretty HUJSON (default) or JSON ACL policy                |                                                           |                              |                  |
| `tscli get settings`                                                                                                      | Current tailnet-wide toggles                              |                                                           |                              |                  |
| `tscli set settings --devices-approval=true …`                                                                            | Patch any subset of settings (requires at least one flag) |                                                           |                              |                  |
| **Webhooks**                                                                                                              |                                                           |                                                           |                              |                  |
| `tscli list webhooks`                                                                                                     | List                                                      |                                                           |                              |                  |
| `tscli get  webhook --id <id>`                                                                                            | Show                                                      |                                                           |                              |                  |
| `tscli delete webhook --id <id>`                                                                                          | Delete                                                    |                                                           |                              |                  |
| `tscli create webhook --url <https://…> --provider generic \`<br>`--subscription nodeCreated --subscription policyUpdate` | Create with provider + events                             |                                                           |                              |                  |
| **Posture integrations**                                                                                                  |                                                           |                                                           |                              |                  |
| `tscli list  posture-integrations`                                                                                        | List integrations                                         |                                                           |                              |                  |
| `tscli get   posture-integration --id <id>`                                                                               | Get one                                                   |                                                           |                              |                  |
| `tscli create posture-integration --provider falcon --cloud-id us-1 --client-secret …`                                    | Create new                                                |                                                           |                              |                  |
| `tscli set    posture-integration --id <id> --provider jamfpro --client-id XXX`                                           | Patch existing                                            |                                                           |                              |                  |

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
