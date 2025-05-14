# **tscli**

`tscli` is a fast, single-binary CLI for the [Tailscale HTTP API](https://tailscale.com/api).
List, query, tag, retag, authorize, or expire devices; manage keys; inspect users; fetch your ACL policy file—straight from the terminal.

---

## ✨ Features

| Area                  | What you can do                                                                                            |
| --------------------- | ---------------------------------------------------------------------------------------------------------- |
| **Devices**           | list, get, authorize / de-authorize, change name, set IPv4, enable routes, add / delete posture attributes |
| **Keys**              | list reusable auth-keys, get a single key                                                                  |
| **Users**             | list users (filter by type / role), get a single user                                                      |
| **Policy file (ACL)** | fetch as raw HUJSON or canonical JSON                                                                      |
| **Raw endpoints**     | helper for calling un-wrapped API paths                                                                    |
| **Config precedence** | *flags* → *env vars* → `config.yaml` (local dir or `~/.tscli/`)                                            |

---

## 🔧 Installation

```bash
go install github.com/jaxxstorm/tscli@latest
```

Binary goes to `$(go env GOPATH)/bin`.

---

## ⚙️ Configuration

| Option            | Flag / Env Var                          | Config key | Default |
| ----------------- | --------------------------------------- | ---------- | ------- |
| Tailscale API key | `--api-key`, `-k` / `TAILSCALE_API_KEY` | `api-key`  | ―       |
| Tailnet name      | `--tailnet`, `-n` / `TAILSCALE_TAILNET` | `tailnet`  | `-`     |

`config.yaml` example:

```yaml
api-key: tskey-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
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
-n, --tailnet string   Tailnet name (default "-")
```

---

## 📜 Command map (selected)

| Command                                                        | Purpose / Filters                               |                     |       |      |                                  |
| -------------------------------------------------------------- | ----------------------------------------------- | ------------------- | ----- | ---- | -------------------------------- |
| `tscli device list [--all]`                                    | List all devices (`--all` adds advanced fields) |                     |       |      |                                  |
| `tscli device get --device <id> [--all]`                       | Get one device                                  |                     |       |      |                                  |
| `tscli device authorize --device <id> [--approve=<bool>]`      | (De)authorize device                            |                     |       |      |                                  |
| `tscli device name --device <id> --name <host>`                | Rename device                                   |                     |       |      |                                  |
| `tscli device routes --device <id> --route <cidr> [--route …]` | Replace enabled subnet routes                   |                     |       |      |                                  |
| `tscli device ip --device <id> --ip <addr>`                    | Force a Tailscale IPv4 address                  |                     |       |      |                                  |
| `tscli device posture --device <id> --key custom:x --value 42` | Set posture attribute                           |                     |       |      |                                  |
| `tscli delete attribute --device <id> --key custom:x`          | Delete posture attribute                        |                     |       |      |                                  |
| `tscli list routes --device <id>`                              | Show advertised / enabled routes                |                     |       |      |                                  |
| `tscli list keys`                                              | List auth-keys                                  |                     |       |      |                                  |
| `tscli get key --key <id>`                                     | Get one auth-key                                |                     |       |      |                                  |
| \`tscli list users \[--type member                             | shared                                          | all] \[--role owner | admin | …]\` | List users with optional filters |
| `tscli get user --user <id>`                                   | Get one user                                    |                     |       |      |                                  |
| `tscli get policy [--json]`                                    | Print ACL policy (HUJSON or JSON)               |                     |       |      |                                  |

---

### Examples

```bash
# List devices, pretty-print with jq
tscli device list | jq '.[] | {id, hostname, authorized}'

# Tag a device
tscli set tags --device node-abc123 --tag tag:web --tag tag:prod

# Rotate a device's IPv4
tscli set ip --device node-abc123 --ip 100.64.0.42

# Fetch ACL policy as JSON
tscli get policy --json | jq '.groups'
```

---

## 🛠 Development

```bash
git clone https://github.com/jaxxstorm/tscli
cd tscli
TAILSCALE_API_KEY=tskey-... TAILSCALE_TAILNET=example.com \
  go run ./cmd/tscli list devices
```

Tests & lint:

```bash
go test ./...
```

---

## 📄 License

MIT (see `LICENSE`).
