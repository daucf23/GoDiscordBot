# Discord Bot Service Setup on Raspberry Pi 5

This guide provides step-by-step instructions to compile a Go-based Discord bot and set it up as a systemd service on a Raspberry Pi 5. By the end of this guide, your Discord bot will run automatically on startup and can be managed using standard `systemctl` commands.

---

## Table of Contents

- [Prerequisites](#prerequisites)
- [Compiling the Go Binary](#compiling-the-go-binary)
- [Setting Up the Systemd Service](#setting-up-the-systemd-service)
  - [1. Create the Service File](#1-create-the-service-file)
  - [2. Configure the Service File](#2-configure-the-service-file)
  - [3. Reload Systemd and Enable the Service](#3-reload-systemd-and-enable-the-service)
  - [4. Start and Test the Service](#4-start-and-test-the-service)
- [Environment Variables](#environment-variables)
  - [Method 1: Define in Service File](#method-1-define-in-service-file)
  - [Method 2: Use an Environment File](#method-2-use-an-environment-file)
- [Commands Summary](#commands-summary)
- [Conclusion](#conclusion)

---

## Prerequisites

- **Raspberry Pi 5** with Raspberry Pi OS installed.
- **Go Programming Language** installed on your development machine.
- **Go-based Discord bot** source code.
- **User account** with `sudo` privileges on the Raspberry Pi.
- **Internet connection** for the Raspberry Pi.
- **OpenAI API Key** (if your bot uses OpenAI services).

---

## Compiling the Go Binary

To ensure your Go binary is compatible with the Raspberry Pi 5 architecture, you need to compile it for Linux ARM64.

1. **Set the Go Environment Variables:**

   ```bash
   export GOOS=linux
   export GOARCH=arm64
   ```

2. **Compile the Go Program:**

   Navigate to your Go project's root directory and run:

   ```bash
   go build -o /home/pi/discord-bot /path/to/your/main.go
   ```

   - Replace `/home/pi/discord-bot` with the desired output path on your Raspberry Pi.
   - Replace `/path/to/your/main.go` with the path to your `main.go` file.

3. **Transfer the Binary to Raspberry Pi (if compiled elsewhere):**

   Use `scp` or another file transfer method:

   ```bash
   scp /path/to/discord-bot pi@raspberry_pi_ip:/home/pi/discord-bot
   ```

4. **Set Executable Permissions on Raspberry Pi:**

   ```bash
   chmod +x /home/pi/discord-bot
   ```

---

## Setting Up the Systemd Service

### 1. Create the Service File

On your Raspberry Pi, create a new service file:

```bash
sudo nano /etc/systemd/system/discord-bot.service
```

### 2. Configure the Service File

Paste the following content into the file:

```ini
[Unit]
Description=Discord Bot Service
After=network-online.target
Requires=network-online.target

[Service]
Type=simple
User=pi
WorkingDirectory=/home/pi/
ExecStart=/home/pi/discord-bot
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

- **User:** Replace `pi` with your username if different.
- **WorkingDirectory:** The directory where your binary resides.
- **ExecStart:** The path to your Go binary.

### 3. Reload Systemd and Enable the Service

```bash
sudo systemctl daemon-reload
sudo systemctl enable discord-bot.service
```

### 4. Start and Test the Service

```bash
sudo systemctl start discord-bot.service
sudo systemctl status discord-bot.service
```

- Verify that the service is **active (running)**.
- To view real-time logs:

  ```bash
  sudo journalctl -u discord-bot.service -f
  ```

---

## Environment Variables

If your Discord bot requires environment variables (e.g., `OPENAI_API_KEY`), you need to provide them to the service.

### Method 1: Define in Service File

Edit the service file:

```bash
sudo nano /etc/systemd/system/discord-bot.service
```

Add the `Environment` directive under the `[Service]` section:

```ini
Environment="OPENAI_API_KEY=your_api_key_here"
```

Replace `your_api_key_here` with your actual API key.

### Method 2: Use an Environment File

1. **Create an Environment File:**

   ```bash
   sudo nano /home/pi/discord-bot.env
   ```

2. **Add Your Environment Variables:**

   ```ini
   OPENAI_API_KEY=your_api_key_here
   ```

3. **Secure the Environment File:**

   ```bash
   sudo chown pi:pi /home/pi/discord-bot.env
   sudo chmod 600 /home/pi/discord-bot.env
   ```

4. **Modify the Service File to Include the Environment File:**

   ```bash
   sudo nano /etc/systemd/system/discord-bot.service
   ```

   Add the following line under the `[Service]` section:

   ```ini
   EnvironmentFile=/home/pi/discord-bot.env
   ```

5. **Reload Systemd and Restart the Service:**

   ```bash
   sudo systemctl daemon-reload
   sudo systemctl restart discord-bot.service
   ```

---

## Commands Summary

### Compiling the Go Binary

```bash
export GOOS=linux
export GOARCH=arm64
go build -o /home/pi/discord-bot /path/to/your/main.go
```

### Setting Up the Service

```bash
sudo nano /etc/systemd/system/discord-bot.service
```

**Service File Content:**

```ini
[Unit]
Description=Discord Bot Service
After=network-online.target
Requires=network-online.target

[Service]
Type=simple
User=pi
WorkingDirectory=/home/pi/
ExecStart=/home/pi/discord-bot
EnvironmentFile=/home/pi/discord-bot.env  # Only if using an environment file
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

### Enabling and Starting the Service

```bash
sudo systemctl daemon-reload
sudo systemctl enable discord-bot.service
sudo systemctl start discord-bot.service
```

### Checking Service Status and Logs

```bash
sudo systemctl status discord-bot.service
sudo journalctl -u discord-bot.service -f
```

### Environment Variables

**Create Environment File:**

```bash
sudo nano /home/pi/discord-bot.env
```

**Set Permissions:**

```bash
sudo chown pi:pi /home/pi/discord-bot.env
sudo chmod 600 /home/pi/discord-bot.env
```

---

## Conclusion

You have successfully compiled your Go-based Discord bot and set it up as a systemd service on your Raspberry Pi 5. The bot will now start automatically on boot and can be managed using `systemctl` commands.

**Tips:**

- **Managing the Service:**
  - **Stop the Service:** `sudo systemctl stop discord-bot.service`
  - **Restart the Service:** `sudo systemctl restart discord-bot.service`
  - **Disable the Service at Boot:** `sudo systemctl disable discord-bot.service`

- **Security:**
  - Keep your API keys secure by limiting file permissions.
  - Avoid hardcoding sensitive information into your code or service files.

- **Troubleshooting:**
  - Check logs using `journalctl` if the bot isn't working as expected.
  - Ensure the binary is executable and compiled for the correct architecture.

---

**Feel free to reach out if you encounter any issues or have questions. Happy coding!**