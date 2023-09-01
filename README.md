# 📔🐐Gotes CLI

Welcome to Gotes! A delightful, minimalist command-line tool for managing your notes inspired by the fusion of 'gotes' and 'notes'. Because we believe that sometimes, simple things, when merged, can create pure magic! ✨

## Features:

- **Effortless Note Creation**: Just provide a category and title, and you're ready to jot down your thoughts!
- **Quick Access**: Dive straight into your notes with the `gotes open` command. Navigating your notes has never been this fun!
  
## Getting Started:

### 1. Installation:

$ go install github.com/trygvizl/gotes@latest

Make sure you've Go installed and your `$GOPATH/bin` is in `$PATH`.

### 2. Command Reference:

**Create a New Note**:

Syntax: `gotes create <category> <title>`

```bash
$ gotes create recipes "Grandma's Apple Pie"
```

**Open and Browse Notes**:

Syntax: `gotes open`

```bash
$ gotes open
```

## Notes on Gotes:

- **Storage**: By default, Gotes keeps all your notes safe and snug under `~/gotes` in your home directory.
  
- **Format**: Every note you create is saved as a `.md` (markdown) file, ensuring compatibility with many editors and platforms.

- **Editor**: We love Vim! Therefore, whenever you create or open a note, Gotes fires up Vim for you to type away. Don't worry, if Vim's not your cup of tea, we're looking to support other editors soon!

## Why Gotes?

- **Minimalistic Design**: No clutter, just the features you need.
    
- **Rapid Access**: Whether you're jotting down a thought or revisiting old ones, Gotes ensures your notes are just a command away.

- **Flexibility**: Change the storage location easily with the `GOTES_PATH` environment variable if you want to move away from the default.

---

📝 Happy Note-Taking!

