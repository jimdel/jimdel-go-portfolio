<!-- TODO: add front matter support -->

# Comprehensive Markdown Syntax Guide

This document serves as a comprehensive guide to **Markdown**, showcasing its main features and syntax.

## 1. Basic Text Formatting

### Bold Text

To make text bold, wrap it in double asterisks (`**`) or double underscores (`__`):

**This is bold text**  
**This is also bold text**

### Italics Text

To make text italic, wrap it in single asterisks (`*`) or single underscores (`_`):

_This is italic text_  
_This is also italic text_

### Strike-through Text

To strike through text, use double tildes (`~~`):

~~This text is struck through~~

### Code Inline

For inline code, wrap the code snippet in backticks:

`This is inline code`

## 2. Lists

### Unordered List

An unordered list can be created with asterisks (`*`), plus (`+`), or minus (`-`):

- Item 1
- Item 2
  - Subitem 1
  - Subitem 2

### Ordered List

An ordered list uses numbers followed by periods:

1. First item
2. Second item
3. Third item

### Nested Lists

You can nest lists by adding spaces before the bullets:

1. First item
   - Subitem A
   - Subitem B
2. Second item

## 3. Links

### Simple Link

You can create a simple link by wrapping the text in square brackets (`[]`) and the URL in parentheses (`()`):

[OpenAI](https://www.openai.com)

### Link with Title

You can add a title to the link by including it in quotes after the URL:

[OpenAI](https://www.openai.com "OpenAI Homepage")

### Email Link

Email links are created with `mailto:` followed by the email address:

[Contact Us](mailto:contact@openai.com)

## 4. Images

### Simple Image

You can embed an image by adding an exclamation mark (`!`), followed by the alt text in square brackets (`[]`), and the image URL in parentheses (`()`):

![Markdown Logo](https://markdown-here.com/img/icon256.png)

### Image with Title

Like links, images can also have titles:

![Markdown Logo](https://markdown-here.com/img/icon256.png "Markdown Logo")

## 5. Blockquotes

Blockquotes are created by using the greater-than symbol (`>`):

> "This is a blockquote. It can be used to highlight quotes or important text."

You can also nest blockquotes:

> "This is a blockquote."
>
> > "This is a nested blockquote."

## 6. Code Blocks

### Fenced Code Block

A fenced code block is created by using three backticks (```) or three tildes (~~~):

```python
def hello_world():
    print("Hello, World!")
```

## 7. Horizontal Rule

A horizontal rule (line) can be created using three hyphens (`---`), three asterisks (`***`), or three underscores (`___`):

---

## 8. Tables

You can create tables using pipes (`|`) and hyphens (`-`):

| Header 1     | Header 2     | Header 3     |
| ------------ | ------------ | ------------ |
| Row 1, Col 1 | Row 1, Col 2 | Row 1, Col 3 |
| Row 2, Col 1 | Row 2, Col 2 | Row 2, Col 3 |

You can align text in table columns with colons (`:`):

| Header 1 | Header 2 | Header 3 |
| :------: | :------: | :------: |
| Centered | Centered | Centered |
|   Left   |  Right   | Centered |

## 9. Footnotes

You can add footnotes like this:

Here is a footnote reference[^1].

[^1]: This is the footnote text.

## 10. Task Lists

You can create task lists by using dashes (`-`) or asterisks (`*`) followed by square brackets (`[ ]` or `[x]`):

- [x] Task 1 (Completed)
- [ ] Task 2 (Pending)
- [ ] Task 3 (Pending)

## 11. Emoji

Markdown supports a variety of emojis. You can use them by typing the emoji's short code between colons:

:smile: :thumbsup: :rocket:

## 12. HTML Elements

Since Markdown is compatible with HTML, you can include HTML elements inside Markdown. For example, a `<div>` element:

<div style="background-color: lightblue; padding: 20px;">
    This is a custom HTML element inside Markdown.
</div>

## 13. Definition Lists

A definition list is created using a term and a definition:

### Term 1

: This is the definition of Term 1.

### Term 2

: This is the definition of Term 2.

## 14. LaTeX

Markdown supports LaTeX-style math for rendering equations:

Inline math:  
`$\sqrt{x^2 + y^2}$`

Block math:

$$
E = mc^2
$$
