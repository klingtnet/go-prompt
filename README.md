# go-prompt

A tiny Go program that generates my WSL bash prompt.

## Setup

Run `make install` and patch your `.bashrc` like this (or see this [example](https://github.com/klingtnet/wsl-environment/commit/c0df02cd3a29e46bf1cbb8256377bb981d2624d4)):

```diff
git diff HEAD~
diff --git a/.bashrc b/.bashrc
index 2e8200d..863c583 100644
--- a/.bashrc
+++ b/.bashrc
@@ -115,3 +115,10 @@ if ! shopt -oq posix; then
     . /etc/bash_completion
   fi
 fi
+
+if [ -e ~/.local/bin/go-prompt ]; then
+       go_prompt() {
+               PS1="$(go-prompt $?)"
+       }
+       export PROMPT_COMMAND=go_prompt
+fi
```
