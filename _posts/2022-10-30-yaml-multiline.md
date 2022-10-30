---

date: 2022-10-30T17:55:11+0800
title: yaml multilines
layout: post

---

```yaml
# Multiple-line strings can be written either as a 'literal block' (using |), 
# or a 'folded block' (using '>').
# Literal block turn every newline within the string into a literal newline (\n).
# Folded block removes newlines within the string.
literal_block: |
  This entire block of text will be the value of the 'literal_block' key,
  with line breaks being preserved.

  The literal continues until de-dented, and the leading indentation is
  stripped.

      Any lines that are 'more-indented' keep the rest of their indentation -
      these lines will be indented by 4 spaces.
folded_style: >
  This entire block of text will be the value of 'folded_style', but this
  time, all newlines will be replaced with a single space.

  Blank lines, like above, are converted to a newline character.

      'More-indented' lines keep their newlines, too -
      this text will appear over two lines.

# |- and >- removes the trailing blank lines (also called literal/block "strip")
literal_strip: |-
  This entire block of text will be the value of the 'literal_block' key,
  with trailing blank line being stripped.
block_strip: >-
  This entire block of text will be the value of 'folded_style', but this
  time, all newlines will be replaced with a single space and 
  trailing blank line being stripped.

# |+ and >+ keeps trailing blank lines (also called literal/block "keep")
literal_keep: |+
  This entire block of text will be the value of the 'literal_block' key,
  with trailing blank line being kept.

block_keep: >+
  This entire block of text will be the value of 'folded_style', but this
  time, all newlines will be replaced with a single space and 
  trailing blank line being kept.
```
