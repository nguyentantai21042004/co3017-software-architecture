# LaTeX Formatting Requirements

Tài liệu này mô tả chi tiết các yêu cầu format cho tất cả các file LaTeX trong dự án báo cáo Kiến trúc Phần mềm (CO3017).

## 1. Cấu trúc Document

### 1.1 Sectioning Commands
- **Section**: Sử dụng `\section{}` cho các phần chính (ví dụ: "Tổng Quan Dự Án", "Phân Tích Bối Cảnh Và Yêu Cầu")
- **Subsection**: Sử dụng `\subsection{}` cho các phần con (ví dụ: "Yêu Cầu Chức Năng", "Ràng Buộc và Giả Định")
- **Subsubsection**: Sử dụng `\subsubsection{}` cho các phần con nhỏ hơn (ví dụ: "User Stories", "Architecture Characteristics")
- **Không sử dụng số thứ tự**: Loại bỏ tất cả số thứ tự như "1.4", "1.4.1" trong tiêu đề

### 1.2 Paragraph Formatting
- **Đoạn văn đầu tiên**: Sử dụng `\indentpar \indentpar` cho đoạn văn đầu tiên sau section/subsection/subsubsection
- **Đoạn văn tiếp theo**: Không cần `\indentpar`, LaTeX sẽ tự động indent theo cấu hình `\parindent`
- **Tiêu đề phụ**: Sử dụng `\noindent\textbf{}` cho các tiêu đề phụ trong nội dung (ví dụ: "Nguyên tắc phân tách Aggregates:")

## 2. Text Formatting

### 2.1 Bold và Italic
- **Bold**: Sử dụng `\textbf{}` cho text in đậm
- **Italic**: Sử dụng `\textit{}` cho text in nghiêng
- **Bold + Italic**: Sử dụng `\textbf{\textit{}}` nếu cần

### 2.2 Quotes
- **Double quotes**: Sử dụng LaTeX quotes ``...'' (backticks và single quotes)
- **Single quotes**: Sử dụng `'...'` (single quotes)

