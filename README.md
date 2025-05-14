# **tscli**

`tscli` is a lightweight Go-based command-line tool for interacting with the [Tailscale API](https://tailscale.com/).
It lets you list, inspect, and (de)-authorize devices in your tailnet from any machine that has Go installed.

---

## ✨ Features

* **List** all devices in your tailnet (`devices list`)
* **Get** full details for a single device (`devices get`)
* **Authorize / De-authorize** a device (`devices authorize`)
* Flag-overrides for showing **all API fields**
* Works with **environment variables** or **flags** – whichever you prefer

---

## 🔧 Installation

```bash
go install github.com/jaxxstorm/tscli@latest
```

The binary will appear in `$(go env GOPATH)/bin`.

---

## ⚙️ Configuration

| Option            | Flag / Env Var                          | Required | Notes                                                              |
| ----------------- | --------------------------------------- | -------- | ------------------------------------------------------------------ |
| Tailscale API key | `--api-key`, `-k` / `TAILSCALE_API_KEY` | ✅        | [Generate a key](https://login.tailscale.com/admin/settings/keys). |
| Tailnet name      | `--tailnet`, `-n` / `TAILSCALE_TAILNET` | ✅        | Looks like `example.com` or `corp.tailscale.net`.                  |

You can set them once as environment variables, pass them as flags every time, or mix and match—the CLI always uses **flags › env vars** precedence.

```bash
export TAILSCALE_API_KEY=tskey-...
export TAILSCALE_TAILNET=mycorp.com
```

---

## 🚀 Usage

```text
tscli <command> [flags]
```

### Top-level flags

```
-k, --api-key string   Tailscale API key.
-n, --tailnet string   Tailscale tailnet.
-h, --help             Help for any command.
```

---

## 📜 Command overview

| Command                   | Purpose                           | Key Flags                                                                              |
| ------------------------- | --------------------------------- | -------------------------------------------------------------------------------------- |
| `tscli devices list`      | List every device in the tailnet. | `--all`   Print **all** API fields instead of the default subset.                      |
| `tscli devices get`       | Get one device by ID.             | `--device <id>` (✔)   Device ID to fetch.<br>`--all`   Show all fields.                |
| `tscli devices authorize` | Approve / reject a device.        | `--device <id>` (✔)<br>`--approve` (default **true**; set `--approve=false` to reject) |

> ✔ = required flag

---

### Examples

#### List devices (default view)

```bash
tscli devices list
```

#### List devices with every field the API returns

```bash
tscli devices list --all
```

#### Get a single device

```bash
tscli devices get --device 123456abcdef
```

#### Approve a device

```bash
tscli devices authorize --device 123456abcdef --approve
```

#### Un-approve (disable) a device

```bash
tscli devices authorize --device 123456abcdef --approve=false
```

All commands print pretty-formatted JSON, making them easy to pipe into `jq`:

```bash
tscli devices list | jq '.[] | {id, hostname, authorized}'
```

---

## 🛠 Development

1. Clone the repo and make sure you have Go 1.22+.

2. Run the CLI from source:

   ```bash
   TAILSCALE_API_KEY=tskey-... \
   TAILSCALE_TAILNET=mycorp.com \
   go run ./cmd/tscli --help
   ```

3. Lint & test:

   ```bash
   make test      # or go test ./...
   make lint      # runs your preferred linter
   ```

---

## 📄 License

This project is licensed under the MIT license (see `LICENSE` file).
