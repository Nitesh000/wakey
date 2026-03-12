# Wakey

![Wakey Logo](./logo.png)

A lightweight CLI tool written in **Go** that prevents your computer from going idle by **gently moving the mouse cursor**.

---

## Features

- Keeps system awake by moving the cursor periodically
- Runs indefinitely by default
- Supports custom run durations
- Adjustable movement interval
- Lightweight and fast
- Cross-platform CLI (macOS, Linux, Windows)

---

## Installation

### Build from Source

Requirements:

- Go **1.20+**

Clone the repository:

```
git clone https://github.com/YOUR_USERNAME/wakey.git
cd wakey
```

Build the binary:

```
make build
```

Run:

```
./bin/wakey
```

---

### Install Globally

```
make install
```

Run from anywhere:

```
wakey
```

---

## Usage

### Run indefinitely

```
wakey
```

---

### Run for a specific duration

Run for **30 minutes**

```
wakey --minutes 30
```

Run for **2 hours**

```
wakey --hours 2
```

Run for **90 seconds**

```
wakey --seconds 90
```

---

### Adjust movement interval

```
wakey --interval 10
```

---

## Example Output

```
Screen detected: 2560x1440
Moving cursor to 1200,640
Moving cursor to 800,900
Moving cursor to 2100,500
```

---

## Makefile Commands

Build project:

```
make build
```

Run project:

```
make run
```

Run with flags:

```
make run ARGS="--minutes 30"
```

Clean build files:

```
make clean
```

Install globally:

```
make install
```

---

## macOS Permission

Because Wakey moves the cursor, macOS requires accessibility permission.

Go to:

```
System Settings
→ Privacy & Security
→ Accessibility
```

Add:

- Terminal
  or
- the wakey binary

---

## Project Structure

```
wakey
├─ cmd
│  └─ wakey
│     └─ main.go
├─ bin
├─ Makefile
├─ go.mod
├─ README.md
└─ LICENSE
```

---

## Stop the Program

```
Ctrl + C
```

---

## License

MIT License