### 2.3 Special Characters
- **Em dash (—)**: Thay bằng `--` (double hyphen)
- **En dash (–)**: Thay bằng `--` (double hyphen)
- **Ampersand (&)**: Escape thành `\&` trong text, hoặc `&` trong bảng
- **Percent (%)**: Escape thành `\%` trong text
- **Dollar ($)**: Escape thành `\$` trong text, hoặc `$...$` cho math mode
- **Hash (#)**: Escape thành `\#`
- **Underscore (_)**: Escape thành `\_` trong text, hoặc `_` trong math mode
- **Caret (^)**: Escape thành `\^{}` trong text, hoặc `^` trong math mode
- **Curly braces**: Escape thành `\{` và `\}`
- **Backslash**: Escape thành `\textbackslash`

### 2.4 Unicode Characters
- **Loại bỏ**: Tất cả các ký tự Unicode đặc biệt như `⸻`, `🔹`, `✅` phải được loại bỏ hoặc thay thế
- **Thay thế**: 
  - `⸻` → Loại bỏ hoặc thay bằng `---`
  - `🔹` → Thay bằng `\noindent\textbf{Lưu ý:}` hoặc `\noindent\textbf{Nguyên tắc:}`
  - `✅` → Loại bỏ hoặc thay bằng `\noindent\textbf{Tổng kết}`

## 3. Mathematical Expressions

### 3.1 Inline Math
- **Số**: Sử dụng `$...$` cho các số và biểu thức toán học ngắn
- **Ví dụ**: `$\geq 90\%$`, `$< 500$ms`, `$\$300$/tháng`

### 3.2 Display Math
- **Biểu thức dài**: Sử dụng `\[...\]` hoặc `$$...$$` cho biểu thức toán học dài
- **Không dùng**: `$$...$$` (deprecated), ưu tiên `\[...\]`

### 3.3 Common Symbols
- **Greater than or equal**: `$\geq$`
- **Less than or equal**: `$\leq$`
- **Greater than**: `$>$`
- **Less than**: `$<$`
- **Arrow**: `$\rightarrow$` hoặc `$\to$`
- **Percent**: `\%` trong text, `$\%$` trong math mode

## 4. Lists

### 4.1 Itemize (Bullet Points)
```latex
\begin{itemize}
    \item Item 1
    \item Item 2
    \begin{itemize}
        \item Sub-item 1
        \item Sub-item 2
    \end{itemize}
\end{itemize}
```

### 4.2 Enumerate (Numbered Lists)
```latex
\begin{enumerate}
    \item First item
    \item Second item
\end{enumerate}
```

### 4.3 List Configuration
- **Spacing**: Đã được cấu hình trong `main.tex` với `enumitem` package:
  - `nosep, topsep=0pt, partopsep=0pt, parsep=0pt, itemsep=0.5em, leftmargin=1.5em`
- **Không cần điều chỉnh thêm**: Các cấu hình này đảm bảo spacing đồng nhất

## 5. Tables

### 5.1 Table Structure
- **Regular tables**: Sử dụng `\begin{table}[ht]` với `tabularx`
- **Long tables**: Sử dụng `\begin{longtable}` cho bảng có thể trải qua nhiều trang
- **Float placement**: Sử dụng `[ht]` (here, top) cho table và figure

### 5.2 Column Alignment
- **Vertical centering**: Sử dụng `m{width}` cho căn giữa dọc
- **Horizontal alignment**:
  - **Centered**: `>{\centering\arraybackslash}m{width}`
  - **Justified**: `>{\noindent\justifying\arraybackslash}X` hoặc `>{\noindent\justifying\arraybackslash}p{width}`
  - **Left aligned**: `>{\raggedright\arraybackslash}X` (không khuyến khích, ưu tiên justify)
  - **Right aligned**: `>{\raggedleft\arraybackslash}X`

### 5.3 Table Formatting Requirements
- **Căn giữa dọc**: Tất cả các ô phải căn giữa theo chiều dọc
- **Căn đều 2 bên (justify)**: Tất cả các cột text phải căn đều 2 bên, không thụt đầu dòng
- **No indent**: Sử dụng `\noindent` trong định nghĩa cột để loại bỏ indent
- **Column width**: Sử dụng `tabularx` với `\textwidth` để bảng tự động điều chỉnh độ rộng
- **Row height**: Sử dụng `\renewcommand{\arraystretch}{1.5}` hoặc giá trị phù hợp để tăng khoảng cách dòng

### 5.4 Table Caption và Label
- **Caption**: Luôn thêm `\caption{}` sau `\end{tabularx}` hoặc `\end{longtable}`
- **Label**: Luôn thêm `\label{}` sau `\caption{}` với format `tab:table_name`
- **Position**: Caption luôn nằm dưới bảng (đã cấu hình trong `main.tex`)

### 5.5 Table Example
```latex
\begin{table}[ht]
\centering
\small
\renewcommand{\tabularxcolumn}[1]{m{#1}}
\renewcommand{\arraystretch}{1.5}
\begin{tabularx}{\textwidth}{|>{\centering\arraybackslash}m{3.5cm}|>{\noindent\justifying\arraybackslash}X|}
\hline
\textbf{Column 1} & \textbf{Column 2} \\
\hline
Content 1 & Content 2 \\
\hline
\end{tabularx}
\renewcommand{\arraystretch}{1.0}
\caption{Table Caption}
\label{tab:table_name}
\end{table}
\FloatBarrier
```

### 5.6 Longtable Example
```latex
\small
\setlength{\tabcolsep}{3pt}
\begin{longtable}{|>{\centering\arraybackslash}m{1.5cm}|>{\noindent\justifying\arraybackslash}p{2.5cm}|}
\caption{Long Table}
\label{tab:long_table}
\\
\hline
\textbf{Column 1} & \textbf{Column 2} \\
\hline
\endfirsthead
\caption[]{Long Table (tiếp theo)}
\\
\hline
\textbf{Column 1} & \textbf{Column 2} \\
\hline
\endhead
\hline
\endfoot
\hline
\endlastfoot
Content 1 & Content 2 \\
\hline
\end{longtable}
\normalsize
```

## 6. Figures

### 6.1 Figure Structure
```latex
\begin{figure}[ht]
    \centering
    \includegraphics[width=0.6\textwidth]{images/figure_name.png}
    \caption{Figure Caption}
    \label{fig:figure_name}
\end{figure}
\FloatBarrier
```

### 6.2 Figure Placement
- **Float placement**: Sử dụng `[ht]` (here, top)
- **Centering**: Luôn sử dụng `\centering` trong figure environment
- **Width**: Điều chỉnh `width` phù hợp (ví dụ: `0.6\textwidth`, `1.0\textwidth`)

### 6.3 Multiple Figures Side-by-Side
```latex
\begin{figure}[ht]
    \centering
    \begin{minipage}{0.48\textwidth}
        \centering
        \includegraphics[width=\textwidth]{images/figure1.png}
    \end{minipage}
    \hfill
    \begin{minipage}{0.48\textwidth}
        \centering
        \includegraphics[width=\textwidth]{images/figure2.png}
    \end{minipage}
    \caption{Two figures side by side}
    \label{fig:two_figures}
\end{figure}
\FloatBarrier
```

### 6.4 Figure Caption và Label
- **Caption**: Luôn thêm `\caption{}` sau `\includegraphics`
- **Label**: Luôn thêm `\label{}` sau `\caption{}` với format `fig:figure_name`
- **Position**: Caption luôn nằm dưới hình (đã cấu hình trong `main.tex`)

## 7. Spacing và Layout

### 7.1 Paragraph Spacing
- **Parskip**: `\setlength{\parskip}{0.3em}` (đã cấu hình trong `main.tex`)
- **Parindent**: `\setlength{\parindent}{1.5em}` (đã cấu hình trong `main.tex`)

### 7.2 Section Spacing
- **Section spacing**: Đã cấu hình trong `main.tex` với `titlesec`:
  - `\titlespacing*{\section}{0pt}{0.2em}{0.2em}`
  - `\titlespacing*{\subsection}{0pt}{0.2em}{0.2em}`
  - `\titlespacing*{\subsubsection}{0pt}{0.2em}{0.2em}`

### 7.3 Float Spacing
- **Float separation**: Đã cấu hình trong `main.tex`:
  - `\setlength{\floatsep}{5pt plus 2pt minus 2pt}`
  - `\setlength{\textfloatsep}{5pt plus 2pt minus 2pt}`
  - `\setlength{\intextsep}{10pt plus 2pt minus 2pt}`

### 7.4 Float Barrier
- **Sử dụng `\FloatBarrier`**: Sau mỗi bảng hoặc hình cuối cùng trong một section/subsection để đảm bảo text tiếp theo không bị nằm ở trang cũ
- **Không dùng `\clearpage`**: Thay bằng `\FloatBarrier` từ package `placeins`

## 8. Code Blocks

### 8.1 Inline Code
- Sử dụng `\texttt{}` cho inline code
- Ví dụ: `\texttt{docker build}`

### 8.2 Code Blocks
- Sử dụng `\begin{verbatim}...\end{verbatim}` cho code blocks
- Hoặc sử dụng `lstlisting` environment nếu cần syntax highlighting

## 9. Packages Required

Các package sau đã được thêm vào `main.tex`:
- `enumitem`: Cho list spacing
- `titlesec`: Cho section spacing
- `tabularx`: Cho bảng tự động điều chỉnh độ rộng
- `longtable`: Cho bảng trải qua nhiều trang
- `caption`: Cho caption formatting
- `placeins`: Cho `\FloatBarrier`
- `ragged2e`: Cho `\justifying` command

## 10. Checklist Formatting

Khi format một file LaTeX mới, kiểm tra:

- [ ] Loại bỏ tất cả số thứ tự trong tiêu đề (section/subsection/subsubsection)
- [ ] Sử dụng `\indentpar \indentpar` cho đoạn văn đầu tiên
- [ ] Sử dụng `\noindent\textbf{}` cho tiêu đề phụ
- [ ] Loại bỏ tất cả ký tự Unicode đặc biệt (⸻, 🔹, ✅)
- [ ] Escape tất cả ký tự đặc biệt (&, %, $, #, _, ^, {, }, \)
- [ ] Thay `—` bằng `--`
- [ ] Sử dụng `$...$` cho số và biểu thức toán học
- [ ] Chuyển bullet points thành `itemize`
- [ ] Chuyển numbered lists thành `enumerate`
- [ ] Format bảng với `tabularx` hoặc `longtable`
- [ ] Căn giữa dọc: sử dụng `m{width}`
- [ ] Căn đều 2 bên: sử dụng `>{\noindent\justifying\arraybackslash}`
- [ ] Thêm `\caption{}` và `\label{}` cho tất cả bảng
- [ ] Thêm `\caption{}` và `\label{}` cho tất cả hình
- [ ] Thêm `\FloatBarrier` sau bảng/hình cuối cùng trong section
- [ ] Kiểm tra không có lỗi LaTeX

## 11. Common Issues và Solutions

### 11.1 Text trong bảng bị thụt đầu dòng
- **Vấn đề**: Text trong ô bảng bị indent
- **Giải pháp**: Thêm `\noindent` vào định nghĩa cột: `>{\noindent\justifying\arraybackslash}`

### 11.2 Text không căn đều 2 bên
- **Vấn đề**: Text trong ô bảng không căn đều
- **Giải pháp**: Sử dụng `\justifying` thay vì `\raggedright`: `>{\noindent\justifying\arraybackslash}`

### 11.3 Caption nằm trên bảng/hình
- **Vấn đề**: Caption xuất hiện phía trên bảng/hình
- **Giải pháp**: Đã cấu hình `position=bottom` trong `main.tex`, kiểm tra lại cấu hình

### 11.4 Text nằm ở trang cũ sau bảng
- **Vấn đề**: Text tiếp theo bị nằm ở trang cũ khi có bảng ở cuối
- **Giải pháp**: Thêm `\FloatBarrier` sau bảng/hình cuối cùng trong section

### 11.5 Bảng không căn giữa dọc
- **Vấn đề**: Nội dung trong ô không căn giữa theo chiều dọc
- **Giải pháp**: Sử dụng `m{width}` thay vì `p{width}`, và thêm `\renewcommand{\tabularxcolumn}[1]{m{#1}}`

### 11.6 Khoảng cách không đồng nhất
- **Vấn đề**: Khoảng cách giữa các phần không đồng nhất
- **Giải pháp**: Đã cấu hình trong `main.tex`, không cần điều chỉnh thêm

## 12. Best Practices

1. **Consistency**: Đảm bảo format đồng nhất trong toàn bộ document
2. **Readability**: Ưu tiên tính dễ đọc, không quá phức tạp
3. **Maintainability**: Sử dụng các cấu hình chung trong `main.tex` thay vì hardcode
4. **Labels**: Sử dụng naming convention nhất quán cho labels (ví dụ: `tab:`, `fig:`)
5. **Comments**: Thêm comments trong code nếu cần giải thích logic phức tạp
6. **Testing**: Luôn compile LaTeX sau khi format để kiểm tra lỗi

## 13. Naming Conventions

### 13.1 Table Labels
- Format: `tab:table_name`
- Ví dụ: `tab:user_stories`, `tab:architecture_characteristics`

### 13.2 Figure Labels
- Format: `fig:figure_name`
- Ví dụ: `fig:usecase-9`, `fig:domain-model-class-diagram`

### 13.3 File Names
- Format: `section_number_section_name.tex`
- Ví dụ: `2.3_functional_requirements.tex`, `2.4_non_functional_requirements.tex`

## 14. Summary

Tài liệu này mô tả đầy đủ các yêu cầu format cho LaTeX documents trong dự án. Tất cả các file `.tex` trong thư mục `report/contents/` phải tuân thủ các quy tắc này để đảm bảo tính nhất quán và chất lượng của báo cáo.

