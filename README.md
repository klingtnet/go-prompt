# go-prompt

A tiny Go program that generates my WSL bash prompt.

## Setup

Run `make install` and patch your `.bashrc` like this:

```diff
diff --git a/.bashrc b/.bashrc
index 2e8200d..19d30b5 100644
--- a/.bashrc
+++ b/.bashrc
@@ -115,3 +115,12 @@ if ! shopt -oq posix; then
     . /etc/bash_completion
   fi
 fi
+
+if [ -e ~/.local/bin/go-prompt ]; then
+	go_prompt() {
+		go-prompt
+	}
+	export -f go_prompt
+	unset PS1
+	export PROMPT_COMMAND=go_prompt
+fi
````