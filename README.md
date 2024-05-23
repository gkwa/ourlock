# ourlock

Purpose:
This is stop gap until I learn elasticsearch.

My notes are getting out of hand over the years; first it was org mode, then org mode with everything in one file, then markdown with very well define small files.

Now I need to find these notes and I'm using ripgrep, like this monstrocity:

```bash

rg test1\(\?s:.\*\?\)test2\|test2\(\?s:.\*\?\)test1 --ignore-case --glob-case-insensitive --multiline-dotall --color=always --glob=\*.\{md,txt,org} --context=20 '/Users/mtm/Documents/Obsidian Vault/' /Users/mtm/pdev/taylormonacelli/notes

```

Sigh...




## example usage

```bash

# find notes on golang charm package
ourlock search charm golang

# nothings' working, just show command that would be run
ourlock search charm golang --dry-run

```

## install ourlock


on macos/linux:
```bash

brew install taylormonacelli/homebrew-tools/ourlock

```


on windows:

```powershell

TBD

```
