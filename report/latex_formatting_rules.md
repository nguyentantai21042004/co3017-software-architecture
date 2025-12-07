# LaTeX Formatting Rules and Common Errors

This file tracks common LaTeX formatting errors and their solutions for this repository.

## Rules

### 1. General Formatting
- Always escape special characters: `_`, `%`, `$`, `#`, `&`, `{`, `}`.
- Use `\texttt{}` for code snippets inline.
- Use `listings` or `minted` package for code blocks.

## Common Errors and Fixes

### Error: `Undefined control sequence`
- **Cause**: Using a command that hasn't been defined or requires a package.
- **Fix**: Check if the package is included in `main.tex` or define the command.

### Error: `Misplaced alignment tab character &`
- **Cause**: Using `&` in text without escaping it as `\&`.
- **Fix**: Replace `&` with `\&`.

### Error: `Something's wrong--perhaps a missing \item`
- **Cause**: Often caused by list environments (itemize, enumerate) without `\item`.
- **Fix**: Ensure list environments have at least one `\item`.

### Error: `Missing $ inserted`
- **Cause**:
    - Unbalanced `$` signs (e.g., opening but not closing).
    - Using special characters like `_`, `^`, `&` outside of math mode or without escaping.
    - Hidden invalid characters in the text (often from copy-pasting).
- **Fix**:
    - Check for balanced `$` signs.
    - Escape special characters (`\_`, `\&`, `\%`).
    - Retype the problematic line to remove hidden characters.

## File-Specific Notes

### `report/contents/2.3_functional_requirements.tex`
- **Tables**: This file uses `longtable` and `tabularx`.
    - **Constraint**: Do NOT use `\begin{lstlisting}` or `\begin{verbatim}` inside table cells. It will cause compilation errors.
    - **Workaround**: Use `\texttt{code}` for short snippets. For longer code, place it outside the table or use `\parbox` with `\lstinline` (advanced).
- **Special Characters**: Ensure characters like `_`, `&`, `%` are escaped in text columns.

## Adding Code
The repository uses the `listings` package (configured in `main.tex`).

### Correct Usage
```latex
\begin{lstlisting}[language=Java, caption={Code Description}, label={lst:example}]
public class Example {
    public static void main(String[] args) {
        System.out.println("Hello");
    }
}
\end{lstlisting}
```

### Common Pitfalls
- **Inside Tables**: As mentioned, avoid `lstlisting` in tables.
- **Caption/Label**: Always add a caption and label for referencing.

