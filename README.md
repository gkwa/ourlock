# ourlock

Purpose:
This is stop gap until I learn elasticsearch.

My notes are getting out of hand over the years; first it was org mode, then org mode with everything in one file, then markdown with very well define small files.



So far, this seems useful:

- Find notes that contain 'golang' and 'charm' in order to remind myself how to setup golang charm

- Amongtst all my notes, show me blobs of notes that have 'elasticsearch' and 'getting started' within 100 characters of each other.








## example usage

Show all blobs/notes where 'elasticsearch' is within 100 characters of 'getting started'


```log

[mtm@taylors-MacBook-Pro-2:ourlock(master)]$ ourlock distancesearch elasticsearch "getting started" --distance 100 --dry-run
rg 'elasticsearch.{0,100}?getting started' --ignore-case --glob-case-insensitive --multiline --multiline-dotall --color=always --glob=\*.\{md,txt,org} --context=20 '/Users/mtm/Documents/Obsidian Vault/' /Users/mtm/pdev/taylormonacelli/notes
[mtm@taylors-MacBook-Pro-2:ourlock(master)]$

```


## example usage


Get files that have golang and charm in them

```log

[mtm@taylors-MacBook-Pro-2:ourlock(master)]$ ourlock existancesearch golang charm --dry-run
2 Permutations:
[golang charm]
[charm golang]
rg golang.+\?charm\|charm.+\?golang --ignore-case --glob-case-insensitive --multiline --multiline-dotall --color=always --glob=\*.\{md,txt,org} --files-with-matches '/Users/mtm/Documents/Obsidian Vault/' /Users/mtm/pdev/taylormonacelli/notes
[mtm@taylors-MacBook-Pro-2:ourlock(master)]$ ourlock existancesearch golang charm
/Users/mtm/Documents/Obsidian Vault/Golang Prompting packages.md
/Users/mtm/Documents/Obsidian Vault/Pushbullet History.md
/Users/mtm/Documents/Obsidian Vault/Charm.md
/Users/mtm/pdev/taylormonacelli/notes/notes.txt
/Users/mtm/Documents/Obsidian Vault/bubbletea.md
/Users/mtm/pdev/taylormonacelli/notes/notes2.txt
[mtm@taylors-MacBook-Pro-2:ourlock(master)]$

```

## install ourlock


on macos/linux:
```bash

brew install gkwa/homebrew-tools/ourlock

```


on windows:

```powershell

TBD

```
