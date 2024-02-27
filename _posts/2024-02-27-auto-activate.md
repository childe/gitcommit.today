---

date: 2024-02-27T16:10:38+0800
layout: post

---

```
function auto_activate_venv() {
    if [ -e "venv/bin/activate" ]; then
        source venv/bin/activate
    else
        local dir=$(pwd)
        while [[ "$dir" != "" && ! -e "$dir/venv/bin/activate" ]]; do
            dir=${dir%/*}
        done
        if [ -e "$dir/venv/bin/activate" ]; then
            source $dir/venv/bin/activate
        fi
    fi
}
#export PROMPT_COMMAND="auto_activate_venv"

function cd() {
  type deactivate > /dev/null && deactivate

  builtin cd $1

  auto_activate_venv
}
auto_activate_venv
```
